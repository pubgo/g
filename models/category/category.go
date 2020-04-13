package category

import "github.com/pubgo/g/models"

// 资源命名空间

type Category struct {
	models.BaseModel

	// 名字
	Name       string
	CategoryID string `json:"category_id"` // 文章类别编号

	// 公开
	// 用户
	// 权限
	// 公开时间
	// 订阅者
}
