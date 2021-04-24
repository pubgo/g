package syncutil

import (
	"fmt"
	"testing"
	"time"
)

func handleTime(wg *WaitGroup, i int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(i) * 5)
	fmt.Println("ok")
}

func TestWaitGroup(t *testing.T) {
	var wg WaitGroup

	for i := 100; i > 0; i-- {
		wg.Inc()
		i := i
		go handleTime(&wg, i)
	}

	wg.Wait()
}

func TestWaitGroupWithCheck(t *testing.T) {
	// enable checked
	var wg = WaitGroup{Check: true}

	for i := 100; i > 0; i-- {
		wg.Inc()
		i := i
		go handleTime(&wg, i)
	}

	wg.Wait()
}
