package retry

import (
	"github.com/pubgo/g/pkg/randutil"
	"github.com/pubgo/g/xtry"
	"log"
	"math"
	"time"
)

type retry struct {
	attempt        uint
	delay          time.Duration
	deadline       time.Time
	handle         func(uint, time.Duration)
	strategy       func(uint) time.Duration
	transformation Transformation
}

func Retry(handle ...interface{}) error {
	_r := &retry{}
	for _, h := range handle {
		switch hdl := h.(type) {
		case time.Duration:
			_r.delay = hdl
		case time.Time:
			_r.deadline = hdl
		case uint:
			_r.attempt = hdl
		case func(uint, time.Duration):
			_r.handle = hdl
		case func(uint) time.Duration:
			_r.strategy = hdl
		case Transformation:
			_r.transformation = hdl
		case *retry:
			_r = hdl
		default:
			log.Fatalf("type error %#v", hdl)
		}
	}
	return _r.Do()
}

func WithConstant(factor time.Duration) *retry {
	return &retry{
		strategy: func(uint) time.Duration {
			return factor
		},
	}
}

func WithRandom(factor time.Duration) *retry {
	return &retry{
		strategy: func(uint) time.Duration {
			return time.Duration(randutil.Uint64n(uint64(factor)))
		},
	}
}

func WithIncremental(initial, increment time.Duration) *retry {
	return &retry{
		strategy: func(attempt uint) time.Duration {
			return initial + (increment * time.Duration(attempt))
		},
	}
}

func WithLinear(factor time.Duration) *retry {
	return WithIncremental(0, factor)
}

func WithExponential(factor time.Duration, base float64) *retry {
	return &retry{
		strategy: func(attempt uint) time.Duration {
			return factor * time.Duration(math.Pow(base, float64(attempt)))
		},
	}
}

func WithBinaryExponential(factor time.Duration) *retry {
	return WithExponential(factor, 2)
}

func WithFibonacci(factor time.Duration) *retry {
	fibonacci := func() func() int {
		a, b := 1, 1
		return func() int {
			a, b = b, a+b
			return a
		}
	}()
	return &retry{
		strategy: func(_ uint) time.Duration {
			return factor * time.Duration(fibonacci())
		},
	}
}

func (t *retry) Do() (err error) {
	if t.delay > 0 {
		time.Sleep(t.delay)
	}

	if t.attempt < 1 {
		t.attempt = math.MaxInt32
	}

	if t.transformation == nil {
		t.transformation = func(duration time.Duration) time.Duration {
			return duration
		}
	}

	if t.handle == nil {
		log.Fatalln("handle func is nil")
	}

	var dur time.Duration
	for i := uint(0); i < t.attempt && t.deadline.Before(time.Now()); i++ {
		if err = xtry.Try(func() {
			dur = t.transformation(t.strategy(i))
			t.handle(i, dur)
		})()(); err == nil {
			return
		}
		time.Sleep(dur)
	}
	return
}

func (t *retry) WithDeadline(deadline time.Time) *retry {
	t.deadline = deadline
	return t
}

func (t *retry) WithAttempt(attempt uint) *retry {
	t.attempt = attempt
	return t
}

func (t *retry) WithHandle(handle func(attempt uint, dur time.Duration)) *retry {
	t.handle = handle
	return t
}

func (t *retry) WithDelay(dur time.Duration) *retry {
	t.delay = dur
	return t
}

func (t *retry) WithTransformation(tf Transformation) *retry {
	t.transformation = tf
	return t
}
