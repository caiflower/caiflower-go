package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：
//
//
// 只使用数字1到9
// 每个数字 最多使用一次
//
//
// 返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。
//
//
//
// 示例 1:
//
//
//输入: k = 3, n = 7
//输出: [[1,2,4]]
//解释:
//1 + 2 + 4 = 7
//没有其他符合的组合了。
//
// 示例 2:
//
//
//输入: k = 3, n = 9
//输出: [[1,2,6], [1,3,5], [2,3,4]]
//解释:
//1 + 2 + 6 = 9
//1 + 3 + 5 = 9
//2 + 3 + 4 = 9
//没有其他符合的组合了。
//
// 示例 3:
//
//
//输入: k = 4, n = 1
//输出: []
//解释: 不存在有效的组合。
//在[1,9]范围内使用4个不同的数字，我们可以得到的最小和是1+2+3+4 = 10，因为10 > 1，没有有效的组合。
//
//
//
//
// 提示:
//
//
// 2 <= k <= 9
// 1 <= n <= 60
//
//
// Related Topics 数组 回溯 👍 958 👎 0

func TestCombinationSum3(t *testing.T) {
	tests := []struct {
		name string
		k    int
		n    int
		want [][]int
	}{
		{
			name: "k=3,n=7",
			k:    3,
			n:    7,
			want: [][]int{{1, 2, 4}},
		},
		{
			name: "k=3,n=9",
			k:    3,
			n:    9,
			want: [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}},
		},
		{
			name: "k=1,n=5",
			k:    1,
			n:    5,
			want: [][]int{{5}},
		},
		{
			name: "k=2,n=1 (no solution)",
			k:    2,
			n:    1,
			want: [][]int{},
		},
		{
			name: "k=0,n=0 (invalid input)",
			k:    0,
			n:    0,
			want: [][]int{},
		},
		{
			name: "k=9,n=45 (max sum)",
			k:    9,
			n:    45,
			want: [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		},
		{
			name: "k=4,n=1 (invalid combination)",
			k:    4,
			n:    1,
			want: [][]int{},
		},
		{
			name: "k=2,n=17 (edge case)",
			k:    2,
			n:    17,
			want: [][]int{{8, 9}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum3(tt.k, tt.n)
			assert.ElementsMatch(t, tt.want, got, "combinationSum3(%d, %d) should equal %v", tt.k, tt.n, tt.want)
		})
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
var (
	cur []int
	sum int
	ans [][]int
)

func combinationSum3(k int, n int) [][]int {
	cur = make([]int, 0)
	sum = 0
	ans = make([][]int, 0)
	if n == 0 {
		return ans
	}

	dfsCombinationSum3(1, k, n)
	return ans
}

func dfsCombinationSum3(num, k, n int) {
	if sum == n && k == 0 {
		t := make([]int, len(cur))
		copy(t, cur)
		ans = append(ans, t)
		return
	}

	if num > 9 || sum >= n || 9-num+1 < k || k*num+sum > n {
		return
	}

	cur = append(cur, num)
	sum += num
	dfsCombinationSum3(num+1, k-1, n)
	cur = cur[:len(cur)-1]
	sum -= num
	dfsCombinationSum3(num+1, k, n)
}

//leetcode submit region end(Prohibit modification and deletion)
