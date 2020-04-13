package retry

import (
	"github.com/pubgo/g/pkg/randutil"
	"github.com/pubgo/g/xtry"
	"log"
	"math"
	"time"
)

type retry struct {
	attempt        uint
	factor         time.Duration
	delay          time.Duration
	deadline       time.Time
	handle         func(uint, time.Duration)
	strategy       func(uint) time.Duration
	transformation Transformation
}

func Retry(factor time.Duration, handle ...interface{}) error {
	_r := &retry{factor: factor}
	for _, h := range handle {
		switch hdl := h.(type) {
		case time.Duration:
			_r.delay = hdl
		case time.Time:
			_r.deadline = hdl
		case uint:
			_r.attempt = hdl
		case func(uint, time.Duration):
			_r.handle = hdl
		case func(uint) time.Duration:
			_r.strategy = hdl
		case Transformation:
			_r.transformation = hdl
		case *retry:
			_r = hdl
		default:
			log.Fatalf("type error %#v", hdl)
		}
	}
	return _r.Do()
}

func With(factor time.Duration) *retry {
	return &retry{
		factor: factor,
	}
}

func Every(factor time.Duration) *retry {
	return With(factor).Constant()
}

func (t *retry) Constant() *retry {
	t.strategy = func(u uint) time.Duration {
		return t.factor
	}
	return t
}

// Random 随机时间
func (t *retry) Random() *retry {
	t.strategy = func(u uint) time.Duration {
		return time.Duration(randutil.Uint64n(uint64(t.factor)))
	}
	return t
}

func (t *retry) RandomIncremental() *retry {
	t.strategy = func(u uint) time.Duration {
		return time.Duration(randutil.Uint64n(uint64(t.factor) * uint64(u)))
	}
	return t
}

// Incremental 线性递增
func (t *retry) Incremental() *retry {
	t.strategy = func(attempt uint) time.Duration {
		return t.factor * time.Duration(attempt)
	}
	return t
}

// Exponential 指数递增
func (t *retry) Exponential(base float64) *retry {
	t.strategy = func(attempt uint) time.Duration {
		return t.factor * time.Duration(math.Pow(base, float64(attempt)))
	}
	return t
}

// BinaryExponential 二进制指数递增
func (t *retry) BinaryExponential() *retry {
	return t.Exponential(2)
}

// Fibonacci 斐波那契递增
func (t *retry) Fibonacci() *retry {
	fibonacci := func() func() int {
		a, b := 1, 1
		return func() int {
			a, b = b, a+b
			return a
		}
	}()
	t.strategy = func(_ uint) time.Duration {
		return t.factor * time.Duration(fibonacci())
	}
	return t
}

// Go 运行
func (t *retry) Go() (err error) {
	return t.Do()
}

// Do 运行
func (t *retry) Do() (err error) {
	if t.delay > 0 {
		time.Sleep(t.delay)
	}

	if t.attempt < 1 {
		t.attempt = math.MaxInt32
	}

	if t.handle == nil {
		log.Fatalln("handle func is nil")
	}

	var dur time.Duration
	for i := uint(1); i <= t.attempt && t.deadline.Before(time.Now()); i++ {
		if err = xtry.Try(func() {
			if t.transformation != nil {
				dur = t.transformation(t.strategy(i))
			}

			t.handle(i, dur)
		})()(); err == nil {
			return
		}
		time.Sleep(dur)
	}
	return
}

// Deadline 重试截止日期
func (t *retry) Deadline(deadline time.Time) *retry {
	t.deadline = deadline
	return t
}

// Attempt 尝试时间
func (t *retry) Attempt(attempt uint) *retry {
	t.attempt = attempt
	return t
}

// Handle 逻辑业务函数
func (t *retry) Handle(handle func(attempt uint, dur time.Duration)) *retry {
	t.handle = handle
	return t
}

// Delay 执行handle之前的延长时间
func (t *retry) Delay(dur time.Duration) *retry {
	t.delay = dur
	return t
}

// Transformation 等待时间转换函数
func (t *retry) Transformation(tf Transformation) *retry {
	t.transformation = tf
	return t
}
