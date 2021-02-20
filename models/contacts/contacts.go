package contacts

import "github.com/pubgo/x/models"

// 名字
// 姓氏
// 所在公司名字
// 是否是公司
// 个人头像
// 电话，手机，iphone，家庭电话，工作电话，主要电话，
// url, 主页，家庭，工作
// 生日
// 住址
// 备注
// 电子邮件
// 昵称
// 朋友
// 及时信息，通讯信息
// 社交账户，社交账户类型，地址，名称，ID，内容，主页

type Contacts struct {
	models.BaseModel

	Name          string `json:"name"` // 用户昵称
	Username      string `json:"-" gorm:"type:varchar(128);index"`
	Salt          string `json:"-"`
	NickName      string `json:"nickName" gorm:"default:'QMPlusUser'"`
	Password      string `json:"password"`        // 用户密码
	PhoneID       string `json:"phone_id"`        // 绑定的手机号
	PhoneBindTime int    `json:"phone_bind_time"` // 手机号绑定时间
	Email         string `json:"email"`           // 绑定的Email
	EmailBindTime int    `json:"email_bind_time"` // Email绑定时间
	Domain        string `json:"domain"`          // 邮箱域名
	AvatarURL     string `json:"avatar_url"`
	HeaderImg     string `json:"headerImg" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`
	HeadImgID     string `json:"head_img_id"`                                                        // 头像id
	BackImgID     string `json:"back_img_id"`                                                        // 背景图id
	HeadImg       string `json:"head_img" gorm:"ForeignKey:HeadImgID;AssociationForeignKey:ImageID"` // 头像id
	BackImg       string `json:"back_img" gorm:"ForeignKey:HeadImgID;AssociationForeignKey:ImageID"` // 背景图id
	ActiveLevel   int    `json:"is_active"`
	UserLevel     string

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
}
