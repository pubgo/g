package fx

import (
	"context"
	"time"
)

var defaultProcess = &process{}

func MemStatsPrint()                   { defaultProcess.memStatsPrint() }
func CostWith(fn func()) time.Duration { return defaultProcess.costWith(fn) }
func Count(n int) <-chan int           { return defaultProcess.count(n) }

// Tick 简单定时器
// Example: Tick(100, time.Second)
func Tick(args ...interface{}) <-chan time.Time { return defaultProcess.tick(args...) }

// Go
// 启动一个goroutine
func Go(fn func(ctx context.Context)) context.CancelFunc { return defaultProcess.goCtx(fn) }

// GoLoop
// 启动一个goroutine loop
// 是为了替换 `go func() {for{ }}()` 这类的代码
func GoLoop(fn func(ctx context.Context)) context.CancelFunc { return defaultProcess.goLoopCtx(fn) }

func Loop(fn func(i int)) error { return defaultProcess.loopCtx(fn) }

// GoDelay
// 延迟goroutine
func GoDelay(fn func(), dur ...time.Duration)  { defaultProcess.goWithDelay(fn, dur...) }
func Delay(dur time.Duration, fn func()) error { return defaultProcess.delay(dur, fn) }

// Timeout
// 执行超时函数, 超时后, 函数自动退出
func Timeout(dur time.Duration, fn func()) error {
	return defaultProcess.goWithTimeout(dur, fn)
}
