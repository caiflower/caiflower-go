package leetcode

import (
	"testing"
)

//给定两个单词 word1 和
// word2 ，返回使得
// word1 和
// word2 相同所需的最小步数。
//
// 每步 可以删除任意一个字符串中的一个字符。
//
//
//
// 示例 1：
//
//
//输入: word1 = "sea", word2 = "eat"
//输出: 2
//解释: 第一步将 "sea" 变为 "ea" ，第二步将 "eat "变为 "ea"
//
//
// 示例 2:
//
//
//输入：word1 = "leetcode", word2 = "etco"
//输出：4
//
//
//
//
// 提示：
//
//
//
// 1 <= word1.length, word2.length <= 500
// word1 和 word2 只包含小写英文字母
//
//
// Related Topics 字符串 动态规划 👍 746 👎 0

func Test583(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  int
	}{
		{
			name:  "示例1",
			word1: "sea",
			word2: "eat",
			want:  2,
		},
		{
			name:  "示例2",
			word1: "leetcode",
			word2: "etco",
			want:  4,
		},
		{
			name:  "相同字符串",
			word1: "abc",
			word2: "abc",
			want:  0,
		},
		{
			name:  "完全不同的字符串",
			word1: "abc",
			word2: "def",
			want:  6,
		},
		{
			name:  "一个空字符串",
			word1: "",
			word2: "abc",
			want:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance(tt.word1, tt.word2); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
// 这题目就是求word1和word2的最长公共子序列，(n1 + n2) - 2*求word1和word2的最长公共子序列就是答案
func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return n1 + n2 - (dp[n1][n2] * 2)
}

//leetcode submit region end(Prohibit modification and deletion)
