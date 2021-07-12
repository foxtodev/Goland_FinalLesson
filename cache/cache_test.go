package cache

import (
	"testing"
)

const N = 20

var sink int

func BenchmarkFib(b *testing.B) {
	f := NewCache()
	sink = f.Fib(N)
}
