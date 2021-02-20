package xcmd1

import (
	"github.com/pubgo/x/xconfig/xconfig_log"
	"github.com/pubgo/x/xdi"
)

func init() {
	// 引入log初始化
	xdi.InitInvoke(func(log xconfig_log.Log) {
		log.Debug().Msg("log init ok")
	})
}
