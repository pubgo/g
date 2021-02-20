package article

import "github.com/pubgo/x/models"

// 权限, 职责
type Permission struct {
	models.BaseModel

	ArticleID string

	// 用户操作判断判断
	IsManager int    `json:"is_manager"`
	IsDonate  int    `json:"is_donate"`
	NeedAuth  string `json:"need_auth,omitempty"`
	IsManage  int    `json:"is_manage"`  //
	IsRead    int    `json:"is_read"`    //
	IsCollect int    `json:"is_collect"` //
	//public - 公开级别 [0 - 私密, 1 - 公开]
	//status - 状态 [0 - 草稿, 1 - 发布]

	// 支付
	PayStatus int `json:"pay_status"` //1:已支付 2：通过企业支付 3:企业订阅 4:个人支付  5 是否是
}
