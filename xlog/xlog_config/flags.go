package xlog_config

import (
	"github.com/spf13/pflag"
)

func GetFlags() *pflag.FlagSet {
	var flags = &pflag.FlagSet{}

	flags.StringVar()
	return flags
}
