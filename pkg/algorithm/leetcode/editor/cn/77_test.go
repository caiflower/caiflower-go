package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
//
// 你可以按 任何顺序 返回答案。
//
//
//
// 示例 1：
//
//
//输入：n = 4, k = 2
//输出：
//[
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
//]
//
// 示例 2：
//
//
//输入：n = 1, k = 1
//输出：[[1]]
//
//
//
// 提示：
//
//
// 1 <= n <= 20
// 1 <= k <= n
//
//
// Related Topics 回溯 👍 1809 👎 0

func Test77(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want [][]int
	}{
		{
			name: "n=4,k=2",
			n:    4,
			k:    2,
			want: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			name: "k=1 (single element combinations)",
			n:    4,
			k:    1,
			want: [][]int{{1}, {2}, {3}, {4}},
		},
		{
			name: "k=n (full set)",
			n:    5,
			k:    5,
			want: [][]int{{1, 2, 3, 4, 5}},
		},
		{
			name: "n=0 (empty input)",
			n:    0,
			k:    0,
			want: [][]int{},
		},
		{
			name: "k=0 (invalid input)",
			n:    4,
			k:    0,
			want: [][]int{},
		},
		{
			name: "k > n (invalid input)",
			n:    3,
			k:    5,
			want: [][]int{},
		},
		{
			name: "n=3,k=2",
			n:    3,
			k:    2,
			want: [][]int{{1, 2}, {1, 3}, {2, 3}},
		},
		{
			name: "n=5,k=3",
			n:    5,
			k:    3,
			want: [][]int{
				{1, 2, 3}, {1, 2, 4}, {1, 2, 5},
				{1, 3, 4}, {1, 3, 5}, {1, 4, 5},
				{2, 3, 4}, {2, 3, 5}, {2, 4, 5},
				{3, 4, 5},
			},
		},
		{
			name: "n=5,k=4",
			n:    5,
			k:    4,
			want: [][]int{
				{1, 2, 3, 4}, {1, 2, 3, 5}, {1, 2, 4, 5},
				{1, 3, 4, 5}, {2, 3, 4, 5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combine(tt.n, tt.k)
			assert.ElementsMatch(t, tt.want, got, "combine(%d, %d) should equal %v", tt.n, tt.k, tt.want)
		})
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	if k == 0 {
		return ans
	}
	travelCombine(1, n, k, []int{}, &ans)
	return ans
}

func travelCombine(s, n, k int, cur []int, ans *[][]int) {
	if k == 0 {
		t := make([]int, len(cur))
		copy(t, cur)
		*ans = append(*ans, t)
		return
	}
	if (n - s + 1) < k {
		return
	}

	travelCombine(s+1, n, k, cur, ans)
	cur = append(cur, s)
	travelCombine(s+1, n, k-1, cur, ans)

	return
}

//leetcode submit region end(Prohibit modification and deletion)
