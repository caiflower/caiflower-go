package leetcode

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//给你一个字符串 s，请你将 s 分割成一些 子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
//
//
//
// 示例 1：
//
//
//输入：s = "aab"
//输出：[["a","a","b"],["aa","b"]]
//
//
// 示例 2：
//
//
//输入：s = "a"
//输出：[["a"]]
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 16
// s 仅由小写英文字母组成
//
//
// Related Topics 字符串 动态规划 回溯 👍 2118 👎 0

func Test131(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected [][]string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: [][]string{{}},
		},
		{
			name:     "single character",
			input:    "a",
			expected: [][]string{{"a"}},
		},
		{
			name:  "multiple partitions",
			input: "aab",
			expected: [][]string{
				{"a", "a", "b"},
				{"aa", "b"},
			},
		},
		{
			name:  "entire string is palindrome",
			input: "aaa",
			expected: [][]string{
				{"a", "a", "a"},
				{"a", "aa"},
				{"aa", "a"},
				{"aaa"},
			},
		},
		{
			name:     "no palindrome partition",
			input:    "abc",
			expected: [][]string{{"a", "b", "c"}},
		},
		{
			name:     "efe",
			input:    "efe",
			expected: [][]string{{"e", "f", "e"}, {"efe"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := partition(tt.input)
			// 对结果和预期进行排序，确保顺序不影响比较
			sortResults(result)
			sortResults(tt.expected)
			assert.Equal(t, tt.expected, result, "Partition(%q) result mismatch", tt.input)
		})
	}
}

// 辅助函数：对结果进行标准化排序
func sortResults(results [][]string) {
	// 先对每个子切片排序
	for _, res := range results {
		sort.Strings(res)
	}
	// 再对整个切片排序
	sort.Slice(results, func(i, j int) bool {
		return strings.Join(results[i], "") < strings.Join(results[j], "")
	})
}

// leetcode submit region begin(Prohibit modification and deletion)
func partition(s string) [][]string {
	ans := make([][]string, 0)
	dfsPartition(&ans, []string{}, s)
	return ans
}

func dfsPartition(ans *[][]string, cur []string, s string) {
	if s == "" {
		t := make([]string, len(cur))
		copy(t, cur)
		*ans = append(*ans, t)
		return
	}

	for i := 0; i < len(s); i++ {
		index := i + 1
		if !isCursor(s[0:index]) {
			continue
		}
		cur = append(cur, s[0:index])
		dfsPartition(ans, cur, s[index:])
		cur = cur[:len(cur)-1]
	}
}

func isCursor(s string) bool {
	i, j := 0, len(s)-1

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)
