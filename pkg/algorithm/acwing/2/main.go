package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 01背包问题，要么要，要么不要

/**
有N件物品和一个容量是V的背包。每件物品只能使用一次。
第 i件物品的体积是 vi，价值是 wi。
求解将哪些物品装入背包，可使这些物品的总体积不超过背包容量，且总价值最大。
输出最大价值。

输入格式
第一行两个整数，N，V，用空格隔开，分别表示物品数量和背包容积。接下来有 N 行，每行两个整数 vi,wi，用空格隔开，分别表示第 i 件物品的体积和价值。

输入样例
4 5
1 2
2 4
3 4
4 5
输出样例：
8
*/

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	text := s.Text()
	fields := strings.Fields(text)
	n, _ := strconv.Atoi(fields[0])
	v, _ := strconv.Atoi(fields[1])
	goods := make([][2]int, n)
	for i := 0; i < n; i++ {
		s.Scan()
		text = s.Text()
		fields = strings.Fields(text)
		v1, _ := strconv.Atoi(fields[0])
		v2, _ := strconv.Atoi(fields[1])
		goods[i] = [2]int{v1, v2}
	}

	dp := make([]int, v+1)

	for i := 0; i < n; i++ {
		for j := v; j >= goods[i][0]; j-- {
			dp[j] = max(dp[j], dp[j-goods[i][0]]+goods[i][1])
		}
	}

	fmt.Println(dp[v])
}
