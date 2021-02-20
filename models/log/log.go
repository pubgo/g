package log

import "github.com/pubgo/x/models"

type Log struct {
	models.BaseModel

	UserId   string
	Action   string
	Status   string
	Category string
	Msg      string
}
