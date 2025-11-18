package leetcode

import (
	"testing"
)

//给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数 。
//
// 你可以对一个单词进行如下三种操作：
//
//
// 插入一个字符
// 删除一个字符
// 替换一个字符
//
//
//
//
// 示例 1：
//
//
//输入：word1 = "horse", word2 = "ros"
//输出：3
//解释：
//horse -> rorse (将 'h' 替换为 'r')
//rorse -> rose (删除 'r')
//rose -> ros (删除 'e')
//
//
// 示例 2：
//
//
//输入：word1 = "intention", word2 = "execution"
//输出：5
//解释：
//intention -> inention (删除 't')
//inention -> enention (将 'i' 替换为 'e')
//enention -> exention (将 'n' 替换为 'x')
//exention -> exection (将 'n' 替换为 'c')
//exection -> execution (插入 'u')
//
//
//
//
// 提示：
//
//
// 0 <= word1.length, word2.length <= 500
// word1 和 word2 由小写英文字母组成
//
//
// Related Topics 字符串 动态规划 👍 3800 👎 0

func Test72(t *testing.T) {
	tests := []struct {
		name   string
		word1  string
		word2  string
		expect int
	}{
		{
			name:   "示例1",
			word1:  "horse",
			word2:  "ros",
			expect: 3,
		},
		{
			name:   "示例2",
			word1:  "intention",
			word2:  "execution",
			expect: 5,
		},
		{
			name:   "空字符串测试",
			word1:  "",
			word2:  "abc",
			expect: 3,
		},
		{
			name:   "空字符串测试2",
			word1:  "abc",
			word2:  "",
			expect: 3,
		},
		{
			name:   "相同字符串",
			word1:  "abc",
			word2:  "abc",
			expect: 0,
		},
		{
			name:   "一个字符差异",
			word1:  "abc",
			word2:  "abd",
			expect: 1,
		},
		{
			name:   "wrong case1",
			word1:  "sea",
			word2:  "eat",
			expect: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minDistance1(tt.word1, tt.word2)
			if got != tt.expect {
				t.Errorf("minDistance(%s, %s) = %d, 期望 %d", tt.word1, tt.word2, got, tt.expect)
			}
		})
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
func minDistance1(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)

	// dp[i][j] = word1前i个字符转化成word2前i个字符最小的步骤数
	dp := make([][]int, n1+1)
	for i := range dp {
		t := make([]int, n2+1)
		t[0] = i
		dp[i] = t
	}
	for j := 0; j <= n2; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// dp[i][j-1]+1 表示word1插入一个字符或者word2删除一个字符
				// dp[i-1][j]+1 表示word1删除一个字符或者word2插入一个字符
				// dp[i-1][j-1]+1 表示word1替换一个字符或者word2替换一个字符
				dp[i][j] = min(min(dp[i][j-1]+1, dp[i-1][j]+1), dp[i-1][j-1]+1)
			}
		}
	}

	return dp[n1][n2]
}

//leetcode submit region end(Prohibit modification and deletion)
