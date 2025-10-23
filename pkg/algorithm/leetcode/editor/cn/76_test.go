package leetcode

import (
	"fmt"
	"testing"
)

//给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
//
//
// 注意：
//
//
// 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
// 如果 s 中存在这样的子串，我们保证它是唯一的答案。
//
//
//
//
// 示例 1：
//
//
//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"
//解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
//
//
// 示例 2：
//
//
//输入：s = "a", t = "a"
//输出："a"
//解释：整个字符串 s 是最小覆盖子串。
//
//
// 示例 3:
//
//
//输入: s = "a", t = "aa"
//输出: ""
//解释: t 中两个字符 'a' 均应包含在 s 的子串中，
//因此没有符合条件的子字符串，返回空字符串。
//
//
//
// 提示：
//
//
// m == s.length
// n == t.length
// 1 <= m, n <= 10⁵
// s 和 t 由英文字母组成
//
//
//
//进阶：你能设计一个在
//o(m+n) 时间内解决此问题的算法吗？
//
// Related Topics 哈希表 字符串 滑动窗口 👍 3419 👎 0

func Test76(t *testing.T) {
	type testCase struct {
		s, t     string
		expected string
	}
	cases := []testCase{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
		{"bba", "ab", "ba"},
		{"aa", "aa", "aa"},
		{"a", "b", ""},
		{"a", "", ""},
		{"ab", "a", "a"},
		{"cabwefgewcwaefgcf", "cae", "cwae"},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d: s=%q, t=%q", i+1, tc.s, tc.t), func(t *testing.T) {
			result := minWindow(tc.s, tc.t)
			if result != tc.expected {
				t.Errorf("输出: %q, 期望: %q", result, tc.expected)
			}
		})
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
// o(m+n) 时间复杂度
func minWindow(s string, t string) string {
	ans := ""
	cnts := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		c := t[i]
		cnts[c]++
	}
	l, r, curLen := 0, 0, len(t)

	for l <= r {
		if curLen == 0 {
			if ans == "" || len(s[l:r]) < len(ans) {
				ans = s[l:r]
			}
			c := s[l]
			if _, ok := cnts[c]; ok {
				cnts[c]++
				if cnts[c] > 0 {
					curLen++
				}
			}
			l++
		} else if r < len(s) {
			c := s[r]
			if _, ok := cnts[c]; ok {
				cnts[c]--
				if cnts[c] >= 0 {
					curLen--
				}
			}
			r++
		} else {
			break
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
