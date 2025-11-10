package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给定一个 m x n 的整数数组 grid。一个机器人初始位于 左上角（即 grid[0][0]）。机器人尝试移动到 右下角（即 grid[m - 1][
//n - 1]）。机器人每次只能向下或者向右移动一步。
//
// 网格中的障碍物和空位置分别用 1 和 0 来表示。机器人的移动路径中不能包含 任何 有障碍物的方格。
//
// 返回机器人能够到达右下角的不同路径数量。
//
// 测试用例保证答案小于等于 2 * 10⁹。
//
//
//
// 示例 1：
//
//
//输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
//输出：2
//解释：3x3 网格的正中间有一个障碍物。
//从左上角到右下角一共有 2 条不同的路径：
//1. 向右 -> 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右 -> 向右
//
//
// 示例 2：
//
//
//输入：obstacleGrid = [[0,1],[0,0]]
//输出：1
//
//
//
//
// 提示：
//
//
// m == obstacleGrid.length
// n == obstacleGrid[i].length
// 1 <= m, n <= 100
// obstacleGrid[i][j] 为 0 或 1
//
//
// Related Topics 数组 动态规划 矩阵 👍 1450 👎 0

func Test63(t *testing.T) {
	tests := []struct {
		name         string
		obstacleGrid [][]int
		want         int
	}{
		{
			name: "示例1",
			obstacleGrid: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			want: 2,
		},
		{
			name: "示例2",
			obstacleGrid: [][]int{
				{0, 1},
				{0, 0},
			},
			want: 1,
		},
		{
			name: "起点有障碍",
			obstacleGrid: [][]int{
				{1, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			want: 0,
		},
		{
			name: "终点有障碍",
			obstacleGrid: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 1},
			},
			want: 0,
		},
		{
			name: "单元格",
			obstacleGrid: [][]int{
				{0},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, uniquePathsWithObstacles(tt.obstacleGrid), tt.name)
		})
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid)
	n1 := len(obstacleGrid[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n1)
	}

	if obstacleGrid[0][0] == 1 {
		return 0
	}
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < n1; j++ {
			if obstacleGrid[i][j] == 0 {
				if i > 0 {
					dp[i][j] += dp[i-1][j]
				}
				if j > 0 {
					dp[i][j] += dp[i][j-1]
				}
			}

		}
	}

	return dp[n-1][n1-1]
}

//leetcode submit region end(Prohibit modification and deletion)
