package rng

import (
	"math/bits"
	"sync"
)

type (
	u64 = uint64
)

const (
	_wyp0 = 0xa0761d6478bd642f
	_wyp1 = 0xe7037ed1a0b428db
	_wyp2 = 0x8ebc6af09c88c6e3
	_wyp3 = 0x589965cc75374cc3
	_wyp4 = 0x1d8e4e27c47d124f
)

func wyMum(A, B u64) u64 {
	hi, lo := bits.Mul64(A, B)
	return hi ^ lo
}

// RNG is a random number generator.
// The zero value is valid.
type RNG uint64

// Int returns a random positive int.
// Not safe for concurrent callers.
func (r *RNG) Int() int {
	return int(uint(r.Uint64()) >> 1)
}

// Intn returns an int uniformly in [0, n).
// Not safe for concurrent callers.
func (r *RNG) Intn(n int) int {
	if n <= 0 {
		return 0
	}
	return int(r.Uint64n(uint64(n)))
}

// Uint64 returns a random uint64.
// Not safe for concurrent callers.
func (r *RNG) Uint64() uint64 {
	*r += _wyp0
	return wyMum(u64(*r)^_wyp1, u64(*r))
}

// Uint64n returns a uint64 uniformly in [0, n).
// Not safe for concurrent callers.
func (r *RNG) Uint64n(n uint64) uint64 {
	if n == 0 {
		return 0
	}

	x := r.Uint64()
	h, l := bits.Mul64(x, n)

	if l < n {
		t := -n
		if t >= n {
			t -= n
			if t >= n {
				t = t % n
			}
		}

	again:
		if l < t {
			x = r.Uint64()
			h, l = bits.Mul64(x, n)
			goto again
		}
	}

	return h
}

// Float64 returns a float64 uniformly in [0, 1).
// Not safe for concurrent callers.
func (r *RNG) Float64() (v float64) {
again:
	v = float64(r.Uint64()>>11) / (1 << 53)
	if v == 1 {
		goto again
	}
	return
}

// global is a parallel rng for the package functions.
var global struct {
	sync.Mutex
	RNG
}

// Int returns a random int.
// Safe for concurrent callers.
func Int() int {
	global.Lock()
	out := global.Int()
	global.Unlock()
	return out
}

// Intn returns a int uniformly in [0, n).
// Safe for concurrent callers.
func Intn(n int) int {
	global.Lock()
	out := global.Intn(n)
	global.Unlock()
	return out
}

// Uint64 returns a random uint64.
// Safe for concurrent callers.
func Uint64() uint64 {
	global.Lock()
	out := global.Uint64()
	global.Unlock()
	return out
}

// Uint64n returns a uint64 uniformly in [0, n).
// Safe for concurrent callers.
func Uint64n(n uint64) uint64 {
	global.Lock()
	out := global.Uint64n(n)
	global.Unlock()
	return out
}

// Float64 returns a float64 uniformly in [0, 1).
// Safe for concurrent callers.
func Float64() float64 {
	global.Lock()
	out := global.Float64()
	global.Unlock()
	return out
}
