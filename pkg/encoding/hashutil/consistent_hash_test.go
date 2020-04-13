package hashutil_test

import (
	"fmt"
	"github.com/pubgo/g/pkg/encoding/hashutil"
	"testing"
)

func TestConsistentHash(t *testing.T) {
	const a = 16

	for i := 0; i < 100; i++ {
		fmt.Println(hashutil.ConsistentHash(uint64(i), a))
	}
}
