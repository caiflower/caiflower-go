package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	fields := strings.Fields(s.Text())
	n, _ := strconv.Atoi(fields[0])
	v, _ := strconv.Atoi(fields[1])
	goods := make([][2]int, 0)

	for i := 0; i < n; i++ {
		// 二进制优化，将多重背包问题简化成01背包问题
		// 原理11个数量的物品可以拆分成1 2 4 4表示，由1 2 4 4(11 - (1 + 2 + 4)) 能组成1-11的任何一个数
		s.Scan()
		fields := strings.Fields(s.Text())
		v1, _ := strconv.Atoi(fields[0])
		v2, _ := strconv.Atoi(fields[1])
		v3, _ := strconv.Atoi(fields[2])

		prev := 0
		j := 1
		for j = 1; j+prev <= v3; j *= 2 {
			goods = append(goods, [2]int{j * v1, j * v2})
			prev += j
		}

		if t := v3 - prev; t > 0 {
			goods = append(goods, [2]int{t * v1, t * v2})
		}
	}

	dp := make([]int, v+1)

	for i := 0; i < len(goods); i++ {
		for j := v; j >= goods[i][0]; j-- {
			dp[j] = max(dp[j], dp[j-goods[i][0]]+goods[i][1])
		}

	}

	fmt.Println(dp[v])
}
