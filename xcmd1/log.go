package xcmd1

import (
	"github.com/pubgo/g/xconfig/xconfig_log"
	"github.com/pubgo/g/xdi"
)

func init() {
	// 引入log初始化
	xdi.InitInvoke(func(log xconfig_log.Log) {
		log.Debug().Msg("log init ok")
	})
}
