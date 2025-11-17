package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
//
// 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而
//"aec"不是）。
//
// 进阶：
//
// 如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代
//码？
//
// 致谢：
//
// 特别感谢 @pbrother 添加此问题并且创建所有测试用例。
//
//
//
// 示例 1：
//
//
//输入：s = "abc", t = "ahbgdc"
//输出：true
//
//
// 示例 2：
//
//
//输入：s = "axc", t = "ahbgdc"
//输出：false
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 100
// 0 <= t.length <= 10^4
// 两个字符串都只由小写字符组成。
//
//
// Related Topics 双指针 字符串 动态规划 👍 1208 👎 0

func Test392(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "示例1",
			s:    "abc",
			t:    "ahbgdc",
			want: true,
		},
		{
			name: "示例2",
			s:    "axc",
			t:    "ahbgdc",
			want: false,
		},
		{
			name: "空字符串",
			s:    "",
			t:    "ahbgdc",
			want: true,
		},
		{
			name: "相同字符串",
			s:    "ahbgdc",
			t:    "ahbgdc",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSubsequence(tt.s, tt.t)
			assert.Equal(t, tt.want, got)
		})
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
// O(n * m)
func isSubsequence(s string, t string) bool {
	n1, n2 := len(s), len(t)
	dp := make([][]bool, n1+1)
	for i := range dp {
		dp[i] = make([]bool, n2+1)
	}
	for j := 0; j <= n2; j++ {
		dp[0][j] = true
	}
	if s == "" {
		return true
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1]
				if i == n1 && dp[i][j] {
					return true
				}
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)
