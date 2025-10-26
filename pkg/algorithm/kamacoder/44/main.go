package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
题目描述
在一个城市区域内，被划分成了n * m个连续的区块，每个区块都拥有不同的权值，代表着其土地价值。目前，有两家开发公司，A 公司和 B 公司，希望购买这个城市区域的土地。

现在，需要将这个城市区域的所有区块分配给 A 公司和 B 公司。

然而，由于城市规划的限制，只允许将区域按横向或纵向划分成两个子区域，而且每个子区域都必须包含一个或多个区块。 为了确保公平竞争，你需要找到一种分配方式，使得 A 公司和 B 公司各自的子区域内的土地总价值之差最小。

注意：区块不可再分。

输入描述
第一行输入两个正整数，代表 n 和 m。

接下来的 n 行，每行输出 m 个正整数。

输出描述
请输出一个整数，代表两个子区域内土地总价值之间的最小差距。
输入示例
3 3
1 2 3
2 1 3
1 2 3
输出示例
0
*/

// 使用前缀和
func main() {
	// 读取数据
	var n, m int
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	text := s.Text()
	fields := strings.Fields(text)
	n, _ = strconv.Atoi(fields[0])
	m, _ = strconv.Atoi(fields[1])

	lands := make([][]int, n)

	for i := 0; i < n; i++ {
		lands[i] = make([]int, m)
		s.Scan()
		fields := strings.Fields(s.Text())
		for j := 0; j < m; j++ {
			lands[i][j], _ = strconv.Atoi(fields[j])
		}
	}

	// 计算前缀和
	rows := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum := 0
		for j := 0; j < m; j++ {
			sum += lands[i-1][j]
		}
		rows[i] = rows[i-1] + sum
	}

	cols := make([]int, m+1)
	for i := 1; i <= m; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			sum += lands[j][i-1]
		}
		cols[i] = cols[i-1] + sum
	}

	ans := 1<<63 - 1
	for i := 1; i < n; i++ {
		ans = minInt(int(math.Abs(float64(rows[n]-2*rows[i]))), ans)
	}

	for i := 1; i < m; i++ {
		ans = minInt(int(math.Abs(float64(cols[m]-2*cols[i]))), ans)
	}

	fmt.Println(ans)
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
