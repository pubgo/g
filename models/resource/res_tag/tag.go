package res_tag

import "github.com/pubgo/x/models"

type ResTag struct {
	models.BaseModel

	ResType string
	ResId   int
	Tag     int
}

// ResType,ResId,Tag 联合起来唯一, unique_index
// 可以跟任何的资源进行打标签, 图片 文件 链接 文章等
