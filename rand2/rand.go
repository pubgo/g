package rand2

import (
	"crypto/rand"
	"encoding/binary"
	"log"
	"math"
	"math/big"
	"strconv"
	"sync/atomic"
	"unsafe"

	"golang.org/x/crypto/blake2b"
)

// randReader
// A randReader produces random values via repeated hashing. The entropy field
// is the concatenation of an initial seed and a 128-bit counter. Each time
// the entropy is hashed, the counter is incremented.
type randReader struct {
	counter      uint64 // First 64 bits of the counter.
	counterExtra uint64 // Second 64 bits of the counter.
	entropy      [32]byte
}

// Reader is a global, shared instance of a cryptographically strong pseudo-
// random generator. It uses blake2b as its hashing function. Reader is safe
// for concurrent use by multiple goroutines.
var Reader *randReader

// init provides the initial entropy for the reader that will seed all numbers
// coming out of fastrand.
func init() {
	r := &randReader{}
	n, err := rand.Read(r.entropy[:])
	if err != nil || n != len(r.entropy) {
		panic("not enough entropy to fill fastrand reader at startup")
	}
	Reader = r
}

// Read
// fills b with random data. It always returns len(b), nil.
func (r *randReader) Read(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	// Grab a unique counter from the reader, while atomically updating the
	// counter so that concurrent callers also end up with unique values.
	counter := atomic.AddUint64(&r.counter, 1)
	counterExtra := atomic.LoadUint64(&r.counterExtra)

	// Increment counterExtra when counter is close to overflowing. We cannot
	// wait until counter == math.MaxUint64 to increment counterExtra, because
	// another goroutine could call Read, overflowing counter to 0 before the
	// first goroutine increments counterExtra. The second goroutine would then
	// be reusing the counter pair (0, 0). Instead, we increment at 1<<63 so
	// that there is little risk of an overflow.
	//
	// There is still a potential overlap near 1<<63, though, because another
	// goroutine could see counter == 1<<63+1 before the first goroutine
	// increments counterExtra. The counter pair (1<<63+1, 1) would then be
	// reused. To prevent this, we also increment at math.MaxUint64. This means
	// that in order for an overlap to occur, 1<<63 goroutine would need to
	// increment counter before the first goroutine increments counterExtra.
	//
	// This strategy means that many counters will be omitted, and that the
	// total space cycle time is potentially as low as 2^126. This is fine
	// however, as the security model merely mandates that no counter is ever
	// used twice.
	if counter == 1<<63 || counter == math.MaxUint64 {
		atomic.AddUint64(&r.counterExtra, 1)
	}

	// Copy the counter and entropy into a separate slice, so that the result
	// may be used in isolation of the other threads. The counter ensures that
	// the result is unique to this thread.
	seed := make([]byte, 64)
	binary.LittleEndian.PutUint64(seed[0:8], counter)
	binary.LittleEndian.PutUint64(seed[8:16], counterExtra)
	// Leave 16 bytes for the inner counter.
	copy(seed[32:], r.entropy[:])

	// Set up an inner counter, that can be incremented to produce unique
	// entropy within this thread.
	n := 0
	innerCounter := uint64(0)
	innerCounterExtra := uint64(0)
	for n < len(b) {
		// Copy in the inner counter values.
		binary.LittleEndian.PutUint64(seed[16:24], innerCounter)
		binary.LittleEndian.PutUint64(seed[24:32], innerCounterExtra)

		// Hash the seed to produce the next set of entropy.
		result := blake2b.Sum512(seed)
		n += copy(b[n:], result[:])

		// Increment the inner counter. Because we are the only thread accessing
		// the counter, we can wait until the first 64 bits have reached their
		// maximum value before incrementing the next 64 bits.
		innerCounter++
		if innerCounter == math.MaxUint64 {
			innerCounterExtra++
		}
	}
	return n, nil
}

// Read is a helper function that calls Reader.Read on b. It always fills b
// completely.
func Read(b []byte) { Reader.Read(b) }

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
	i, _ := rand.Int(Reader, n)
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
