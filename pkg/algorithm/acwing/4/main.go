package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 多重背包问题
func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	fields := strings.Fields(s.Text())
	n, _ := strconv.Atoi(fields[0])
	v, _ := strconv.Atoi(fields[1])
	goods := make([][3]int, n)
	for i := 0; i < n; i++ {
		s.Scan()
		fields = strings.Fields(s.Text())
		v1, _ := strconv.Atoi(fields[0])
		v2, _ := strconv.Atoi(fields[1])
		v3, _ := strconv.Atoi(fields[2])
		goods[i] = [3]int{v1, v2, v3}
	}

	dp := make([]int, v+1)

	for i := 0; i < n; i++ {
		for j := 1; j <= goods[i][2]; j++ {
			gv := j * goods[i][0]
			for k := v; k >= gv; k-- {
				dp[k] = max(dp[k], dp[k-goods[i][0]]+goods[i][1])
			}
		}
	}

	fmt.Println(dp[v])
}
