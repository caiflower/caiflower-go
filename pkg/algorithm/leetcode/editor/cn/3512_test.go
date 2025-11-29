package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给你一个整数数组 nums 和一个整数 k。你可以执行以下操作任意次：
//
//
// 选择一个下标 i，并将 nums[i] 替换为 nums[i] - 1。
//
//
// 返回使数组元素之和能被 k 整除所需的最小操作次数。
//
//
//
// 示例 1：
//
//
// 输入： nums = [3,9,7], k = 5
//
//
// 输出： 4
//
// 解释：
//
//
// 对 nums[1] = 9 执行 4 次操作。现在 nums = [3, 5, 7]。
// 数组之和为 15，可以被 5 整除。
//
//
// 示例 2：
//
//
// 输入： nums = [4,1,3], k = 4
//
//
// 输出： 0
//
// 解释：
//
//
// 数组之和为 8，已经可以被 4 整除。因此不需要操作。
//
//
// 示例 3：
//
//
// 输入： nums = [3,2], k = 6
//
//
// 输出： 5
//
// 解释：
//
//
// 对 nums[0] = 3 执行 3 次操作，对 nums[1] = 2 执行 2 次操作。现在 nums = [0, 0]。
// 数组之和为 0，可以被 6 整除。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 1000
// 1 <= nums[i] <= 1000
// 1 <= k <= 100
//
//
// Related Topics 数组 数学 👍 6 👎 0

func Test3512(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "示例1",
			nums:     []int{3, 9, 7},
			k:        5,
			expected: 4,
		},
		{
			name:     "示例2",
			nums:     []int{4, 1, 3},
			k:        4,
			expected: 0,
		},
		{
			name:     "示例3",
			nums:     []int{3, 2},
			k:        6,
			expected: 5,
		},
		{
			name:     "数组只有一个元素",
			nums:     []int{7},
			k:        3,
			expected: 1,
		},
		{
			name:     "数组和已经是k的倍数",
			nums:     []int{5, 10, 15},
			k:        5,
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := minOperations1(test.nums, test.k)
			assert.Equal(t, test.expected, result, "minOperations(%v, %d) = %d, 期望 %d", test.nums, test.k, result, test.expected)
			result = minOperations(test.nums, test.k)
			assert.Equal(t, test.expected, result, "minOperations(%v, %d) = %d, 期望 %d", test.nums, test.k, result, test.expected)
		})
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
// 动态规划 时间复杂度O(n * k * k)
func minOperations1(nums []int, k int) int {
	// dp[j] 指移动到nums[i]所有数的和mod/k余j的最小操作次数
	dp := make([]int, k)
	for i := range dp {
		dp[i] = (1<<63 - 1) / 2
	}
	dp[0] = 0

	// 遍历nums
	for i := 0; i < len(nums); i++ {
		t := make([]int, k)
		copy(t, dp)

		// 遍历dp[j]
		for j := 0; j < k; j++ {
			// 计算状态转移
			for l := 0; l < k; l++ {
				mod := (l + nums[i]) % k
				gap := 0
				if mod >= j {
					gap = mod - j
				} else {
					gap = mod + k - j
				}

				if l != 0 {
					dp[j] = min(gap+t[l], dp[j])
				} else {
					dp[j] = gap + t[l]
				}
			}
		}
	}

	return dp[0]
}

// 贪心，nums的和越大，那么操作数就越小
func minOperations(nums []int, k int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	i := 0
	for i = 1; i*k <= sum; i++ {

	}

	return sum - (i-1)*k
}

//leetcode submit region end(Prohibit modification and deletion)
