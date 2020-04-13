package xcmd

import (
	"github.com/pubgo/g/xenv"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/g/xinit"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

type Command = cobra.Command

var rootCmd = &Command{}

func Init(cfn ...func(cmd *Command)) func(...string) {
	rootCmd.Use = xenv.Cfg.Service
	rootCmd.AddCommand(
		ss(), &Command{
			Use:     "version",
			Aliases: []string{"v"},
			Short:   "version info",
			Run: func(cmd *Command, args []string) {
				xenv.Version()
			},
		})
	rootCmd.PersistentPreRunE = func(cmd *Command, args []string) (err error) {
		defer xerror.RespErr(&err)
		xerror.PanicM(viper.BindPFlags(cmd.Flags()), "Flags Error")
		xerror.PanicM(xinit.Start(), "xinit error")
		return
	}

	// 环境变量处理
	if len(cfn) != 0 {
		cfn[0](rootCmd)
	}

	return func(defaultHome ...string) {
		_defaultHome := "$PWD"
		if len(defaultHome) > 0 {
			_defaultHome = defaultHome[0]
		}
		_defaultHome = os.ExpandEnv(_defaultHome)

		rootCmd.PersistentFlags().StringP("home", "", _defaultHome, "project home dir")
		rootCmd.PersistentFlags().BoolP("debug", "d", false, "debug mode")
		rootCmd.PersistentFlags().StringP("log_level", "l", "", "log level(debug|info|warn|error|fatal|panic)")
		rootCmd.PersistentFlags().StringP("env", "e", "", "running mode(dev|test|stag|prod|release)")
		xerror.Exit(rootCmd.Execute(), "command error")
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
