package cmds

import (
	"github.com/pubgo/g/version"
	"github.com/pubgo/g/xcmd"
	"github.com/pubgo/g/xerror"
)

const Service = "server"

// Execute exec
var Execute = xcmd.Init(func(cmd *xcmd.Command) {
	defer xerror.Assert()

	cmd.Use = Service
	cmd.Version = version.Version
	cmd.AddCommand(ExampleCmd())

})
