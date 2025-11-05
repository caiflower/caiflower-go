package leetcode

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

//给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是回文串。
//
// 返回符合要求的 最少分割次数 。
//
//
//
//
//
//
//
// 示例 1：
//
//
//输入：s = "aab"
//输出：1
//解释：只需一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。
//
//
// 示例 2：
//
//
//输入：s = "a"
//输出：0
//
//
// 示例 3：
//
//
//输入：s = "ab"
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 2000
// s 仅由小写英文字母组成
//
//
// Related Topics 字符串 动态规划 👍 843 👎 0

func Test132(t *testing.T) {
	tests := []struct {
		name string
		want int
		cast string
	}{
		{name: "", want: 0, cast: ""},
		{name: "aab", want: 1, cast: "aab"},
		{name: "a", want: 0, cast: "a"},
		{name: "ab", want: 1, cast: "ab"},
		{name: "abababa", want: 0, cast: "abababa"},
		{name: "ababab", want: 1, cast: "ababab"},
		{name: "cdd", want: 1, cast: "cdd"},
		{name: "cabababcbc", want: 3, cast: "cabababcbc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, minCut(tt.cast), tt.name)
		})
	}

}

// leetcode submit region begin(Prohibit modification and deletion)
func minCut(s string) int {
	n := len(s)
	if s == "" {
		return 0
	}
	cursor := make([][]bool, 0)
	for i := 0; i < len(s); i++ {
		t := make([]bool, len(s))
		t[i] = true
		cursor = append(cursor, t)
	}

	for i := 0; i < len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if (j+1 > i-1 || cursor[j+1][i-1] == true) && s[i] == s[j] {
				cursor[j][i] = true
			}
		}
	}

	// 这里是难点
	f := make([]int, n)
	for i := range f {
		if cursor[0][i] {
			continue
		}
		f[i] = math.MaxInt64
		for j := 0; j < i; j++ {
			if cursor[j+1][i] && f[j]+1 < f[i] {
				f[i] = f[j] + 1
			}
		}
	}
	return f[n-1]
}

//leetcode submit region end(Prohibit modification and deletion)
