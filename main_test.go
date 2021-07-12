package main

import (
	"github/foxtodev/Goland_FinalLesson/cache"
	"github/foxtodev/Goland_FinalLesson/closure"
	"github/foxtodev/Goland_FinalLesson/recursion"
	"github/foxtodev/Goland_FinalLesson/web"
	"testing"
)

const N = 20

var sink int

var fibImplementations = []web.Fiber{recursion.NewRecursion(), cache.NewCache(), closure.NewClosure()}

func BenchmarkRecursion(b *testing.B) {
	sink = fibImplementations[0].Fib(N)
}

func BenchmarkCache(b *testing.B) {
	sink = fibImplementations[1].Fib(N)
}

func BenchmarkClosure(b *testing.B) {
	sink = fibImplementations[2].Fib(N)
}
