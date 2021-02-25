package signal2

import (
	"os"
)

var shutdownSignals = []os.Signal{os.Interrupt}
