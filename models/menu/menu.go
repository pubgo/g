package menu

import (
	"github.com/pubgo/x/models"
)

type BaseMenu struct {
	models.BaseModel

	MenuLevel uint   `json:"-"`
	ParentId  string `json:"parentId"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	Hidden    bool   `json:"hidden"`
	Component string `json:"component"`

	NickName string     `json:"nickName"`
	Children []BaseMenu `json:"children"`
}
