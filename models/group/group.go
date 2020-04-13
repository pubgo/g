package group

import "github.com/pubgo/g/models"

// 送鱼 打鱼 晒鱼 买鱼 卖鱼 岛上面挖矿
// island 海岛 开放 封闭
// 仓库
// 矿井 mine 不同类型的矿，需要自己挖取和筛选
// 创建知识，整合数据，协作筛选

type GroupRep struct {
	CreatedAt       int      `json:"created_at"`          // 创建时间
	UpdatedAt       int      `json:"updated_at"`          // 更新时间
	GroupID         string   `json:"group_id"`            // 圈子编号
	ImageID         string   `json:"-"`                   // 图片id
	HeadImgID       string   `json:"-"`                   // 头像id
	BackImgID       string   `json:"-"`                   // 背景图id
	CreateUid       string   `json:"-"`                   // 创建者ID
	Status          int      `json:"status"`              // 圈子状态
	Name            string   `json:"name"`                // 圈子名称
	GroupDesc       string   `json:"group_desc"`          // 圈子描述
	CompanyType     int      `json:"company_type"`        // 企业/企业个人/个人
	GroupType       int      `json:"group_type"`          // 企业/企业个人/个人
	JoinUserType    int      `json:"join_user_type"`      // 0=自由加入 1=申请加入
	JoinArticleType int      `json:"join_article_type"`   // 0=自由转入 1=同意后方可转入 2=仅自己可转入
	PublicType      int      `json:"public_type"`         // 1=私密 2=节点公开 3=企业内公开
	Subscribe       int      `json:"subscribe"`           // 订阅状态
	ArticleCount    int      `json:"article_count"`       // 文章数
	MemberCount     int      `json:"member_count"`        // 成员数
	IsCreator       int      `json:"is_creator"`          // 是管理员
	IsNotify        int      `json:"is_notify,omitempty"` // 订阅状态
	JoinStatus      int      `json:"join_status"`         // 加入状态 0=未加入 1=已加入 2=不可加入
	IsAvailable     int      `json:"is_available"`        // 加入状态 0=不可加入 1=可加入
	ActiveTime      int      `json:"active_time"`         // 更新时间
	CompanyDomain   string   `json:"company_domain"`      // companyDomain
	UnRead          int      `json:"un_read"`
	NodeAddress     string   `json:"node_address"`
	NodeStatus      int      `json:"node_status"`
	HighlightLower  []string `json:"highlight"`
	IsSubscribe     int      `json:"is_subscribe"`
}

type Group struct {
	models.BaseModel

	Name    string `json:"name"`     // 圈子名称
	GroupID string `json:"group_id"` // 圈子编号

	// public - 公开状态 [1 - 公开, 0 - 私密]
	// user_id - 所属的团队/用户编号
	// namespace - 仓库完整路径 user.login/book.slug
	// creator_id - 创建人 User Id
	//likes_count - 喜欢数量
	//watches_count - 订阅数量

	GroupDesc       string `json:"group_desc"`        // 圈子描述
	CreateUid       string `json:"create_uid"`        // 圈子创建者内部编号(内部流转)
	HeadImgID       string `json:"head_img_id"`       // 头像id
	BackImgID       string `json:"back_img_id"`       // 背景图id
	JoinUserType    int    `json:"join_user_type"`    // 0=自由加入 1=申请加入
	JoinArticleType int    `json:"join_article_type"` // 0=自由转入 1=同意后方可转入 2=仅自己可转入
	CompanyType     int    `json:"company_type"`      // 类别: -1=未知 1=企业 2=企业个人 3=个人
	GroupType       int    `json:"group_type"`        // 圈子类型 -1=未知 1=企业 2=企业个人 3=个人 4=企业主页  5=订阅圈子
	PublicType      int    `json:"public_type"`       // 公开: 0=未知 1私密 2节点公开 3企业公开
	CompanyDomain   string `json:"company_domain"`    // 企业域名
	Signature       string `json:"signature"`         // 数据签名
	Subscribe       int    `json:"subscribe"`         // 订阅状态
}
