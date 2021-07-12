package recursion

import checkoverflow "github/foxtodev/Goland_FinalLesson/check-overflow"

type ctx struct {
}

func NewRecursion() *ctx {
	return &ctx{}
}

func (c *ctx) Fib(n int) int {
	if n <= 1 {
		return n
	}
	return checkoverflow.CheckOverflow(c.Fib(n-1), c.Fib(n-2))
}

func (c *ctx) Name() string {
	return "Fibonacci (Recursion)"
}
