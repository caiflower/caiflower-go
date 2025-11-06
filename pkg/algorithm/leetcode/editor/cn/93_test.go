package leetcode

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
//
//
// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312"
//和 "192.168@1.1" 是 无效 IP 地址。
//
//
// 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你 不能 重新
//排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。
//
//
//
// 示例 1：
//
//
//输入：s = "25525511135"
//输出：["255.255.11.135","255.255.111.35"]
//
//
// 示例 2：
//
//
//输入：s = "0000"
//输出：["0.0.0.0"]
//
//
// 示例 3：
//
//
//输入：s = "101023"
//输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 20
// s 仅由数字组成
//
//
// Related Topics 字符串 回溯 👍 1548 👎 0

func TestRestoreIpAddresses(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected []string
	}

	cases := []testCase{
		{
			name:     "normal case",
			input:    "25525511135",
			expected: []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			name:     "all zeros",
			input:    "0000",
			expected: []string{"0.0.0.0"},
		},
		{
			name:     "single digits",
			input:    "1111",
			expected: []string{"1.1.1.1"},
		},
		{
			name:     "with leading zeros",
			input:    "010010",
			expected: []string{"0.10.0.10", "0.100.1.0"},
		},
		{
			name:     "multiple combinations",
			input:    "101023",
			expected: []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"},
		},
		{
			name:     "invalid with 256",
			input:    "25625511135",
			expected: []string{},
		},
		{
			name:     "empty string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "too short",
			input:    "123",
			expected: []string{},
		},
		{
			name:     "too long",
			input:    "1234567890123",
			expected: []string{},
		},
		{
			name:     "max valid numbers",
			input:    "255255255255",
			expected: []string{"255.255.255.255"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := restoreIpAddresses(tc.input)
			assert.ElementsMatch(t, tc.expected, actual, "Input: %s", tc.input)
		})
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
func restoreIpAddresses(s string) []string {
	ans := make([]string, 0)
	cur := make([]byte, 0)

	dfsRestoreIpAddresses(&ans, s, cur)

	return ans
}

func dfsRestoreIpAddresses(ans *[]string, s string, cur []byte) {
	if s == "" {
		t := string(cur)
		if len(strings.Split(t, ".")) != 4 {
			return
		}

		*ans = append(*ans, t)
		return
	}
	if len(strings.Split(string(cur), ".")) > 4 {
		return
	}

	n := len(s)
	var num []byte
	for i := 0; i < n; i++ {
		num = append(num, s[i])
		if !isValidNum(num) {
			break
		}

		next := cur
		if len(next) != 0 {
			next = append(next, '.')
		}
		next = append(next, num...)

		dfsRestoreIpAddresses(ans, s[i+1:], next)
	}
}

func isValidNum(num []byte) bool {
	if len(num) == 0 || len(num) > 3 {
		return false
	}

	if num[0] == '0' && len(num) != 1 {
		return false
	}

	val, _ := strconv.Atoi(string(num))
	if val < 0 || val > 255 {
		return false
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)
