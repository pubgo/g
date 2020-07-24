package xcmd

import (
	"github.com/pubgo/x/xinit"
	"github.com/pubgo/xerror"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Command = cobra.Command

var rootCmd = &Command{}

func Init(cfn ...func(cmd *Command)) func(fn ...func(*Command)) {
	if len(cfn) != 0 {
		cfn[0](rootCmd)
	}

	if ex, err := os.Executable(); rootCmd.Use == "" && err == nil {
		rootCmd.Use = filepath.Base(ex)
	}

	rootCmd.PersistentPreRunE = func(cmd *Command, args []string) (err error) {
		defer xerror.RespErr(&err)
		xerror.Panic(viper.BindPFlags(cmd.Flags()))
		xinit.Start()
		return
	}

	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	}

	return func(fn ...func(*Command)) {
		defer xerror.RespExit()

		for _, f := range fn {
			f(rootCmd)
		}

		xerror.Panic(rootCmd.Execute())
	}
}

func Args(fn func(cmd *Command)) func(cmd *Command) *Command {
	return func(cmd *Command) *Command {
		if fn != nil {
			fn(cmd)
		}
		return cmd
	}
}
