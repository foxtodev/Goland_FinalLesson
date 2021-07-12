
#### Changed. Random function on refresh web page.
###### main.go

```go
func main() {
	fibImplementations := []web.Fiber{recursion.NewRecursion(), cache.NewCache(), closure.NewClosure()}
	fmt.Printf("Starting server http://localhost:8080/. Use after ?n=10&n=20")
	web.Serve(fibImplementations)
}
```

###### web.go

```go
func Serve(fiber []Fiber) {
    ...
    rand.Seed(time.Now().UnixNano())
    fiberRnd := rand.Intn(len(fiber))
    ...
        res := fiber[fiberRnd].Fib(n)
		if res != -1 {
			data += fmt.Sprintf("Fib(%d) = %d\n", n, res)
		} else {
			data += fmt.Sprintf("Fib(%d) = integer overflow\n", n)
		}
    ...
}
```

#### Added. Check overflow integer.
###### check-overflow.go

```go
func CheckOverflow(n1, n2 int) int {
	if n2 > 0 {
		if n1 > math.MaxInt32-n2 {
			return -1
		}
	} else {
		if n1 < math.MinInt32-n2 {
			return -1
		}
	}
	return n1 + n2
}
```

###### changed fibonacci function
###### closure.go

```go
func gen() func() int {
	a, b := 0, 1
	return func() int {
		defer func() { a, b = b, checkoverflow.CheckOverflow(a, b) }()
		return a
	}
}
```

###### cashe.go

```go
func (c *ctx) Fib(n int) int {
	if n <= 1 {
		return n
	}
	return checkoverflow.CheckOverflow(c.cacheFib(n-1), c.cacheFib(n-2))
}
```

###### recursion.go

```go
func (c *ctx) Fib(n int) int {
	if n <= 1 {
		return n
	}
	return checkoverflow.CheckOverflow(c.Fib(n-1), c.Fib(n-2))
}
```

#### Benchmark
###### main_test.go

```go
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
```
``
go test -bench=.
``
```
goos: linux
goarch: amd64
pkg: github/foxtodev/Goland_FinalLesson
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkRecursion-2    1000000000               0.0000555 ns/op
BenchmarkCache-2        1000000000               0.0000009 ns/op
BenchmarkClosure-2      1000000000               0.0000047 ns/op
PASS
ok      github/foxtodev/Goland_FinalLesson      0.039s
```

<br />
