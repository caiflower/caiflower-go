package main

import "fmt"

func trap(height []int) int {
	lefts := make([]int, len(height))

	m := 0
	for i, v := range height {
		lefts[i] = m
		if v > m {
			m = v
		}
	}

	rights := make([]int, len(height))
	m = 0
	for i := len(height) - 1; i >= 0; i-- {
		rights[i] = m
		if height[i] > m {
			m = height[i]
		}
	}

	res := 0
	for i, v := range height {
		h := 0
		if lefts[i] > v && rights[i] > v {
			if lefts[i] < rights[i] {
				h = lefts[i]
			} else {
				h = rights[i]
			}
			res += h - v
		}
	}

	return res
}

func main() {
	i := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})

	fmt.Println(i)
}
