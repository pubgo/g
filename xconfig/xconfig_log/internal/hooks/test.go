package hooks

import (
	"github.com/rs/zerolog"
)

type hostHook struct{}

func (ch hostHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	//v, _ := mem.VirtualMemory()
	//c, _ := cpu.Info()
	//cc, _ := cpu.Percent(time.Second, false)
	//d, _ := disk.Usage("/")
	//n, _ := host.Info()
}

var HH = hostHook{}
