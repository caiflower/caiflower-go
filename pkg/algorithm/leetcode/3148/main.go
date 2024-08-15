package main

import (
	"fmt"
	"math"
)

//给你一个由 正整数 组成、大小为 m x n 的矩阵 grid。你可以从矩阵中的任一单元格移动到另一个位于正下方或正右侧的任意单元格（不必相邻）。从值为
//c1 的单元格移动到值为 c2 的单元格的得分为 c2 - c1 。
//
// 你可以从任一单元格开始，并且必须至少移动一次。
//
// 返回你能得到的 最大 总得分。
//
//
//
// 示例 1：
//
//
// 输入：grid = [[9,5,7,3],[8,9,6,1],[6,7,14,3],[2,5,3,1]]
//
//
// 输出：9
//
// 解释：从单元格 (0, 1) 开始，并执行以下移动： - 从单元格 (0, 1) 移动到 (2, 1)，得分为 7 - 5 = 2 。 - 从单元格 (2
//, 1) 移动到 (2, 2)，得分为 14 - 7 = 7 。 总得分为 2 + 7 = 9 。
//
// 示例 2：
//
//
//
//
// 输入：grid = [[4,3,2],[3,2,1]]
//
//
// 输出：-1
//
// 解释：从单元格 (0, 0) 开始，执行一次移动：从 (0, 0) 到 (0, 1) 。得分为 3 - 4 = -1 。
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 2 <= m, n <= 1000
// 4 <= m * n <= 10⁵
// 1 <= grid[i][j] <= 10⁵
//
//
// Related Topics 数组 动态规划 矩阵 👍 34 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

func main() {
	fmt.Println(maxScore([][]int{{4, 3, 2}, {3, 2, 1}}))
	fmt.Println(maxScore([][]int{{9, 5, 7, 3}, {8, 9, 6, 1}, {6, 7, 14, 3}, {2, 5, 3, 1}}))
}

func maxScore(grid [][]int) (res int) {
	res = math.MinInt
	r, c := len(grid), len(grid[0])
	t := make([][]int, r) // 记录当前位置能够获取到的最大分数
	for i, _ := range t {
		t[i] = make([]int, c)
		for j, _ := range t[i] {
			t[i][j] = 0
		}
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			jump := false
			for k := i + 1; k < r; k++ {
				t[k][j], jump = maxInt(t[k][j], grid[k][j]-grid[i][j]+t[i][j])
				if jump {
					res, _ = maxInt(res, t[k][j])
				} else {
					// 如果没有跳跃，也就是跳跃了反而分数小于0，那么根据题目要求要至少跳跃一次，这里强行跳跃
					res, _ = maxInt(res, grid[k][j]-grid[i][j])
				}
			}
			for k := j + 1; k < c; k++ {
				t[i][k], jump = maxInt(t[i][k], grid[i][k]-grid[i][j]+t[i][j])
				if jump {
					res, _ = maxInt(res, t[i][k])
				} else {
					res, _ = maxInt(res, grid[i][k]-grid[i][j])
				}
			}
		}
	}
	return res
}

func maxInt(a, b int) (int, bool) {
	if a > b {
		return a, false
	} else {
		return b, true
	}
}

func minInt(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
