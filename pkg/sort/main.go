package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{9, 7, 8, 5, 6, 3, 4}

	sort.Ints(nums)

	fmt.Printf("nums = %v \n", nums)
}
