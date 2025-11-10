package leetcode

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
// [1,2,1],
// [2,1,1]]
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 8
// -10 <= nums[i] <= 10
//
//
// Related Topics 数组 回溯 排序 👍 1746 👎 0

func Test47(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "示例1",
			nums: []int{1, 1, 2},
			want: [][]int{
				{1, 1, 2},
				{1, 2, 1},
				{2, 1, 1},
			},
		},
		{
			name: "示例2",
			nums: []int{1, 2, 3},
			want: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, permuteUnique(tt.nums), tt.name)
		})
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
	ans := [][]int{}
	n := len(nums)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	used := make([]bool, n)

	var dfs func(cur []int)
	dfs = func(cur []int) {
		if len(cur) == n {
			t := make([]int, n)
			copy(t, cur)
			ans = append(ans, t)
		}

		hasSet := make(map[int]bool)
		for i := 0; i < n; i++ {
			if !used[i] && !hasSet[nums[i]] {
				used[i] = true
				hasSet[nums[i]] = true
				dfs(append(cur, nums[i]))
				used[i] = false
			}
		}
	}

	dfs([]int{})

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
