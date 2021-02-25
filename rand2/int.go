package rand2

import (
	"log"
)

// Int63n implements rand.Int63n on the grpcrand global source.
func Int63n(n int64) int64 {
	mu.Lock()
	res := r.Int63n(n)
	mu.Unlock()
	return res
}

// Intn implements rand.Intn on the grpcrand global source.
func Intn(n int) int {
	mu.Lock()
	res := r.Intn(n)
	mu.Unlock()
	return res
}

// Float64 implements rand.Float64 on the grpcrand global source.
func Float64() float64 {
	mu.Lock()
	res := r.Float64()
	mu.Unlock()
	return res
}

// Shuffle
// randomizes the order of elements. n is the number of elements. It
// panics if n < 0. swap swaps the elements with indexes i and j.
func Shuffle(n int, swap func(i, j int)) {
	if n < 0 {
		log.Fatalln("fastrand: argument to Shuffle is < 0")
	}

	// Fisher-Yates
	for i := n - 1; i > 0; i-- {
		j := IntN(i + 1)
		swap(i, j)
	}
}
