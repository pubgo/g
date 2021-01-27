package xconfig_rabbitmq

import (
	"fmt"
	"github.com/pubgo/x/kts"
	"github.com/pubgo/x/pkg"
	"github.com/pubgo/x/pkg/randutil"
	"github.com/pubgo/x/retry"
	"github.com/pubgo/x/xconfig"
	"github.com/pubgo/x/xerror"
	"github.com/streadway/amqp"
	"time"
)

type _Connection struct {
	client   *amqp.Connection
	delivery map[string]<-chan amqp.Delivery
	channel  map[string]*amqp.Channel
}

var _mq map[string]*_Connection

// InitRabbitMq init rabbitMQ
func InitRabbitMq() (err error) {
	defer xerror.RespErr(&err)

	// 加载配置
	_cfg := xconfig.Default().Mq.RabbitMQ
	xerror.PanicT(_cfg.Default == "", "default name is empty")
	xerror.PanicT(len(_cfg.Cfg) == 0, "rabbitMQ config count is 0")

	// 默认重试次数
	if _cfg.PublishRetryTime == 0 {
		_cfg.PublishRetryTime = 5
	}

	_mq = make(map[string]*_Connection, len(_cfg.Cfg))

	for _, cfg := range _cfg.Cfg {
		if len(cfg.Channel) > 0 {
			continue
		}

		_conn := xerror.PanicErr(amqp.Dial(cfg.URL)).(*amqp.Connection)
		_mq[cfg.Name] = &_Connection{client: _conn, delivery: make(map[string]<-chan amqp.Delivery), channel: make(map[string]*amqp.Channel)}
		for _, ch := range cfg.Channel {
			_chan := xerror.PanicErr(_conn.Channel()).(*amqp.Channel)
			if ch.ExchangeName != "" {
				if err := _chan.ExchangeDeclarePassive(ch.ExchangeName, ch.ExchangeType, ch.Durable, ch.AutoDelete, false, ch.NoWait, nil); err != nil {
					xerror.PanicM(_chan.ExchangeDeclare(ch.ExchangeName, ch.ExchangeType, ch.Durable, ch.AutoDelete, false, ch.NoWait, nil), "failed to declare an exchange")
				}
			}

			args := make(amqp.Table)
			if ch.XMaxPriority > 0 {
				args["x-max-priority"] = ch.XMaxPriority
			}
			if ch.XQueueMode != "" {
				args["x-queue-mode"] = ch.XQueueMode
			}

			// 用于检查队列是否存在,已经存在不需要重复声明
			if _, err := _chan.QueueDeclarePassive(ch.QueueName, ch.Durable, ch.AutoDelete, ch.Exclusive, ch.NoWait, args); err != nil {
				xerror.PanicErr(_chan.QueueDeclare(ch.QueueName, ch.Durable, ch.AutoDelete, ch.Exclusive, ch.NoWait, args))
			}

			xerror.PanicM(_chan.Qos(ch.PrefetchCount, ch.PrefetchSize, ch.Global), "chan qos")
			xerror.PanicM(_chan.QueueBind(ch.QueueName, ch.RoutingKey, ch.ExchangeName, ch.NoWait, nil), "Failed to bind a queue")
			_mq[cfg.Name].delivery[ch.QueueName] = xerror.PanicErr(_chan.Consume(ch.QueueName, ch.Consumer, ch.AutoAck, ch.Exclusive, ch.NoLocal, ch.NoWait, args)).(<-chan amqp.Delivery)
			_mq[cfg.Name].channel[ch.QueueName] = _chan
		}
	}
	_mq[xconfig.DefaultName] = _mq[_cfg.Default]
	return
}

// GetRedis get redis instance with name
func GetRabbitMq(name ...string) (c map[string]*amqp.Channel) {
	_name := xconfig.DefaultName
	if len(name) > 0 {
		_name = name[0]
	}

	c = _mq[_name].channel
	xerror.PanicT(pkg.IsNone(c), "rabbitMQ instance %s is nil", _name)
	return
}

func GetOneChan(name ...string) *amqp.Channel {
	_name := xconfig.DefaultName
	if len(name) > 0 {
		_name = name[0]
	}

	_chan := _mq[_name]
	xerror.PanicT(pkg.IsNone(_chan), "rabbitMQ instance %s is nil", _name)

	for k := range _chan.channel {
		return _chan.channel[k]
	}

	return nil
}

func Consume(names []string, fn func(amqp.Delivery)) {
	defer xerror.Resp(func(err *xerror.Err) {
		fmt.Println(err.P())
	})

	_name := xconfig.DefaultName
	if len(names) > 0 {
		_name = names[0]
	}

	for {
		select {
		case d := <-_mq[_name].delivery[names[1]]:
			fn(d)
		case <-time.NewTimer(time.Second).C:
		}
	}
}

func ConsumeRPC(names []string, fn func(amqp.Delivery) []byte) {
	defer xerror.Resp(func(err *xerror.Err) {
		fmt.Println(err.P())
	})

	_name := xconfig.DefaultName
	if len(names) > 0 {
		_name = names[0]
	}

	for {
		select {
		case d := <-_mq[_name].delivery[names[1]]:
			xerror.PanicM(_mq[_name].channel[names[1]].Publish(d.Exchange, d.ReplyTo, false, false,
				amqp.Publishing{ContentType: "text/plain", Priority: d.Priority, DeliveryMode: amqp.Persistent, CorrelationId: d.CorrelationId, Body: fn(d),}), "发送数据失败")
		case <-time.NewTimer(time.Second).C:
		}
	}
}

// Publish mq publish
func Publish(exchange, key string, tsk *kts.Task) (err error) {
	defer xerror.RespErr(&err)

	_cfg := xconfig.Default().Mq.RabbitMQ
	return retry.Retry(_cfg.PublishRetryTime, func(u uint, duration time.Duration) {
		if err := GetOneChan().Publish(exchange, key, false, false, amqp.Publishing{
			ContentType:  "text/plain",
			Body:         tsk.Marshal(),
			DeliveryMode: amqp.Persistent,
			Priority:     tsk.Priority,
			Headers:      amqp.Table{},
		}); err != nil {
			xerror.PanicM(err, "消息队列失败 %s", key)
		}
	})
}

// Publish mq publish
func PublishRPC(exchange, key string, tsk *kts.Task) (err error) {
	defer xerror.RespErr(&err)

	_cfg := xconfig.Default().Mq.RabbitMQ

	return gotry.Retry(_cfg.PublishRetryTime, func() {
		xerror.PanicM(GetOneChan().Publish(exchange, key, false, false,
			amqp.Publishing{
				ContentType:   "text/plain",
				Body:          tsk.Marshal(),
				DeliveryMode:  amqp.Persistent,
				Priority:      tsk.Priority,
				Headers:       amqp.Table{},
				CorrelationId: randutil.String(32),
				ReplyTo:       key,
			}), "publish error")
	})
}
