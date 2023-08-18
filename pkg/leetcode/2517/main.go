package main

import (
	"fmt"
	"sort"
)

/**
给你一个正整数数组 price ，其中 price[i] 表示第 i 类糖果的价格，另给你一个正整数 k 。

商店组合 k 类 不同 糖果打包成礼盒出售。礼盒的 甜蜜度 是礼盒中任意两种糖果 价格 绝对差的最小值。

返回礼盒的 最大 甜蜜度。

2 <= k <= price.length <= 10^5
1 <= price[i] <= 10^9

[13,5,1,8,21,2]
3
ans:
8
*/

func main() {
	fmt.Printf("%v\n", maximumTastiness([]int{3, 1, 1}, 2))
}

func maximumTastiness(price []int, k int) (res int) {
	sort.Ints(price)
	i, j := 0, price[len(price)-1]-price[0]/(k-1)
	for i <= j {
		m := (j-i)/2 + i

		num := k - 1
		p := price[0]
		for index := 1; index < len(price); index++ {
			if price[index]-p >= m {
				p = price[index]
				num--
			}
			if num == 0 {
				break
			}
		}
		if num == 0 {
			res = m
			i = m + 1
		} else {
			j = m - 1
		}
	}
	return
}

//func abs(v1, v2 int) int {
//	if v1 > v2 {
//		return v1 - v2
//	} else {
//		return v2 - v1
//	}
//}

//func min(v1, v2 int) int {
//	if v1 > v2 {
//		return v2
//	} else {
//		return v1
//	}
//}

//func max(v1, v2 int) int {
//	if v1 > v2 {
//		return v1
//	} else {
//		return v2
//	}
//}
