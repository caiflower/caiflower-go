package bench

import (
	"fmt"
	"testing"
)

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Printf("i = %v, b.N = %v \n", i, b.N)
		fib(30)
	}
}

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}
