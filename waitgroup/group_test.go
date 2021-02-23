package waitgroup

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGroup(t *testing.T) {
	g := New(5)
	for i := 30; i > 0; i-- {
		i := i
		g.Go(func(ctx context.Context) {
			fmt.Println("ok", i)
			time.Sleep(time.Second * 2)
		})
	}
	g.Wait()
	time.Sleep(time.Second * 2)
}
