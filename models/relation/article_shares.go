package relation

type ArticleShare struct {
	ID             int    `json:"id"`               // 自动编号
	Version        int    `json:"version"`          // 数据库版本
	CreatedAt      int    `json:"created_at"`       // 创建时间
	UpdatedAt      int    `json:"updated_at"`       // 更新时间
	DeleteAt       int    `json:"delete_at"`        // 删除时间
	Status         int    `json:"status"`           // 记录状态: -1=删除 0=可正常使用
	ShareID        string `json:"share_id"`         // 点赞编号
	ShareDna       string `json:"share_dna"`        // 点赞dna
	ShareParentDna string `json:"share_parent_dna"` // 点赞父dna
	InUid          string `json:"in_uid"`           // 用户内部编号(内部流转)
	ArticleID      string `json:"article_id"`       // 用户内部编号(内部流转)
	GroupID        string `json:"group_id"`         // 圈子编号
	SourceGroupID  string `json:"source_group_id"`  // 转入的圈子编号
	Signature      string `json:"signature"`        // 数据签名
}
