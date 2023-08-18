package main

import (
	"fmt"
)

/**
一个由字母和数字组成的字符串的 值定义如下：

如果字符串 只 包含数字，那么值为该字符串在 10进制下的所表示的数字。
否则，值为字符串的 长度。
给你一个字符串数组strs，每个字符串都只由字母和数字组成，请你返回 strs中字符串的 最大值

strs = ["alic3","bob","3","4","00000"]
5

["5232","yv","22","c","yawgs","928","4003","2"]

*/

func main() {
	fmt.Printf("%v\n", maximumValue([]string{"5232", "yv", "22", "c", "yawgs", "928", "4003", "2"}))
}

func maximumValue(strs []string) (res int) {
	for _, v := range strs {
		val := 0
		for _, c := range v {
			if c < '0' || c > '9' {
				val = len(v)
				break
			} else {
				if c == '0' && val == 0 {
					continue
				}
				val = val*10 + int(c-'0')
			}
		}
		res = max(res, val)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
