package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。
//
// 子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。
//
//
//
// 示例 1：
//
//
//输入：s = "bbbab"
//输出：4
//解释：一个可能的最长回文子序列为 "bbbb" 。
//
//
// 示例 2：
//
//
//输入：s = "cbbd"
//输出：2
//解释：一个可能的最长回文子序列为 "bb" 。
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 仅由小写英文字母组成
//
//
// Related Topics 字符串 动态规划 👍 1382 👎 0

func Test516(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "示例1-bbbab",
			input:    "bbbab",
			expected: 4,
		},
		{
			name:     "示例2-cbbd",
			input:    "cbbd",
			expected: 2,
		},
		{
			name:     "单个字符",
			input:    "a",
			expected: 1,
		},
		{
			name:     "两个相同字符",
			input:    "aa",
			expected: 2,
		},
		{
			name:     "两个不同字符",
			input:    "ab",
			expected: 1,
		},
		{
			name:     "全部相同字符",
			input:    "aaaa",
			expected: 4,
		},
		{
			name:     "完全回文字符串",
			input:    "abcba",
			expected: 5,
		},
		{
			name:     "复杂情况1",
			input:    "abcdcba",
			expected: 7,
		},
		{
			name:     "复杂情况2",
			input:    "abcdef",
			expected: 1,
		},
		{
			name:     "长字符串",
			input:    "abacabad",
			expected: 7,
		},
		{
			name:     "wrong case1",
			input:    "aabaaba",
			expected: 6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := longestPalindromeSubseq(tc.input)
			result1 := longestPalindromeSubseq1(tc.input)
			assert.Equal(t, tc.expected, result, "输入: %s %s", tc.input, "longestPalindromeSubseq")
			assert.Equal(t, tc.expected, result1, "输入: %s %s", tc.input, "longestPalindromeSubseq1")
		})
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
func longestPalindromeSubseq(s string) int {
	n := len(s)

	// dp[i][j]是i到j的最长子序列的长度
	dp := make([][]int, n)
	for i := range dp {
		t := make([]int, n)
		t[i] = 1
		dp[i] = t
	}

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if i == j {
				dp[i][j] = 1 // 单个字符,回文长度为1
			} else if s[i] == s[j] {
				// 两端字符相同,在内部回文基础上+2
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				// 两端字符不同,取删除左端或右端的较大值
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][n-1]
}

func longestPalindromeSubseq1(s string) int {
	n := len(s)

	// dp[i][j]是i到j的最长子序列的长度
	dp := make([][]int, n)
	for i := range dp {
		t := make([]int, n)
		t[i] = 1
		dp[i] = t
	}

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] {
				if j-i <= 1 { // 处理越界
					dp[i][j] = j - i + 1
				} else {
					// 如果s[i] == s[j]
					dp[i][j] = dp[i+1][j-1] + 2
				}
			} else {
				if j-i <= 1 { // 处理越界
					dp[i][j] = 1
				} else {
					t := 0
					if i+1 < n { // 处理越界
						t = dp[i+1][j]
					}
					if j-1 >= 0 { // 处理越界
						t = max(dp[i][j-1], t)
					}

					// dp[i][j] = max(dp[i+1][j-1], dp[i+1][j], dp[i][j-1])
					dp[i][j] = max(dp[i+1][j-1], t)
				}
			}
		}
	}

	return dp[0][n-1]
}

//leetcode submit region end(Prohibit modification and deletion)
