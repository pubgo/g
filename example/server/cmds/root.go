package cmds

import (
	"github.com/pubgo/x/version"
	"github.com/pubgo/x/xcmd"
	"github.com/pubgo/x/xerror"
)

const Service = "server"

// Execute exec
var Execute = xcmd.Init(func(cmd *xcmd.Command) {
	defer xerror.Assert()

	cmd.Use = Service
	cmd.Version = version.Version
	cmd.AddCommand(ExampleCmd())

})
