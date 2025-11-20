package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。
//
// 回文字符串 是正着读和倒过来读一样的字符串。
//
// 子字符串 是字符串中的由连续字符组成的一个序列。
//
//
//
// 示例 1：
//
//
//输入：s = "abc"
//输出：3
//解释：三个回文子串: "a", "b", "c"
//
//
// 示例 2：
//
//
//输入：s = "aaa"
//输出：6
//解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 由小写英文字母组成
//
//
// Related Topics 双指针 字符串 动态规划 👍 1474 👎 0

func Test647(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "示例1",
			input:    "abc",
			expected: 3,
		},
		{
			name:     "示例2",
			input:    "aaa",
			expected: 6,
		},
		{
			name:     "空字符串",
			input:    "",
			expected: 0,
		},
		{
			name:     "单个字符",
			input:    "z",
			expected: 1,
		},
		{
			name:     "复杂示例",
			input:    "abccba",
			expected: 9, // a,b,c,c,b,a,cc,abccba,bccb
		},
		{
			name:     "wrong case1", // 比较特殊
			input:    "fdsklf",
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countSubstrings(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
func countSubstrings(s string) int {
	n := len(s)
	ans := 0

	// dp[i][j] 下标i到j是否为回文串
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, len(s))
	}

	// 因为dp[i][j] 递推关系为 dp[i][j] = dp[i+1][j-1]，所以要倒序遍历，画图很好理解
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] {
				ans++
			}
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
