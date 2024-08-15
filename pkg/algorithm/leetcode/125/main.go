package main

import (
	"fmt"
	"strings"
)

//如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
//
// 字母和数字都属于字母数字字符。
//
// 给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
//输入: s = "A man, a plan, a canal: Panama"
//输出：true
//解释："amanaplanacanalpanama" 是回文串。
//
//
// 示例 2：
//
//
//输入：s = "race a car"
//输出：false
//解释："raceacar" 不是回文串。
//
//
// 示例 3：
//
//
//输入：s = " "
//输出：true
//解释：在移除非字母数字字符之后，s 是一个空字符串 "" 。
//由于空字符串正着反着读都一样，所以是回文串。
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 2 * 10⁵
// s 仅由可打印的 ASCII 字符组成
//
//
// Related Topics 双指针 字符串 👍 756 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		for i < j && (s[i] < 'a' || s[i] > 'z') && (s[i] < '0' || s[i] > '9') {
			i++
		}
		for i < j && (s[j] < 'a' || s[j] > 'z') && (s[j] < '0' || s[j] > '9') {
			j--
		}
		if i >= j {
			break
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	//fmt.Println(isPalindrome("aab"))
	//fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("ab2a"))
	//fmt.Println(isPalindrome("amanaplanacanalpanama"))
}
