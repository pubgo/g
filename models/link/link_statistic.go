package link

import (
	"github.com/pubgo/g/models"
)

type Statistic struct {
	models.BaseModel

	ArticleID string

	InUid string `json:"-"` // 用户内部编号(内部流转)
	Pin string

	Activity float64 `json:"activity"` // 热度

	// 点赞人数
	// 喜欢人数
	// 收藏人数

	// 转发数
	// 转发人数

	// 评论数
	// 评论人数

	// glance 浏览次数
	// 浏览人数

	// 举报数
	// 举报人数

	// 捐赠人数
	// 捐赠收益数 Donate

	// 购买数
	// 购买金额

	// 分类
	// 标签
	// 地理位置
	// 热度计算
	// hash值
	// 修改日期
	// 修改人
}
