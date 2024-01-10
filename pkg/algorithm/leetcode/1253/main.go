package main

import "fmt"

func main() {
	res := reconstructMatrix(2, 1, []int{1, 1, 1})
	fmt.Printf("%v", res)
}

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	res := make([][]int, 2)
	for i, v := range colsum {
		if i == 0 {
			res[0] = make([]int, len(colsum))
			res[1] = make([]int, len(colsum))
		}
		if v == 2 {
			res[0][i] = 1
			res[1][i] = 1
			upper--
			lower--
		} else if v == 1 {
			if upper >= lower {
				res[0][i] = 1
				upper--
			} else {
				res[1][i] = 1
				lower--
			}
		}
	}
	if upper == 0 && lower == 0 {
		return res
	} else {
		return [][]int{}
	}
}
