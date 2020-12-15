package article

import "github.com/pubgo/x/models"

type Article struct {
	models.BaseModel

	InUid string `json:"-"` // 用户内部编号(内部流转)

	ArticleID string `json:"article_id"` // 文章编号

	Title   string `json:"title"`       // 文章标题
	UserID  string `json:"author_name"` // 作者名称
	GroupID string `json:"group_id"`    // 圈子编号

	Abstract    string `json:"abstract"`  // 文章摘要
	Content     string `json:"content"`   // 用户文章正文
	Format      string                    //描述了正文的格式
	ContentType int `json:"content_type"` // 内容类型

	LinkID    string `json:"link_id"`    // 文章外链编号
	LinkTitle string `json:"link_title"` // 链接标题
	LinkURL   string `json:"link_url"`   // 链接URL

	Pin int `json:"pin"`

	Source   string `json:"source"` //火眼文章来源
	Tags     []int  `json:"tag_id" gorm:"index"`
	Activity string `json:"activity"` //火眼文章热度

	// 最后修改人
	// 所属group或者个人
	// 所属的分类

	Desc string `json:"desc"`

	CoverImageUrl string `json:"cover_image_url"`

	ContentUrl string `json:"content"`  // 内容的链接地址
	Category   int    `json:"category"` // 分类，内容的大分类，金融 财经 区块链等

	ContentHash string `json:"content_hash"` // 用户发布评论的hash值,可以重复
	MarkID      string `json:"mark_id"`      // 文章标注编号

	// 基本信息
	ArticleScore   string   `json:"article_score"`
	HighlightLower []string `json:"highlight"`

	// 创建者 所在圈子 链接
	Author       string   `json:"author" gorm:"ForeignKey:InUid;AssociationForeignKey:InUid"`
	Group        string   `json:"group" gorm:"ForeignKey:GroupID;AssociationForeignKey:GroupID"`
	Link         string   `json:"link" gorm:"ForeignKey:LinkID;AssociationForeignKey:LinkID"`
	ImagesInfo   []string `json:"images_info" gorm:"many2many:article_images;ForeignKey:ArticleID;AssociationForeignKey:ImageID;association_jointable_foreignkey:image_id;jointable_foreignkey:article_id;"`
	MentionsInfo []string `json:"mentions"`

	Description string   `json:"description"`
	Price       int      `json:"price"`
	ProductId   string   `json:"product_id"`
	Name        string   `json:"name"`
	LinkHeadImg string   `json:"link_head_img"` //链接
	ImagesList  []string `json:"images_list"`
}
