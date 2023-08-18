package main

import (
	"fmt"
	"math"
)

// [8,-1,6,-7,-4,5,-4,7,-6]
// [11,-10,-11,8,7,-6,9,4,11,6,5,0]

func main() {
	fmt.Printf("%v\n", maximumSum([]int{8, -1, 6, -7, -4, 5, -4, 7, -6}))
}

var dp [100001][2]int

func maximumSum(arr []int) int {
	var res int
	res = math.MinInt
	for i, v := range arr {
		if v >= 0 {
			dp[i+1][0] = max(dp[i][0]+v, v)
			dp[i+1][1] = max(dp[i][1]+v, v)
		} else {
			dp[i+1][0] = max(dp[i][0]+v, v)
			// 计算dp[i+1][1]是一个难点
			dp[i+1][1] = max(dp[i][1]+v, v)
			dp[i+1][1] = max(dp[i][0], dp[i+1][1])
		}
		res = max(res, dp[i+1][0])
		res = max(res, dp[i+1][1])
	}
	if res == 0 {
		res = math.MinInt
		for _, v := range arr {
			res = max(res, v)
		}
	}
	return res
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	} else {
		return v2
	}
}
