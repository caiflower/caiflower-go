package leetcode

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//
//
//
//
// 示例 1：
//
//
//输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//
//
// 示例 2：
//
//
//输入：digits = "2"
//输出：["a","b","c"]
//
//
//
//
// 提示：
//
//
// 1 <= digits.length <= 4
// digits[i] 是范围 ['2', '9'] 的一个数字。
//
//
// Related Topics 哈希表 字符串 回溯 👍 3144 👎 0

func Test17(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "empty input",
			input:    "",
			expected: []string{},
		},
		{
			name:     "single digit 2",
			input:    "2",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "two digits 23",
			input:    "23",
			expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:     "digits with 4 letters",
			input:    "79",
			expected: []string{"pw", "px", "py", "pz", "qw", "qx", "qy", "qz", "rw", "rx", "ry", "rz", "sw", "sx", "sy", "sz"},
		},
		{
			name:     "input with 0",
			input:    "0",
			expected: []string{},
		},
		{
			name:     "input with 1",
			input:    "1",
			expected: []string{},
		},
		{
			name:     "mixed invalid and valid digits",
			input:    "020",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "multiple invalid digits",
			input:    "010",
			expected: []string{},
		},
		{
			name:     "three valid digits",
			input:    "234",
			expected: []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := letterCombinations(tt.input)
			assert.ElementsMatch(t, tt.expected, result, "Input: %s, Expected: %v, Got: %v", tt.input, tt.expected, result)
		})
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
var m = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	ans := make([]string, 0)
	dfsLetterCombinations(&ans, "", 0, digits)
	return ans
}

func dfsLetterCombinations(ans *[]string, curString string, cur int, digits string) {
	if cur >= len(digits) {
		if curString != "" {
			*ans = append(*ans, curString)
		}
		return
	}
	index, _ := strconv.Atoi(string(digits[cur]))
	str := m[index]

	if str == "" {
		dfsLetterCombinations(ans, curString, cur+1, digits)
	} else {
		for _, v := range str {
			dfsLetterCombinations(ans, curString+string(v), cur+1, digits)
		}
	}

	return
}

//leetcode submit region end(Prohibit modification and deletion)
