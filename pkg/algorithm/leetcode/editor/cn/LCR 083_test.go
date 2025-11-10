package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给定一个不含重复数字的整数数组 nums ，返回其 所有可能的全排列 。可以 按任意顺序 返回答案。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
// 示例 2：
//
//
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
//
//
// 示例 3：
//
//
//输入：nums = [1]
//输出：[[1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// nums 中的所有整数 互不相同
//
//
//
//
//
// 注意：本题与主站 46 题相同：https://leetcode-cn.com/problems/permutations/
//
// Related Topics 数组 回溯 👍 88 👎 0

func TestLCR_083(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{name: "示例1", nums: []int{1, 2, 3}, want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		{name: "示例2", nums: []int{0, 1}, want: [][]int{{0, 1}, {1, 0}}},
		{name: "示例3", nums: []int{1}, want: [][]int{{1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permute(tt.nums)
			// 由于排列顺序不影响结果，需要进行特殊比较
			assert.ElementsMatch(t, tt.want, got, "permute() 结果不符合预期")
		})
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	ans := [][]int{}
	n := len(nums)

	used := make([]bool, n)

	var dfs func(cur []int)
	dfs = func(cur []int) {
		if len(cur) == n {
			t := make([]int, n)
			copy(t, cur)
			ans = append(ans, t)
		}

		for i := 0; i < n; i++ {
			if !used[i] {
				used[i] = true
				dfs(append(cur, nums[i]))
				used[i] = false
			}
		}
	}

	dfs([]int{})

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
