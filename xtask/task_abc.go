package xtask

import "time"

// ITask interface
type ITask interface {
	// all tasks
	Size() int32
	// task that running
	CurSize() int32
	// task average execution time
	CurDur() time.Duration
	// wait for task to finish
	Wait()
	Do(args ...interface{})
	Stop()
	IsClosed() bool
}
