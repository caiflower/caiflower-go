package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//在两条独立的水平线上按给定的顺序写下 nums1 和 nums2 中的整数。
//
// 现在，可以绘制一些连接两个数字 nums1[i] 和 nums2[j] 的直线，这些直线需要同时满足：
//
//
// nums1[i] == nums2[j]
// 且绘制的直线不与任何其他连线（非水平线）相交。
//
//
// 请注意，连线即使在端点也不能相交：每个数字只能属于一条连线。
//
// 以这种方法绘制线条，并返回可以绘制的最大连线数。
//
//
//
// 示例 1：
//
//
//输入：nums1 = [1,4,2], nums2 = [1,2,4]
//输出：2
//解释：可以画出两条不交叉的线，如上图所示。
//但无法画出第三条不相交的直线，因为从 nums1[1]=4 到 nums2[2]=4 的直线将与从 nums1[2]=2 到 nums2[1]=2 的直线相
//交。
//
//
//
// 示例 2：
//
//
//
//输入：nums1 = [2,5,1,2,5], nums2 = [10,5,2,1,5,2]
//输出：3
//
//
//
// 示例 3：
//
//
//
//输入：nums1 = [1,3,7,1,7,5], nums2 = [1,9,2,5,1]
//输出：2
//
//
//
// 提示：
//
//
// 1 <= nums1.length, nums2.length <= 500
// 1 <= nums1[i], nums2[j] <= 2000
//
//
//
//
// Related Topics 数组 动态规划 👍 656 👎 0

func Test1035(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		nums1 []int
		nums2 []int
	}{
		{name: "示例1", want: 2, nums1: []int{1, 4, 2}, nums2: []int{1, 2, 4}},
		{name: "示例2", want: 3, nums1: []int{2, 5, 1, 2, 5}, nums2: []int{10, 5, 2, 1, 5, 2}},
		{name: "示例3", want: 2, nums1: []int{1, 3, 7, 1, 7, 5}, nums2: []int{1, 9, 2, 5, 1}},
		{name: "空数组", want: 0, nums1: []int{}, nums2: []int{}},
		{name: "单元素相同", want: 1, nums1: []int{1}, nums2: []int{1}},
		{name: "单元素不同", want: 0, nums1: []int{1}, nums2: []int{2}},
		{name: "完全相同", want: 3, nums1: []int{1, 2, 3}, nums2: []int{1, 2, 3}},
		{name: "完全不同", want: 0, nums1: []int{1, 2, 3}, nums2: []int{4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxUncrossedLines(tt.nums1, tt.nums2), tt.name)
		})
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[n1][n2]
}

//leetcode submit region end(Prohibit modification and deletion)
