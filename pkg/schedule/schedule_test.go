package schedule_test

import (
	"context"
	"github.com/pubgo/x/pkg/schedule"
	"log"
	"testing"
	"time"
)

func EventSecond(param string) {
	log.Printf("second event value:%v\n", param)

	time.Sleep(12 * time.Second)
}

func EventMinute(param string) {
	log.Printf("minute event value:%v\n", param)
}

func EventHour(param string) {
	log.Printf("hour event value:%v\n", param)
}

func EventAtDatetime(param string) {
	log.Printf("AtDatetime event value:%v\n", param)
}

func TestScheduler(t *testing.T) {
	err := schedule.VerySeconds(2).Do(EventSecond, "second")
	if err != nil {
		t.Errorf("test schedule error:%v", err.Error())
		return
	}

	err = schedule.VeryMinutes(1).Do(EventMinute, "minute")
	if err != nil {
		t.Errorf("test schedule error:%v", err.Error())
		return
	}

	err = schedule.VeryHours(1).Do(EventHour, "hour") // minute event value:hours
	if err != nil {
		t.Errorf("test schedule error:%v", err.Error())
		return
	}

	err = schedule.AtDateTime(2018, time.December, 21, 16, 59, 10).Do(EventAtDatetime, "at_datetime")
	if err != nil {
		t.Errorf("test schedule error:%v", err.Error())
		return
	}

	ctx, _ := context.WithCancel(context.Background())
	schedule.Start(ctx)

	select {}
}
