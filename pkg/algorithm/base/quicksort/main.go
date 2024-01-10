// Package main provides ...
package main

import "fmt"

func main() {
	nums := []int{-1, 2, -8, -10}
	quicksort(nums, 0, len(nums)-1)
	fmt.Printf("%v\n", nums)
}

func quicksort(nums []int, l, r int) {
	if l == r {
		return
	}
	mid := (l + r) >> 1
	x := nums[mid]
	i, j := l-1, r+1
	for i < j {
		for {
			i++
			if !(nums[i] < x) {
				break
			}
		}
		for {
			j--
			if !(nums[j] > x) {
				break
			}
		}
		if i < j {
			tmp := nums[i]
			nums[i] = nums[j]
			nums[j] = tmp
			//swap(&nums[i], &nums[j])
		}
	}

	// 关键点是j不是i
	quicksort(nums, l, j)
	quicksort(nums, j+1, r)
}

func swap(x, j *int) {
	tmp := *x
	*x = *j
	*j = tmp
}
