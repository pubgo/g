package models

// 权限, 职责
type Permission struct {
	Status    int8   `json:"status"`  // 记录状态: -1=删除 0=可正常使用
	Version   int8   `json:"version"` // 版本
	ID        int64  `json:"id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt int64  `json:"deleted_at"`
	Slug      string `json:"slug"` // unique_index 人类可读的ID

	// 评论的对象
	ResType int16
	ResID   int64

	Public int8 `json:"public"` // 公开的状态
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
