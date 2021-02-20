package collect

import (
	"github.com/pubgo/x/models"
)

// Collect 资源收藏
type Collect struct {
	models.BaseModel

	ResType string // 资源类型
	ResId   string // 资源Id

	UserId   string
	Category string
}
