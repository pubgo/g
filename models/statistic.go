package models

type Statistic struct {
	Status    int8   `json:"status"`  // 记录状态: -1=删除 0=可正常使用
	Version   int8   `json:"version"` // 版本
	ID        int64  `json:"id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt int64  `json:"deleted_at"`
	Slug      string `json:"slug"` // unique_index 人类可读的ID

	// 对什么资源进行分析
	ResType int16
	ResID   int64

	// 标签数
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
