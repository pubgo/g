package log

import "github.com/pubgo/g/models"

type Log struct {
	models.BaseModel

	UserId   string
	Action   string
	Status   string
	Category string
	Msg      string
}
