package closure

import checkoverflow "github/foxtodev/Goland_FinalLesson/check-overflow"

type ctx struct {
}

func NewClosure() *ctx {
	return &ctx{}
}

func gen() func() int {
	a, b := 0, 1
	return func() int {
		defer func() { a, b = b, checkoverflow.CheckOverflow(a, b) }()
		return a
	}
}

func (c *ctx) Fib(n int) int {
	fib := gen()
	for i := 0; i < n; i++ {
		fib()
	}
	return fib()
}

func (c *ctx) Name() string {
	return "Fibonacci (Closure)"
}
