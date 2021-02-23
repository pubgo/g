package randutil

import (
	"fmt"
	"testing"
)

func TestProbability(t *testing.T) {
	for i := 0; i < 10000; i++ {
		t.Log(Probability(0.2))
	}
}

func TestString(t *testing.T) {
	fmt.Println(String(10))
}
