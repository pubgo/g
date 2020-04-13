package log

import "github.com/pubgo/g/models"

type Log struct {
	models.BaseModel

	FromResType string
	FromResId   string
	Relation    string
	ToResType   string
	ToResId     string
}
