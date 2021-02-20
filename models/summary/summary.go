package summary

import "github.com/pubgo/x/models"

type ArticleGlance struct {
	models.BaseModel

	RecID         string `json:"rec_id"`     // 记录编号
	ArticleID     string `json:"article_id"` // 文章编号
	InUid         string `json:"in_uid"`     // 用户内部编号(内部流转)
	GroupID       string `json:"group_id"`   // 圈子编号
	LinkID        string `json:"link_id"`    // 文章外链编号
	Signature     string `json:"signature"`  // 数据签名
	CompanyDomain string `json:"company_domain"`
}
