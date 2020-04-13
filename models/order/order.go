package order

import (
	"github.com/pubgo/g/pkg/encoding/hashutil"
	"github.com/pubgo/g/pkg/timeutil"
	"github.com/shopspring/decimal"
	"strconv"
)

type Order struct {
	ID         int             `json:"id"`          // 自动编号
	Version    int             `json:"version"`     // 数据库版本
	CreatedAt  int             `json:"created_at"`  // 创建时间
	UpdatedAt  int             `json:"updated_at"`  // 更新时间
	DeleteAt   int             `json:"delete_at"`   // 删除时间
	Status     int             `json:"status"`      // 订单状态：0=未完成 1=已完成
	OrderID    string          `json:"order_id"`    // 订单编号
	TotalPrice decimal.Decimal `json:"total_price"` // 订单总价
	InUid      string          `json:"in_uid"`      // 用户编号
}

func (o Order) GenerateID() string {
	text := strconv.Itoa(o.CreatedAt) + strconv.Itoa(o.UpdatedAt) + o.InUid
	return hashutil.MD5(text)
}

func InitOrderInstance(inUid string) *Order {
	order := &Order{}
	order.CreatedAt = timeutil.GetSystemCurTime()
	order.UpdatedAt = timeutil.GetSystemCurTime()
	order.InUid = inUid
	order.OrderID = order.GenerateID()
	return order
}

type OrderRep struct {
	ID         uint            `json:"-"`
	CreatedAt  int             `json:"created_at"`
	UpdatedAt  int             `json:"updated_at"`
	Status     int             `json:"status"` // 0：待支付 1：支付成功 2：支付失败 3：取消支付
	OrderID    string          `json:"order_id"`
	TotalPrice decimal.Decimal `json:"total_price"`
	InUid      string          `json:"-"`
}
