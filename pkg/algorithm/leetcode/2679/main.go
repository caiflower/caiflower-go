package main

import "fmt"

/**
给你一个下标从 0开始的二维整数数组nums。一开始你的分数为0。你需要执行以下操作直到矩阵变为空：

1.矩阵中每一行选取最大的一个数，并删除它。如果一行中有多个最大的数，选择任意一个并删除。
2.在步骤 1 删除的所有数字中找到最大的一个数字，将它添加到你的 分数中。
请你返回最后的 分数

[[7,2,1],[6,4,2],[6,5,3],[3,2,1]]
15

1 <= nums.length <= 300
1 <= nums[i].length <= 500
0 <= nums[i][j] <= 103

*/

func main() {
	var nums []string = nil
	for _, v := range nums {
		fmt.Printf("%v", v)
	}
}

// matrixSum 暴力法
func matrixSum(nums [][]int) (res int) {
	for k := 0; k < len(nums[0]); k++ {
		m := 0
		for i := 0; i < len(nums); i++ {
			cmax := nums[i][0]
			nums[i][0] = 0
			cj := 0
			for j := 1; j < len(nums[i]); j++ {
				if nums[i][j] > cmax {
					nums[i][cj] = cmax
					cmax = nums[i][j]
					nums[i][j] = 0
					cj = j
				}
			}
			m = max(m, cmax)
		}
		res += m
	}
	return
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	} else {
		return v2
	}
}
