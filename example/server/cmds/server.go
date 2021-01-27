package cmds

import (
	"github.com/pubgo/x/example/server/app"
	"github.com/pubgo/x/xcmd"
	"github.com/pubgo/x/xdi"
	"github.com/pubgo/x/xerror"
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
