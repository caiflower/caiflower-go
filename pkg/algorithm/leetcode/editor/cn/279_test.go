package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
//
// 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
//
//
//
//
// 示例 1：
//
//
//输入：n = 12
//输出：3
//解释：12 = 4 + 4 + 4
//
// 示例 2：
//
//
//输入：n = 13
//输出：2
//解释：13 = 4 + 9
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁴
//
//
// Related Topics 广度优先搜索 数学 动态规划 👍 2254 👎 0

func Test279(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "示例1",
			input:    12,
			expected: 3,
		},
		{
			name:     "示例2",
			input:    13,
			expected: 2,
		},
		{
			name:     "最小情况",
			input:    1,
			expected: 1,
		},
		{
			name:     "完全平方数",
			input:    16,
			expected: 1,
		},
		{
			name:     "较大数字",
			input:    43,
			expected: 3,
		},
		{
			name:     "1",
			input:    1,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numSquares(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
// 面试题
func numSquares(n int) int {
	// 1. 定义dp数组，dp[i] 为组成i的最小平方数的个数
	dp := make([]int, n+1)

	// 2. 初始化
	for i := range dp {
		dp[i] = 1<<63 - 1
	}
	dp[0] = 0

	// 3. 确定遍历方向
	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			// 4. 状态转移方程
			dp[i] = min(dp[i-j*j]+1, dp[i])
		}
	}

	return dp[n]
}

//leetcode submit region end(Prohibit modification and deletion)
