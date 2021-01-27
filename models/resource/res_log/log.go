package log

import "github.com/pubgo/x/models"

type Log struct {
	models.BaseModel

	LogId   string
	ResType string
	ResId   string
}
