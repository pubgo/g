package link

import (
	"github.com/pubgo/x/models"
)

type Statistic struct {
	models.BaseModel

	LinkID int
	// 点赞数
	// 评论数
	// 转发数
	// 喜欢数
	// 收藏数
	// 分类
	// 标签
	// 地理位置

	//slug - 文档路径

	//YuanChuang bool
	// 是否是原创
}
