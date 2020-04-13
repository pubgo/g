package relation

type GroupVisitLog struct {
	ID        int    `json:"id"`         // 自动编号
	CreatedAt int    `json:"created_at"` // 创建时间
	UpdatedAt int    `json:"updated_at"` // 更新时间
	Status    int    `json:"status"`     // 记录状态:  0=可正常使用 1删除
	RecID     string `json:"rec_id"`     // 记录编号
	GroupID   string `json:"group_id"`   // 圈子编号
	InUid     string `json:"in_uid"`     // 圈子成员内部编号(内部流转)
}
