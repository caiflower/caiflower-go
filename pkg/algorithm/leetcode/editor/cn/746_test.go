package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。一旦你支付此费用，即可选择向上爬一个或者两个台阶。
//
// 你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。
//
// 请你计算并返回达到楼梯顶部的最低花费。
//
//
//
// 示例 1：
//
//
//输入：cost = [10,15,20]
//输出：15
//解释：你将从下标为 1 的台阶开始。
//- 支付 15 ，向上爬两个台阶，到达楼梯顶部。
//总花费为 15 。
//
//
// 示例 2：
//
//
//输入：cost = [1,100,1,1,1,100,1,1,100,1]
//输出：6
//解释：你将从下标为 0 的台阶开始。
//- 支付 1 ，向上爬两个台阶，到达下标为 2 的台阶。
//- 支付 1 ，向上爬两个台阶，到达下标为 4 的台阶。
//- 支付 1 ，向上爬两个台阶，到达下标为 6 的台阶。
//- 支付 1 ，向上爬一个台阶，到达下标为 7 的台阶。
//- 支付 1 ，向上爬两个台阶，到达下标为 9 的台阶。
//- 支付 1 ，向上爬一个台阶，到达楼梯顶部。
//总花费为 6 。
//
//
//
//
// 提示：
//
//
// 2 <= cost.length <= 1000
// 0 <= cost[i] <= 999
//
//
// Related Topics 数组 动态规划 👍 1716 👎 0

func Test746(t *testing.T) {
	tests := []struct {
		name string
		cost []int
		want int
	}{
		{name: "示例1", cost: []int{10, 15, 20}, want: 15},
		{name: "示例2", cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, want: 6},
		{name: "两个台阶", cost: []int{1, 2}, want: 1},
		{name: "三个台阶", cost: []int{1, 2, 3}, want: 2},
		{name: "四个台阶", cost: []int{1, 2, 3, 1}, want: 3},
		{name: "全部相同", cost: []int{1, 1, 1, 1}, want: 2},
		{name: "递增", cost: []int{1, 2, 3, 4, 5}, want: 6},
		{name: "递减", cost: []int{5, 4, 3, 2, 1}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, minCostClimbingStairs(tt.cost), tt.name)
		})
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[n]
}

//leetcode submit region end(Prohibit modification and deletion)
