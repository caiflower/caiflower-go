package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,5,11,5]
//输出：true
//解释：数组可以分割成 [1, 5, 5] 和 [11] 。
//
// 示例 2：
//
//
//输入：nums = [1,2,3,5]
//输出：false
//解释：数组不能分割成两个元素和相等的子集。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 200
// 1 <= nums[i] <= 100
//
//
// Related Topics 数组 动态规划 👍 2435 👎 0

func Test416(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{name: "示例1", nums: []int{1, 5, 11, 5}, want: true},
		{name: "示例2", nums: []int{1, 2, 3, 5}, want: false},
		{name: "单个元素", nums: []int{1}, want: false},
		{name: "两个相等元素", nums: []int{1, 1}, want: true},
		{name: "总和为奇数", nums: []int{1, 2, 3, 4, 5}, want: false},
		{name: "总和为偶数但无法平分", nums: []int{1, 2, 5}, want: false},
		{name: "可以平分的复杂情况", nums: []int{3, 3, 3, 4, 5}, want: true},
		{name: "大数值测试", nums: []int{100, 100, 100, 100}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, canPartition(tt.nums), tt.name)
		})
	}
}

// leetcode submit region begin(Prohibit modification and deletion)

// 使用01背包方式解决该问题
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum&1 != 0 {
		return false
	}

	v := sum / 2
	dp := make([]int, v+1)

	for _, v1 := range nums {
		for j := v; j >= v1; j-- {
			dp[j] = max(dp[j], dp[j-v1]+v1)
		}
	}

	return dp[v] == v
}

//leetcode submit region end(Prohibit modification and deletion)
