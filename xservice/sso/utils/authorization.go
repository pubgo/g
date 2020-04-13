package utils

import (
	"github.com/pubgo/g/xservice/sso/model"
)

func SetDefaultRolesBasedOnConfig() {
	model.InitalizeRoles()
}
