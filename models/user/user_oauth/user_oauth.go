package user_oauth

import "github.com/pubgo/g/models"

const OAuthGoogle = 1

type UserOAuth struct {
	models.BaseModel

	UserID     uint
	VendorType uint
	VendorID   string
}

type UserOperation struct {
	ID        int    `json:"id"`         // 自动编号
	CreatedAt int    `json:"created_at"` // 创建时间
	UpdatedAt int    `json:"updated_at"` // 更新时间
	Status    int    `json:"status"`     // 记录状态: -1=删除 0=可正常使用
	OType     int    `json:"o_type"`     // 类型   0:发现页访问
	ArticleID string `json:"article_id"` // 文章编号
	InUid     string `json:"in_uid"`     // 用户内部编号(内部流转)
	GroupID   string `json:"group_id"`   // 圈子编号
	Extra     string `json:"extra"`
}
