package randutil

import (
	"math/rand"
	"sync"
	"time"
)

var (
	r  = rand.New(rand.NewSource(time.Now().UnixNano()))
	mu sync.Mutex
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Probability 小于prob的概率, prob is in [0.0,1.0)
func Probability(prob float64) bool {
	if prob > r.Float64() {
		return true
	}
	return false
}
