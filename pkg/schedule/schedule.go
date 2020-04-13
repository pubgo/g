package schedule

import (
	"context"
	"errors"
	"reflect"
	"runtime"
	"sync"
	"time"
)

// Time location, default set by the time.Local (*time.Location)
var timeLocal = time.Local

func InitLocationTime(newLocal *time.Location) {
	timeLocal = newLocal
}

type Task struct {
	isOnce   bool
	interval time.Duration
	running  bool
	lastRun  time.Time
	gName    string
	gFunc    map[string]interface{}
	gParams  map[string][]interface{}
}

type Scheduler struct {
	running bool
	time    *time.Ticker
	tasks   []*Task
	sync.RWMutex
}

var schedule *Scheduler

func newScheduler() *Scheduler {
	if schedule == nil {
		schedule = &Scheduler{
			running: false,
			tasks:   make([]*Task, 0),
		}
	}

	return schedule
}

func very(interval uint64, once bool) *Task {
	if interval <= 0 {
		interval = 1
	}
	newTask := &Task{
		isOnce:   once,
		interval: time.Duration(interval),
		lastRun:  time.Now(),
		gName:    "",
		gFunc:    make(map[string]interface{}, 0),
		gParams:  make(map[string][]interface{}, 0),
	}

	if once {
		newTask.lastRun = time.Unix(int64(interval), 0)
	}

	if schedule == nil {
		newScheduler()
	}

	schedule.Add(newTask)

	return newTask
}

func VerySeconds(interval uint64) *Task {
	return very(interval, false)
}

func VeryMinutes(interval uint64) *Task {
	return very(interval*60, false)
}

func VeryHours(interval uint64) *Task {
	return very(interval*60*60, false)
}

func VeryDays(interval uint64) *Task {
	return very(interval*60*60*24, false)
}

// note: Execute once at a certain point in time
func AtDateTime(year int, month time.Month, day, hour, minute, second int) *Task {
	return very(uint64(time.Date(year, month, day, hour, minute, second, 0, timeLocal).Unix()), true)
}

func (t *Task) Do(taskFun interface{}, params ...interface{}) error {
	typ := reflect.TypeOf(taskFun)
	if typ.Kind() != reflect.Func {
		return errors.New("param taskFun type error")
	}

	funcName := runtime.FuncForPC(reflect.ValueOf(taskFun).Pointer()).Name()
	t.gName = funcName
	t.gFunc[funcName] = taskFun
	t.gParams[funcName] = params

	return nil
}

func (t *Scheduler) checkTaskStatus(isWait bool) bool {
retry:
	for _, taskItem := range t.tasks {
		locTask := taskItem

		if locTask.running {
			if isWait {
				time.Sleep(5 * time.Millisecond)
				goto retry
			}

			return false
		}
	}

	return true
}

func (t *Task) run(locNow time.Time) (result []reflect.Value, err error) {
	if t.isOnce && (locNow.Unix()-t.lastRun.Unix() > 0) {
		return
	}

	if t.running {
		return
	}

	t.running = true
	defer func() {
		t.running = false
	}()

	if !t.isOnce {
		nextTime := t.lastRun.Add(t.interval * time.Second)
		if (locNow.Unix() - nextTime.Unix()) < 0 {
			return
		}
	} else {
		if (locNow.Unix() - t.lastRun.Unix()) < 0 {
			return
		}
	}

	f := reflect.ValueOf(t.gFunc[t.gName])
	params := t.gParams[t.gName]
	if len(params) != f.Type().NumIn() {
		err = errors.New(" param num not adapted ")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)

	if t.isOnce {
		t.lastRun = time.Unix(0, 0)
	}

	t.lastRun = locNow

	return
}

func (t *Scheduler) Add(value *Task) *Scheduler {
	if t.running {
		return t
	}

	if value == nil {
		return t
	}

	t.Lock()
	t.tasks = append(t.tasks, value)
	t.Unlock()

	return t
}

func (t *Scheduler) runAll(locNow time.Time) {

	for _, taskItem := range t.tasks {
		locTask := taskItem

		go func(task *Task) {
			_, _ = task.run(locNow)
		}(locTask)
	}

	return
}

func Start(context context.Context) {
	if schedule == nil {
		newScheduler()
	}

	if schedule.running {
		return
	}

	schedule.Lock()
	schedule.running = true
	schedule.time = time.NewTicker(1 * time.Second)
	schedule.Unlock()

	go func() {
		for {
			select {
			case locNow := <-schedule.time.C:
				schedule.runAll(locNow)

			case <-context.Done():
				schedule.time.Stop()
				schedule.checkTaskStatus(true)
				return
			}
		}
	}()
}
