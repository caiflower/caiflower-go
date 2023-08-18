package main

import "fmt"

// 知识点前缀和

func main() {
	fmt.Printf("%v\n", pivotInteger(8))
}

var sums [1001]int

func pivotInteger(n int) int {
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + i
	}
	for i := 1; i <= n; i++ {
		if sums[n]-sums[i-1] == sums[i] {
			return i
		}
	}
	return -1
}
