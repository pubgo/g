package log

import "github.com/pubgo/g/models"

type Log struct {
	models.BaseModel

	LogId   string
	ResType string
	ResId   string
}
