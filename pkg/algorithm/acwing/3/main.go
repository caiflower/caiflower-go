package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 完全背包问题，和01背包问题就是将for循环换个方向即可，让一件物品能被多次装

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
		for j := goods[i][0]; j <= v; j++ {
			dp[j] = maxInt(dp[j], dp[j-goods[i][0]]+goods[i][1])
		}
	}

	fmt.Println(dp[v])
}

func maxInt(v1, v2 int) int {
	if v1 > v2 {
		return v1
	} else {
		return v2
	}
}
