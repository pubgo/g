package randutil

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Probability 小于prob的概率, prob is in [0.0,1.0)
func Probability(prob float64) bool {
	if prob > rand.Float64() {
		return true
	}
	return false
}
