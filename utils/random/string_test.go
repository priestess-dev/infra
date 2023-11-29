package random

import (
	"fmt"
	"testing"
)

func TestRandString(t *testing.T) {
	lengths := []int{0, 10, 100, 1001, 1123}
	for _, length := range lengths {
		s := RandString(length)
		if len(s) != length {
			t.Errorf("length of random string is not %d, but %d", length, len(s))
		}
	}
}

func BenchmarkRandString(b *testing.B) {
	for _, length := range []int{100, 1000, 10000} {
		b.Run(fmt.Sprintf("length-%d", length), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				RandString(length)
			}
		})
	}
}
