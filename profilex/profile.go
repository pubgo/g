package profilex

import (
	"os"
	"runtime/pprof"

	"github.com/pubgo/xerror"
)

func CPUProfile(name string) func() {
	f, err := os.Create(name)
	xerror.Panic(err)
	pprof.StartCPUProfile(f)

	return pprof.StopCPUProfile
}
