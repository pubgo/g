package user

import (
	"github.com/pubgo/x/models"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	models.BaseModel

	UUID   uuid.UUID `json:"uuid"`
	InUid  string    `json:"in_uid"`  // 用户内部编号
	OutUid string    `json:"out_uid"` // 用户编号(app端显示)

	Name            string `json:"name"` // 用户昵称
	Username        string `json:"-" gorm:"type:varchar(128);index"`
	Salt            string `json:"-"`
	NickName        string `json:"nickName" gorm:"default:'QMPlusUser'"`
	Password        string `json:"password"`        // 用户密码
	PhoneID         string `json:"phone_id"`        // 绑定的手机号
	PhoneBindTime   int    `json:"phone_bind_time"` // 手机号绑定时间
	Email           string `json:"email"`           // 绑定的Email
	EmailBindTime   int    `json:"email_bind_time"` // Email绑定时间
	Domain          string `json:"domain"`          // 邮箱域名
	AvatarURL       string `json:"avatar_url"`
	SmallAvatarURL  string `json:"small_avatar_url"`
	LargeAvatarURL  string `json:"large_avatar_url"`
	MediumAvatarURL string `json:"medium_avatar_url"`
	Public          int    `json:"public"`
	HeaderImg       string `json:"headerImg" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`
	HeadImgID       string `json:"head_img_id"`                                                        // 头像id
	BackImgID       string `json:"back_img_id"`                                                        // 背景图id
	HeadImg         string `json:"head_img" gorm:"ForeignKey:HeadImgID;AssociationForeignKey:ImageID"` // 头像id
	BackImg         string `json:"back_img" gorm:"ForeignKey:HeadImgID;AssociationForeignKey:ImageID"` // 背景图id
	ActiveLevel     int    `json:"is_active"`
	UserLevel       string

	Category string
	Tags     []string

	WeChatID    string `json:"wechat_id"`
	AuthorityId string `json:"-" gorm:"default:888"`
	UniqueID    string `json:"id" gorm:"type:varchar(128);unique_index"`

	CommDesc  string `json:"comm_desc"` // 自我描述
	Extra     string `json:"extra"`     // 扩展信息
	Signature string `json:"signature"` // 数据签名

	CommAddress   string `json:"comm_address"`   // 通讯地址
	CommCode      string `json:"comm_code"`      // 通讯邮编
	CommSignature string `json:"comm_signature"` // 个性签名

	RegType  int `json:"reg_type"`  // 注册类型: 0=用户名 1=手机 2=Email 3=第三方接口 4=eth
	InType   int `json:"in_type"`   // 用户类别: 0=未知 1=企业用户 2=个人用户
	IsManage int `json:"is_manage"` // 是否是管理员,1=管理员

	Balance string `json:"balance"`
	Address string //地址
}

// 个人信息
	// 昵称
	// 简介
	// 地址
	// 头像
// 账户管理
	// 手机号码
	// 邮箱
	// 帐户密码
	// 个人路径 子域名
	// 绑定第三方帐号
	// 删除帐户
// 会员管理
	// 开通
	// 发票
	// 时长
	// 更改
// 登陆安全
	// 登陆设备 地址 ip 时间
//消息通知 消息+邮件
	// 知识库变更
	// 知识库内容更新
	// 讨论区内容更新
	// 会员消息
	// 团队变更
	// 新评论
	// 提到我
	// 有人关注我
	// 有人赞赏我
	// 团队邀请
	// 团队申请
	// 协作分享
	// 空间邀请
	// 空间添加新成员
	// 空间成员申请
	// 空间成员加入
	// 语雀精选推送
// Token 管理
	// Personal Access Token
// OAuth 应用
// 授权
const (
	// 记录状态:
	CONST_UserBase_Status_Delete             = -1 // 删除
	CONST_UserBase_Status_Normal             = 0  // 可正常使用
	CONST_UserBase_Status_Register           = 1  // 注册中
	CONST_UserBase_Status_RegisterCheck      = 2  // 注册审核中(发送短信,发送Email，第三方授权)
	CONST_UserBase_Status_RegisterFail       = 3  // 注册失败
	CONST_UserBase_Status_BackPassword       = 4  // 找回密码中
	CONST_UserBase_Status_BackPasswordCheck  = 5  // 找回密码审核中
	CONST_UserBase_Status_BackPasswordFail   = 6  // 找回密码失败
	CONST_UserBase_Status_Frozen             = 7  //账户冻结
	CONST_UserBase_Status_PasswordResetCheck = 8  // 密码重置审核中
	CONST_UserBase_Status_PasswordResetFail  = 9  // 密码重置失败
	CONST_UserBase_Status_PasswordReset      = 10 // 密码重置中

	// 注册类型:
	CONST_UserBase_RegType_Name     = 0 // 用户名
	CONST_UserBase_RegType_Phone    = 1 // 手机
	CONST_UserBase_RegType_Email    = 2 // Email
	CONST_UserBase_RegType_Third    = 3 // 第三方接口
	CONST_UserBase_RegType_Ethereum = 4 // Ethereum

	// 用户类型:
	CONST_UserBase_RegType_Company  = 1 // 企业用户
	CONST_UserBase_RegType_Personal = 2 // 个人用户
)
