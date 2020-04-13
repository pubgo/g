package relation

type ArticleGlance struct {
	ID            int    `json:"id"`         // 自动编号
	Version       int    `json:"version"`    // 数据库版本
	CreatedAt     int    `json:"created_at"` // 创建时间
	UpdatedAt     int    `json:"updated_at"` // 更新时间
	DeleteAt      int    `json:"delete_at"`  // 删除时间
	Status        int    `json:"status"`     // 记录状态: -1=删除 0=阅读
	RecID         string `json:"rec_id"`     // 记录编号
	ArticleID     string `json:"article_id"` // 文章编号
	InUid         string `json:"in_uid"`     // 用户内部编号(内部流转)
	GroupID       string `json:"group_id"`   // 圈子编号
	LinkID        string `json:"link_id"`    // 文章外链编号
	Signature     string `json:"signature"`  // 数据签名
	CompanyDomain string `json:"company_domain"`
}
