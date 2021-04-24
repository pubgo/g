package retry

import (
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	With(time.Second).Go()
	Every(time.Second).Go()
	Retry(time.Second,)
}