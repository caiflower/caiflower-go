package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给你一个下标从 0 开始的 m x n 整数矩阵 grid 和一个整数 k 。你从起点 (0, 0) 出发，每一步只能往 下 或者往 右 ，你想要到达终点
//(m - 1, n - 1) 。
//
// 请你返回路径和能被 k 整除的路径数目，由于答案可能很大，返回答案对 10⁹ + 7 取余 的结果。
//
//
//
// 示例 1：
//
//
//
// 输入：grid = [[5,2,4],[3,0,5],[0,7,2]], k = 3
//输出：2
//解释：有两条路径满足路径上元素的和能被 k 整除。
//第一条路径为上图中用红色标注的路径，和为 5 + 2 + 4 + 5 + 2 = 18 ，能被 3 整除。
//第二条路径为上图中用蓝色标注的路径，和为 5 + 3 + 0 + 5 + 2 = 15 ，能被 3 整除。
//
//
// 示例 2：
// 输入：grid = [[0,0]], k = 5
//输出：1
//解释：红色标注的路径和为 0 + 0 = 0 ，能被 5 整除。
//
//
// 示例 3：
// 输入：grid = [[7,3,4,9],[2,3,6,2],[2,3,7,0]], k = 1
//输出：10
//解释：每个数字都能被 1 整除，所以每一条路径的和都能被 k 整除。
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 5 * 10⁴
// 1 <= m * n <= 5 * 10⁴
// 0 <= grid[i][j] <= 100
// 1 <= k <= 50
//
//
// Related Topics 数组 动态规划 矩阵 👍 82 👎 0

func Test2435(t *testing.T) {
	type args struct {
		grid [][]int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例1",
			args: args{
				grid: [][]int{{5, 2, 4}, {3, 0, 5}, {0, 7, 2}},
				k:    3,
			},
			want: 2,
		},
		{
			name: "示例2",
			args: args{
				grid: [][]int{{0, 0}},
				k:    5,
			},
			want: 1,
		},
		{
			name: "示例3",
			args: args{
				grid: [][]int{{7, 3, 4, 9}, {2, 3, 6, 2}, {2, 3, 7, 0}},
				k:    1,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numberOfPaths1(tt.args.grid, tt.args.k)
			assert.Equal(t, tt.want, got)
			got = numberOfPaths(tt.args.grid, tt.args.k)
			assert.Equal(t, tt.want, got)
		})
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
// 暴力，因为结果很多，肯定会超时
func numberOfPaths1(grid [][]int, k int) int {
	paths := [][]int{{1, 0}, {0, 1}}
	n, m := len(grid), len(grid[0])
	ans := 0
	mod := int(1e9 + 7)

	var dfs func(i, j, sum int)

	dfs = func(i, j, sum int) {
		if i >= n || j >= m {
			return
		}
		sum += grid[i][j]
		if i == n-1 && j == m-1 {
			if sum%k == 0 {
				ans = (ans % mod) + 1
			}
			return
		}

		for k := 0; k < len(paths); k++ {
			dfs(i+paths[k][0], j+paths[k][1], sum)
		}
	}

	dfs(0, 0, 0)

	return ans
}

// 动态规划, O(m * n * k)
// 做出来了，但是速度太慢了
func numberOfPaths(grid [][]int, k int) int {
	n, m := len(grid), len(grid[0])
	mod := int(1e9 + 7)

	// 定义dp,dp[i][j][y]表示到达i,j， %k余下的数的次数，答案就是dp[n-1][m-1][0]
	dp := make([][][]int, n)
	for i, _ := range dp {
		dp[i] = make([][]int, m)
		for j := range dp[i] {
			dp[i][j] = make([]int, k)
		}
	}
	dp[0][0][grid[0][0]%k] = 1

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for y := 0; y < k; y++ {
				if i == 0 && j == 0 {
					continue
				}
				if i-1 >= 0 {
					dp[i][j][(grid[i][j]+y)%k] = (dp[i][j][(grid[i][j]+y)%k] + dp[i-1][j][y]) % mod
				}
				if j-1 >= 0 {
					dp[i][j][(grid[i][j]+y)%k] = (dp[i][j-1][y] + dp[i][j][(grid[i][j]+y)%k]) % mod
				}
			}
		}
	}

	return dp[n-1][m-1][0]
}

//leetcode submit region end(Prohibit modification and deletion)
