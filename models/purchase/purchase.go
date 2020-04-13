package purchase

type Purchase struct {
	ID         int    `json:"id"`          // 自动编号
	Version    int    `json:"version"`     // 数据库版本
	CreatedAt  int    `json:"created_at"`  // 创建时间
	UpdatedAt  int    `json:"updated_at"`  // 更新时间
	DeleteAt   int    `json:"delete_at"`   // 删除时间
	Status     int    `json:"status"`      // 状态：-1=删除 0=正常
	PurchaseID string `json:"purchase_id"` // 购买记录编号
	OrderID    string `json:"order_id"`    // 所属订单编号
	ProductID  string `json:"product_id"`  // 购买产品编号
	Amount     int    `json:"amount"`      // 购买数量
	GroupID    string `json:"group_id"`    //圈子ID
	ArticleID  string `json:"article_id"`  //文章ID
	InUid      string `json:"in_uid"`
	FromUid    string `json:"from_uid"`
}

type PurchaseRep struct {
	ID           uint   `json:"-"`
	CreatedAt    int    `json:"created_at"`
	ExpiredAt    int    `json:"expired_at" gorm:"-"` // 过期时间
	UpdatedAt    int    `json:"-"`
	Status       int    `json:"-"`
	PurchaseID   string `json:"purchase_id" form:"purchase_id"`
	OrderID      string `json:"order_id" form:"order_id"`
	ProductID    string `json:"product_id" form:"product_id"`
	GroupID      string `json:"group_id"`
	ArticleTitle string `json:"article_title" gorm:"-"`
	ArticleID    string `json:"article_id"`
	InUid        string `json:"-"`
	FromUid      string `json:"from_uid"`
	Amount       int    `json:"amount" form:"amount"`
}

type GroupSubscribeInfo struct {
	GroupID    string `json:"group_id"`
	SubGroupID string `json:"sub_group_id"`
	StartTime  int    `json:"start_time"`
	ExpireTime int    `json:"expire_time"`
	Amount     int    `json:"amount"`
}
