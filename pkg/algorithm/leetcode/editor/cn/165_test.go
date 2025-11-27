package leetcode

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//给你两个 版本号字符串 version1 和 version2 ，请你比较它们。版本号由被点 '.' 分开的修订号组成。修订号的值 是它 转换为整数 并忽略
//前导零。
//
// 比较版本号时，请按 从左到右的顺序 依次比较它们的修订号。如果其中一个版本字符串的修订号较少，则将缺失的修订号视为 0。
//
// 返回规则如下：
//
//
// 如果 version1 < version2 返回 -1，
// 如果 version1 > version2 返回 1，
// 除此之外返回 0。
//
//
//
//
// 示例 1：
//
//
// 输入：version1 = "1.2", version2 = "1.10"
//
//
// 输出：-1
//
// 解释：
//
// version1 的第二个修订号为 "2"，version2 的第二个修订号为 "10"：2 < 10，所以 version1 < version2。
//
// 示例 2：
//
//
// 输入：version1 = "1.01", version2 = "1.001"
//
//
// 输出：0
//
// 解释：
//
// 忽略前导零，"01" 和 "001" 都代表相同的整数 "1"。
//
// 示例 3：
//
//
// 输入：version1 = "1.0", version2 = "1.0.0.0"
//
//
// 输出：0
//
// 解释：
//
// version1 有更少的修订号，每个缺失的修订号按 "0" 处理。
//
//
//
// 提示：
//
//
// 1 <= version1.length, version2.length <= 500
// version1 和 version2 仅包含数字和 '.'
// version1 和 version2 都是 有效版本号
// version1 和 version2 的所有修订号都可以存储在 32 位整数 中
//
//
// Related Topics 双指针 字符串 👍 486 👎 0

func Test165(t *testing.T) {
	tests := []struct {
		name     string
		version1 string
		version2 string
		expected int
	}{
		{"case1", "1.2", "1.10", -1},
		{"case2", "1.01", "1.001", 0},
		{"case3", "1.0", "1.0.0.0", 0},
		{"case4", "0.1", "1.1", -1},
		{"case5", "1.0.1", "1", 1},
		{"case6", "7.5.2.4", "7.5.3", -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := compareVersion(test.version1, test.version2)
			assert.Equal(t, test.expected, result)
		})
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
func compareVersion(version1, version2 string) int {
	fields1 := strings.Split(version1, ".")
	fields2 := strings.Split(version2, ".")

	i := 0
	for i = 0; i < len(fields1) && i < len(fields2); i++ {
		v1, _ := strconv.Atoi(fields1[i])
		v2, _ := strconv.Atoi(fields2[i])
		if v1 > v2 {
			return 1
		} else if v1 < v2 {
			return -1
		}
	}

	if i < len(fields1) {
		for j := i; j < len(fields1); j++ {
			v, _ := strconv.Atoi(fields1[j])
			if v > 0 {
				return 1
			}
		}
	} else if i < len(fields2) {
		for j := i; j < len(fields2); j++ {
			v, _ := strconv.Atoi(fields2[j])
			if v > 0 {
				return -1
			}
		}
	}

	return 0
}

//leetcode submit region end(Prohibit modification and deletion)
