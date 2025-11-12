package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给你一个非负整数数组 nums 和一个整数 target 。
//
// 向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
//
//
// 例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
//
//
// 返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,1,1,1,1], target = 3
//输出：5
//解释：一共有 5 种方法让最终目标和为 3 。
//-1 + 1 + 1 + 1 + 1 = 3
//+1 - 1 + 1 + 1 + 1 = 3
//+1 + 1 - 1 + 1 + 1 = 3
//+1 + 1 + 1 - 1 + 1 = 3
//+1 + 1 + 1 + 1 - 1 = 3
//
//
// 示例 2：
//
//
//输入：nums = [1], target = 1
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 20
// 0 <= nums[i] <= 1000
// 0 <= sum(nums[i]) <= 1000
// -1000 <= target <= 1000
//
//
// Related Topics 数组 动态规划 回溯 👍 2227 👎 0

func Test494(t *testing.T) {
	tests := []struct {
		name   string
		want   int
		nums   []int
		target int
	}{
		{name: "case1", want: 5, nums: []int{1, 1, 1, 1, 1}, target: 3},
		{name: "case2", want: 1, nums: []int{1}, target: 1},
		{name: "case3", want: 0, nums: []int{1, 1, 1, 1}, target: -1000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findTargetSumWays(tt.nums, tt.target), tt.name)
		})
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
//func findTargetSumWays(nums []int, target int) int {
//	sum := 0
//	for _, v := range nums {
//		sum += v
//	}
//	if sum < target || -sum > target {
//		return 0
//	}
//
//	n := sum*2 + 1
//	dp := make([]int, n)
//	dp[sum] = 1
//
//	for _, v := range nums {
//		dp1 := make([]int, n)
//		// +v
//		for i := n - 1; i >= v; i-- {
//			dp1[i] += dp[i-v]
//		}
//		// -v
//		for i := 0; i <= n-v-1; i++ {
//			dp1[i] += dp[i+v]
//		}
//		dp = dp1
//	}
//
//	return dp[sum+target]
//}

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum < target || -sum > target {
		return 0
	}

	n := sum*2 + 1
	dp := make([]int, n)
	dp[sum] = 1

	for _, v := range nums {
		dp1 := make([]int, n)
		// +v
		for i := n - 1; i >= v; i-- {
			dp1[i] += dp[i-v]
		}
		// -v
		for i := 0; i <= n-v-1; i++ {
			dp1[i] += dp[i+v]
		}
		dp = dp1
	}

	return dp[sum+target]
}

//leetcode submit region end(Prohibit modification and deletion)
