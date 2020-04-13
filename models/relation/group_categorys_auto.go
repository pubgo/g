package relation

type GroupCategory struct {
	ID         int    `json:"id"`          // 自动编号
	Version    int    `json:"version"`     // 数据库版本
	CreatedAt  int    `json:"created_at"`  // 创建时间
	UpdatedAt  int    `json:"updated_at"`  // 更新时间
	DeleteAt   int    `json:"delete_at"`   // 删除时间
	Status     int    `json:"status"`      // 记录状态: -1=删除 0=可正常使用
	RecID      string `json:"rec_id"`      // 记录编号
	CategoryID string `json:"category_id"` // 圈子类别编号
	GroupID    string `json:"group_id"`    // 圈子编号
	Signature  string `json:"signature"`   // 数据签名
}
