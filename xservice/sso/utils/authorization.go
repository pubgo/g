package utils

import (
	"github.com/pubgo/x/xservice/sso/model"
)

func SetDefaultRolesBasedOnConfig() {
	model.InitalizeRoles()
}
