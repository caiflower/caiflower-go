package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。
//
// 连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l <= i < r，都有 nums[i] < nums[i + 1] ，那
//么子序列 [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] 就是连续递增子序列。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,3,5,4,7]
//输出：3
//解释：最长连续递增序列是 [1,3,5], 长度为3。
//尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为 5 和 7 在原数组里被 4 隔开。
//
//
// 示例 2：
//
//
//输入：nums = [2,2,2,2,2]
//输出：1
//解释：最长连续递增序列是 [2], 长度为1。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁴
// -10⁹ <= nums[i] <= 10⁹
//
//
// Related Topics 数组 👍 503 👎 0

func Test674(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "示例1",
			nums:     []int{1, 3, 5, 4, 7},
			expected: 3,
		},
		{
			name:     "示例2",
			nums:     []int{2, 2, 2, 2, 2},
			expected: 1,
		},
		{
			name:     "单个元素",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "完全递增",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "完全递减",
			nums:     []int{5, 4, 3, 2, 1},
			expected: 1,
		},
		{
			name:     "空数组",
			nums:     []int{},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, findLengthOfLCIS(tc.nums))
		})
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
func findLengthOfLCIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	left, right := 0, 0
	ans := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[right] {
			right++
		} else {
			ans = max(ans, right-left+1)
			left = i
			right = i
		}
	}

	ans = max(ans, right-left+1)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
