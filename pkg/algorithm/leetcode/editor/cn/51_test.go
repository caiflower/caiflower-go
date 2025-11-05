package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
//
// n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
// 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
//
//
//
// 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
//
//
//
//
//
// 示例 1：
//
//
//输入：n = 4
//输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
//解释：如上图所示，4 皇后问题存在两个不同的解法。
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：[["Q"]]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 9
//
//
// Related Topics 数组 回溯 👍 2369 👎 0

func TestSolveNQueens(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected [][]string
	}{
		{
			name:     "n=1",
			n:        1,
			expected: [][]string{{"Q"}},
		},
		{
			name: "n=4",
			n:    4,
			expected: [][]string{
				{
					".Q..",
					"...Q",
					"Q...",
					"..Q.",
				},
				{
					"..Q.",
					"Q...",
					"...Q",
					".Q..",
				},
			},
		},
		{
			name:     "n=0",
			n:        0,
			expected: [][]string{},
		},
		{
			name:     "n=2",
			n:        2,
			expected: [][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := solveNQueens(tt.n)
			assert.Len(t, result, len(tt.expected), "Number of solutions mismatch for n=%d", tt.n)

			// Validate each solution is correct
			for _, solution := range result {
				assert.True(t, isValidSolution(solution), "Invalid solution found: %v", solution)
			}

			// Compare results ignoring order
			assert.ElementsMatch(t, tt.expected, result, "Solutions content mismatch for n=%d", tt.n)
		})
	}
}

func isValidSolution(solution []string) bool {
	n := len(solution)
	queens := make([][2]int, 0)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if solution[i][j] == 'Q' {
				queens = append(queens, [2]int{i, j})
			}
		}
	}

	for i := 0; i < len(queens); i++ {
		for j := i + 1; j < len(queens); j++ {
			x1, y1 := queens[i][0], queens[i][1]
			x2, y2 := queens[j][0], queens[j][1]
			if x1 == x2 || y1 == y2 || abs(x1-x2) == abs(y1-y2) {
				return false
			}
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) [][]string {
	ans := make([][]string, 0)
	if n == 0 {
		return ans
	}
	occupy := make([][]int, 0)
	for i := 0; i < 2; i++ {
		occupy = append(occupy, make([]int, n))
	}
	occupy = append(occupy, make([]int, n*2-1))
	occupy = append(occupy, make([]int, n*2-1))

	dfsSolveNQueens(&ans, occupy, []string{}, 0, n)

	return ans
}

func dfsSolveNQueens(ans *[][]string, occupy [][]int, cur []string, i, n int) {
	if i == n {
		t := make([]string, len(cur))
		copy(t, cur)
		*ans = append(*ans, t)
		return
	}

	for j := 0; j < n; j++ {
		if isOccpy(occupy, i, j, n) {
			continue
		}
		toOccupy(occupy, i, j, n, &cur)
		dfsSolveNQueens(ans, occupy, cur, i+1, n)
		deOccupy(occupy, i, j, n, &cur)
	}
}

func isOccpy(occupy [][]int, i, j, n int) bool {
	return occupy[0][i] == 1 || occupy[1][j] == 1 || occupy[2][i-j+n-1] == 1 || occupy[3][i+j] == 1
}

func toOccupy(occupy [][]int, i, j, n int, cur *[]string) {
	occupy[0][i] = 1
	occupy[1][j] = 1
	occupy[2][i-j+n-1] = 1
	occupy[3][j+i] = 1

	row := make([]byte, n)
	for k := range row {
		row[k] = '.'
	}
	row[j] = 'Q'
	t := string(row)
	*cur = append(*cur, t)

}

func deOccupy(occupy [][]int, i, j, n int, cur *[]string) {
	occupy[0][i] = 0
	occupy[1][j] = 0
	occupy[2][i-j+n-1] = 0
	occupy[3][i+j] = 0
	*cur = (*cur)[:len(*cur)-1]
}

//leetcode submit region end(Prohibit modification and deletion)
