package app

import (
	"github.com/pubgo/g/xservice/sso/utils"
	"runtime/debug"

)

func ReloadConfig() {
	debug.FreeOSMemory()
	utils.LoadConfig(utils.CfgFileName)
}
