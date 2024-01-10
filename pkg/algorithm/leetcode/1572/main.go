package main

import (
	"fmt"
)

/**
给你一个正方形矩阵 mat，请你返回矩阵对角线元素的和。

请你返回在矩阵主对角线上的元素和副对角线上且不在主对角线上元素的和。

输入：mat = [[1,1,1,1],
            [1,1,1,1],
            [1,1,1,1],
            [1,1,1,1]]
输出：8
*/

func main() {
	nums := [][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}
	fmt.Printf("%v\n", diagonalSum(nums))
}

func diagonalSum(mat [][]int) int {
	res := 0
	for i, j := 0, 0; i < len(mat) && j < len(mat[0]); {
		res += mat[i][j]
		mat[i][j] = 0
		i += 1
		j += 1
	}
	for i, j := 0, len(mat[0])-1; i < len(mat) && j >= 0; {
		res += mat[i][j]
		mat[i][j] = 0
		i += 1
		j -= 1
	}
	return res
}
