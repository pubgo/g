package randutil

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"math"
	"math/big"
	"strconv"
	"unsafe"
)

// Uint64n returns a uniform random uint64 in [0,n). It panics if n == 0.
func Uint64n(n uint64) uint64 {
	if n == 0 {
		log.Fatalln("fastrand: argument to Uint64n is 0")
	}

	// To eliminate modulo bias, keep selecting at random until we fall within
	// a range that is evenly divisible by n.
	// NOTE: since n is at most math.MaxUint64, max is minimized when:
	//    n = math.MaxUint64/2 + 1 -> max = math.MaxUint64 - math.MaxUint64/2
	// This gives an expected 2 tries before choosing a value < max.
	max := math.MaxUint64 - math.MaxUint64%n
	var r uint64
	b := (*[8]byte)(unsafe.Pointer(&r))[:]
	Read(b)
	for r >= max {
		Read(b)
	}
	return r % n
}

// IntN
// returns a uniform random int in [0,n). It panics if n <= 0.
func IntN(n int) int {
	if n <= 0 {
		log.Fatalln("rand: argument to Intn is <= 0: " + strconv.Itoa(n))
	}
	// NOTE: since n is at most math.MaxUint64/2, max is minimized when:
	//    n = math.MaxUint64/4 + 1 -> max = math.MaxUint64 - math.MaxUint64/4
	// This gives an expected 1.333 tries before choosing a value < max.
	return int(Uint64n(uint64(n)))
}

// BigIntN
// returns a uniform random *big.Int in [0,n). It panics if n <= 0.
func BigIntN(n *big.Int) *big.Int {
	i, _ := rand.Int(rr, n)
	return i
}

// Perm
// returns a random permutation of the integers [0,n).
func Perm(n int) []int {
	m := make([]int, n)
	for i := 1; i < n; i++ {
		j := IntN(i + 1)
		m[i] = m[j]
		m[j] = i
	}
	return m
}

// Random
// return random string from string slice
func Random(ss []string) []string {
	for i := len(ss) - 1; i > 0; i-- {
		num := IntN(i + 1)
		ss[i], ss[num] = ss[num], ss[i]
	}

	ss1 := make([]string, len(ss))
	for i := 0; i < len(ss); i++ {
		ss1[i] = ss[i]
	}
	return ss1
}

// Ints
// returns a random integer array with the specified from, to and size.
func Ints(from, to, size int) []int {
	if to-from < size {
		size = to - from
	}

	var slice []int
	for i := from; i < to; i++ {
		slice = append(slice, i)
	}

	var ret []int
	for i := 0; i < size; i++ {
		idx := IntN(len(slice))
		ret = append(ret, slice[idx])
		slice = append(slice[:idx], slice[idx+1:]...)
	}

	return ret
}

// Int
// returns a random integer in range [min, max].
func Int(min int, max int) int {
	return min + IntN(max-min)
}

func Bits(b []byte) {
	Read(b)
}

// Bytes
// a helper function that returns n bytes of random data.
func Bytes(n int) []byte {
	b := make([]byte, n)
	Read(b)
	return b
}

func String(n int) string {
	return hex.EncodeToString(Bytes(n))
}

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
