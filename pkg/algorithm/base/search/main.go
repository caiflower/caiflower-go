package main

import (
	"fmt"
	"sort"
)

func main() {
	//index := binarySearch([]int{0, 1, 2, 3, 5, 6, 7}, 9)
	//fmt.Println(index)

	nums := []int{0, 1, 2, 3, 5, 6, 7}
	fmt.Println(binarySearchSort(nums, 0) == binarySearch(nums, 0))
	fmt.Println(binarySearchSort(nums, 5) == binarySearch(nums, 5))
	fmt.Println(binarySearchSort(nums, 7) == binarySearch(nums, 7))
	fmt.Println(binarySearchSort(nums, 9) == binarySearch(nums, 9))
}

func binarySearchSort(nums []int, target int) int {
	search := sort.Search(len(nums), func(i int) bool {
		if nums[i] < target {
			return false
		}
		return true
	})
	if search >= len(nums) {
		return -1
	}
	if nums[search] == target {
		return search
	} else {
		return -1
	}
}

func binarySearch(arr []int, target int) int {
	l, r := 0, len(arr)
	for l < r {
		m := (l + r) / 2
		if arr[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}

	if l >= len(arr) {
		return -1
	}

	return l
}
