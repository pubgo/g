package process

import (
	"time"
)

// PidAlive checks whether a pid is alive.
func PidAlive(pid int) bool {
	return _pidAlive(pid)
}

// PidWait blocks for a process to exit.
func PidWait(pid int) error {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if !PidAlive(pid) {
			break
		}
	}

	return nil
}
