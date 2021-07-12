package checkoverflow

import (
	"math"
)

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
