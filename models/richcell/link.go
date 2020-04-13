package link

import (
	"github.com/pubgo/g/models"
)

type Link struct {
	models.BaseModel

	SourceName string `json:"source"`
	SourceUrl  string `json:"source"`

	Category string

	GeoPoint []string

	// 分类
	// 标签
	// 地理位置

	//slug - 文档路径

	//YuanChuang bool
	// 是否是原创

	Title        string `json:"title"` // 文章标题
	Url          string `json:"url"`   // 文章外链url
	Type         string
	Tags         []string
	LinkID       string `json:"link_id"`      // 文章编号
	Abstract     string `json:"abstract"`     // 文章摘要
	Content      string `json:"content"`      // 文章的内容
	PublishTime  int    `json:"publish_time"` // 文章的发布时间
	Publisher    string `json:"publisher"`    // 文章作者
	PublisherUrl string `json:"publisher"`    // 文章作者
	SourceGroup  string
	ContentHash  string `json:"content_hash"`   // 文章的hash值
	HeadImageUrl string `json:"head_image_url"` // 文章缩略图
}
