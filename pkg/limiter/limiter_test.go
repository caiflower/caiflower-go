package limiter

import (
	"fmt"
	"testing"
)

func BenchmarkRun(b *testing.B) {
	var errCount, succCount int
	for i := 0; i < b.N; i++ {
		result := Run()
		if result != Success {
			errCount++
		} else {
			succCount++
		}
	}
	fmt.Printf("succcount is %v, errcount is %v\n", succCount, errCount)
}
