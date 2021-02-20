package log

import "github.com/pubgo/x/models"

type Log struct {
	models.BaseModel

	FromResType string
	FromResId   string
	Relation    string
	ToResType   string
	ToResId     string
}
