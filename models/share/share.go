package share

import "github.com/pubgo/x/models"

type Share struct {
	models.BaseModel

	ShareID        string `json:"share_id"`         // 点赞编号
	InUid          string `json:"in_uid"`           // 用户内部编号(内部流转)
	ArticleID      string `json:"article_id"`       // 用户内部编号(内部流转)
	GroupID        string `json:"group_id"`         // 圈子编号
	SourceGroupID  string `json:"source_group_id"`  // 转入的圈子编号
	Signature      string `json:"signature"`        // 数据签名
}
