package leetcode

import "testing"

//给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。
//
// 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
//
//
//
// 示例 1：
//
//
//输入: s = "leetcode", wordDict = ["leet", "code"]
//输出: true
//解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
//
//
// 示例 2：
//
//
//输入: s = "applepenapple", wordDict = ["apple", "pen"]
//输出: true
//解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
//     注意，你可以重复使用字典中的单词。
//
//
// 示例 3：
//
//
//输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
//输出: false
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 300
// 1 <= wordDict.length <= 1000
// 1 <= wordDict[i].length <= 20
// s 和 wordDict[i] 仅由小写英文字母组成
// wordDict 中的所有字符串 互不相同
//
//
// Related Topics 字典树 记忆化搜索 数组 哈希表 字符串 动态规划 👍 2861 👎 0

func Test139(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		want     bool
	}{
		{
			name:     "示例1",
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			want:     true,
		},
		{
			name:     "示例2",
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			want:     true,
		},
		{
			name:     "示例3",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			want:     false,
		},
		{
			name:     "空字符串",
			s:        "",
			wordDict: []string{"a"},
			want:     true,
		},
		{
			name:     "单个字符匹配",
			s:        "a",
			wordDict: []string{"a"},
			want:     true,
		},
		{
			name:     "单个字符不匹配",
			s:        "a",
			wordDict: []string{"b"},
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordBreak(tt.s, tt.wordDict)
			if got != tt.want {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
// O(n^3)
func wordBreak(s string, wordDict []string) bool {
	if s == "" {
		return true
	}

	wordMap := make(map[string]bool)

	for _, v := range wordDict {
		wordMap[v] = true
	}

	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		t := make([]bool, n)
		dp[i] = t
	}

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			for k := i; k <= j; k++ {
				if k == j {
					dp[i][j] = wordMap[s[i:j+1]]
				} else {
					// 这里写错过状态转移方程!!!  dp[i][j] = (dp[i][k] && dp[k][j]) || dp[i][j]
					dp[i][j] = dp[i][k] && dp[k+1][j]
					if dp[i][j] {
						break
					}
				}
			}
		}
	}

	return dp[0][n-1]
}

//leetcode submit region end(Prohibit modification and deletion)
