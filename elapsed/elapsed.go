package elapsed

import (
	"fmt"
	"time"
)

func TimeFunc(start time.Time) string {
	elapsed := time.Since(start)
	return fmt.Sprintf("Time elapsed took %s\n", elapsed)
}
