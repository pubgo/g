package user_auth

type UserAuth struct {
	ID          int    `json:"id"`          // 自动编号
	Version     int    `json:"version"`     // 数据库版本
	CreatedAt   int    `json:"created_at"`  // 创建时间
	UpdatedAt   int    `json:"updated_at"`  // 更新时间
	DeleteAt    int    `json:"delete_at"`   // 删除时间
	RecID       string `json:"rec_id"`      // 记录编号
	Status      int    `json:"status"`      // 记录状态: -1=删除 0=可正常使用
	InUid       string `json:"in_uid"`      // 用户内部编号(内部流转)
	AuthType    int    `json:"auth_type"`   // 授权类型: 0=用户名 1=手机号 2=邮箱 3=qq 4=微信 5=腾讯微博 6=新浪微博 7=ethereum
	Identifier  string `json:"identifier"`  // 授权的手机号 邮箱 用户名或第三方应用的唯一标识
	Certificate string `json:"certificate"` // 短信验证码、邮箱验证码、第三方token
}

const (
	// 记录状态:
	CONST_Auth_Status_Delete  = -1 // 删除
	CONST_Auth_Status_Normal  = 0  // 可正常使用
	CONST_Auth_Status_Pending = 1  // 待验证

	// 授权类型: 0=用户名 1=手机号 2=邮箱 3=qq 4=微信 5=腾讯微博 6=新浪微博 7=ethereum
	CONST_Auth_Type_Name     = 0
	CONST_Auth_Type_Phone    = 1
	CONST_Auth_Type_Email    = 2
	CONST_Auth_Type_QQ       = 3
	CONST_Auth_Type_Wechat   = 4
	CONST_Auth_Type_Qbo      = 5
	CONST_Auth_Type_Weibo    = 6
	CONST_Auth_Type_Ethereum = 7
)
