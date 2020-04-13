package cmds

import (
	"github.com/pubgo/g/example/server/app"
	"github.com/pubgo/g/xcmd"
	"github.com/pubgo/g/xdi"
	"github.com/pubgo/g/xerror"
)

func ExampleCmd() *xcmd.Command {
	return &xcmd.Command{
		Use:   "s",
		Short: "simple encryption",
		RunE: func(cmd *xcmd.Command, args []string) (err error) {
			defer xerror.RespErr(&err)
			xerror.Panic(xdi.Invoke(func(ctrl app.App) {
				ctrl.InitRoutes()
				ctrl.Run()
			}))
			return
		},
	}
}
