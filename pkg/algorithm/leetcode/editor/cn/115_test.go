package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给你两个字符串 s 和 t ，统计并返回在 s 的 子序列 中 t 出现的个数。
//
// 测试用例保证结果在 32 位有符号整数范围内。
//
//
//
// 示例 1：
//
//
//输入：s = "rabbbit", t = "rabbit"
//输出：3
//解释：
//如下所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
//rabbbit
//rabbbit
//rabbbit
//
// 示例 2：
//
//
//输入：s = "babgbag", t = "bag"
//输出：5
//解释：
//如下所示, 有 5 种可以从 s 中得到 "bag" 的方案。
//babgbag
//babgbag
//babgbag
//babgbag
//babgbag
//
//
//
//
// 提示：
//
//
// 1 <= s.length, t.length <= 1000
// s 和 t 由英文字母组成
//
//
// Related Topics 字符串 动态规划 👍 1381 👎 0

func Test115(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected int
	}{
		{
			s:        "rabbbit",
			t:        "rabbit",
			expected: 3,
		},
		{
			s:        "rabbbitb",
			t:        "rabbit",
			expected: 3,
		},
		{
			s:        "babgbag",
			t:        "bag",
			expected: 5,
		},
		{
			s:        "a",
			t:        "a",
			expected: 1,
		},
		{
			s:        "abc",
			t:        "",
			expected: 1,
		},
		{
			s:        "",
			t:        "a",
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			result := numDistinct(test.s, test.t)
			assert.Equal(t, test.expected, result, "s=%s, t=%s", test.s, test.t)
		})
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
// O(m*n)
func numDistinct(s string, t string) int {
	n1, n2 := len(s), len(t)

	// s中下标以i结尾，t中下标以j结尾的子序列个数
	dp := make([][]int, n1+1)
	for i := range dp {
		tmp := make([]int, n2+1)
		tmp[0] = 1
		dp[i] = tmp
	}

	// 状态转移方程，对我来说hard实至名归，状态转移方程太难想了
	// 如果s[i] == t[j]，可以加上 dp[i][j] += dp[i-1][j-1]，同时可以不以i为结尾，有k小于i，使得s[k] == t[j], dp[i][j] += dp[k][j]

	for i := 1; i <= n1; i++ {
		for j := 1; j <= i && j <= n2; j++ {
			// j等于i为结尾的情况
			if s[i-1] == t[j-1] {
				dp[i][j] += dp[i-1][j-1]
			}

			dp[i][j] += dp[i-1][j]
		}
	}

	return dp[n1][n2]
}

//leetcode submit region end(Prohibit modification and deletion)
