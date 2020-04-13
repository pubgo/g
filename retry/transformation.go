package retry

import (
	"math"
	"math/rand"
	"time"
)

// Transformation 对生成的等待时间进行转换调整
type Transformation func(duration time.Duration) time.Duration

// Full creates a Transformation that transforms a duration into a result
// duration in [0, n) randomly, where n is the given duration.
//
// The given generator is what is used to determine the random transformation.
//
// Inspired by https://www.awsarchitectureblog.com/2015/03/backoff.html
func (t *retry) Full() Transformation {
	return func(duration time.Duration) time.Duration {
		return time.Duration(rand.Int63n(int64(duration)))
	}
}

// Equal creates a Transformation that transforms a duration into a result
// duration in [n/2, n) randomly, where n is the given duration.
//
// The given generator is what is used to determine the random transformation.
//
// Inspired by https://www.awsarchitectureblog.com/2015/03/backoff.html
func (t *retry) Equal() Transformation {
	return func(duration time.Duration) time.Duration {
		return (duration / 2) + time.Duration(rand.Int63n(int64(duration))/2)
	}
}

// Deviation creates a Transformation that transforms a duration into a result
// duration that deviates from the input randomly by a given factor.
//
// The given generator is what is used to determine the random transformation.
//
// Inspired by https://developers.google.com/api-client-library/java/google-http-java-client/backoff
func (t *retry) Deviation(factor float64) Transformation {
	rand.Seed(time.Now().UnixNano())
	return func(duration time.Duration) time.Duration {
		min := int64(math.Floor(float64(duration) * (1 - factor)))
		max := int64(math.Ceil(float64(duration) * (1 + factor)))
		return time.Duration(rand.Int63n(max-min) + min)
	}
}

// NormalDistribution creates a Transformation that transforms a duration into a
// result duration based on a normal distribution of the input and the given
// standard deviation.
//
// The given generator is what is used to determine the random transformation.
func (t *retry) NormalDistribution(standardDeviation float64) Transformation {
	rand.Seed(time.Now().UnixNano())
	return func(duration time.Duration) time.Duration {
		return time.Duration(rand.NormFloat64()*standardDeviation + float64(duration))
	}
}
