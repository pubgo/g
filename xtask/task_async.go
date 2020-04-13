package xtask

import (
	"fmt"
	"github.com/pubgo/g/xtry"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/pubgo/g/xerror"
)

// NewAsyncTask
// max: 最大并发数
// maxDur: 任务最大执行时间
func NewAsyncTask(max int32, maxDur time.Duration, fn interface{}) ITask {
	_t := &_AsyncTask{
		max:     max,
		maxDur:  maxDur,
		_tfn:    xtry.Try(fn),
		mux:     &sync.Mutex{},
		_curDur: make(chan time.Duration, max),
		taskQ:   make(chan func(...interface{}) (err error), max),
	}
	go _t._loop()
	return _t
}

// AsyncTask struct
type _AsyncTask struct {
	ITask

	_tfn func(...interface{}) func(...interface{}) (err error)

	wg sync.WaitGroup

	max    int32
	maxDur time.Duration

	curDur  time.Duration
	_curDur chan time.Duration

	taskL int32
	taskQ chan func(...interface{}) (err error)

	mux *sync.Mutex

	_stop1 int32
	_stop  sync.Once
}

// Size
// get current xtask size
func (t *_AsyncTask) Size() int32 {
	return atomic.LoadInt32(&t.taskL)
}

// IsClosed
func (t *_AsyncTask) IsClosed() bool {
	return atomic.LoadInt32(&t._stop1) == 1
}

// CurSize
// get current active xtask size
func (t *_AsyncTask) CurSize() int32 {
	return atomic.LoadInt32(&t.taskL) - int32(len(t.taskQ))
}

// CurDur
// current duration
func (t *_AsyncTask) CurDur() time.Duration {
	return t.curDur
}

// Wait wait for xtask done
func (t *_AsyncTask) Wait() {
	t.wg.Wait()
}

func (t *_AsyncTask) taskAdd(fn func(...interface{}) error) {
	if !t.IsClosed() {
		t.taskQ <- fn
		t.wg.Add(1)
		atomic.AddInt32(&t.taskL, 1)
	}
}

func (t *_AsyncTask) taskDone(ct time.Time) {
	if !t.IsClosed() {
		t.wg.Done()
		atomic.AddInt32(&t.taskL, -1)
		t._curDur <- time.Now().Sub(ct)
	}
}

// Do
// handle xtask
func (t *_AsyncTask) Do(args ...interface{}) {
	if t.IsClosed() {
		return
	}

	t.mux.Lock()
	defer t.mux.Unlock()

	for !t.IsClosed() {
		if t.Size() < t.max && t.curDur < t.maxDur {
			t.taskAdd(t._tfn(args...))
			return
		}

		if t.Size() < int32(runtime.NumCPU()*2) {
			t.curDur = 0
		}

		xerror.Debug(fmt.Sprintf("info q_l:%d, cur_dur:%s, max_q:%d, max_dur:%s, ", t.taskL, t.curDur, t.max, t.maxDur))
		time.Sleep(time.Microsecond)
	}
}

// _taskHandle
// 此处不允许出错, 所有的错误必须在worker中自行处理
func (t *_AsyncTask) _taskHandle(fn func(...interface{}) (err error)) {
	go func() {
		_t := time.Now()
		xerror.Exit(fn, "XTask.AsyncTask: max_task_queue_len:%d, task_queue_len:%d, cur_dur:%s, max_dur:%s", t.max, t.taskL, t.curDur, t.maxDur)
		t.taskDone(_t)
	}()
}

// Stop stop xtask
func (t *_AsyncTask) Stop() {
	t._stop.Do(func() {
		atomic.CompareAndSwapInt32(&t._stop1, 0, 1)
		atomic.CompareAndSwapInt32(&t.taskL, t.taskL, 0)
		close(t._curDur)
		close(t.taskQ)
	})
}

func (t *_AsyncTask) _loop() {
	for {
		select {
		case _fn, ok := <-t.taskQ:
			if ok {
				t._taskHandle(_fn)
			}
		case _curDur, ok := <-t._curDur:
			if ok {
				t.curDur = (t.curDur + _curDur) / 2
			}
		}
	}
}

func Wait(tasks ...ITask) {
	var isWait bool
	for {
		isWait = false
		for _, t := range tasks {
			if t.IsClosed() {
				continue
			}

			if t.Size() > 0 {
				t.Wait()
				isWait = true
				break
			}
		}

		if !isWait {
			return
		}
	}
}

func Stop(tasks ...ITask) {
	for _, t := range tasks {
		t.Stop()
	}
}

func WaitAndStop(tasks ...ITask) {
	Wait(tasks...)
	Stop(tasks...)
}
