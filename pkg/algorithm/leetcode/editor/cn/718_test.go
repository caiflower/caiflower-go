package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给两个整数数组 nums1 和 nums2 ，返回 两个数组中 公共的 、长度最长的子数组的长度 。
//
//
//
// 示例 1：
//
//
//输入：nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
//输出：3
//解释：长度最长的公共子数组是 [3,2,1] 。
//
//
// 示例 2：
//
//
//输入：nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]
//输出：5
//
//
//
//
// 提示：
//
//
// 1 <= nums1.length, nums2.length <= 1000
// 0 <= nums1[i], nums2[i] <= 100
//
//
// Related Topics 数组 二分查找 动态规划 滑动窗口 哈希函数 滚动哈希 👍 1211 👎 0

func Test718(t *testing.T) {
	testCases := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  int
	}{
		{
			name:  "示例1",
			nums1: []int{1, 2, 3, 2, 1},
			nums2: []int{3, 2, 1, 4, 7},
			want:  3,
		},
		{
			name:  "示例2",
			nums1: []int{0, 0, 0, 0, 0},
			nums2: []int{0, 0, 0, 0, 0},
			want:  5,
		},
		{
			name:  "空数组1",
			nums1: []int{},
			nums2: []int{1, 2, 3},
			want:  0,
		},
		{
			name:  "空数组2",
			nums1: []int{1, 2, 3},
			nums2: []int{},
			want:  0,
		},
		{
			name:  "单元素相同",
			nums1: []int{1},
			nums2: []int{1},
			want:  1,
		},
		{
			name:  "单元素不同",
			nums1: []int{1},
			nums2: []int{2},
			want:  0,
		},
		{
			name:  "无公共子数组",
			nums1: []int{1, 2, 3},
			nums2: []int{4, 5, 6},
			want:  0,
		},
		{
			name:  "完全相同",
			nums1: []int{1, 2, 3, 4, 5},
			nums2: []int{1, 2, 3, 4, 5},
			want:  5,
		},
		{
			name:  "公共子数组在开头",
			nums1: []int{1, 2, 3, 4, 5},
			nums2: []int{1, 2, 3, 6, 7},
			want:  3,
		},
		{
			name:  "公共子数组在结尾",
			nums1: []int{1, 2, 3, 4, 5},
			nums2: []int{6, 7, 3, 4, 5},
			want:  3,
		},
		{
			name:  "多个公共子数组取最长",
			nums1: []int{1, 2, 3, 4, 2, 3, 4, 5},
			nums2: []int{2, 3, 4, 5, 6},
			want:  4,
		},
		{
			name: "长度为1000的边界",
			nums1: func() []int {
				arr := make([]int, 1000)
				for i := range arr {
					arr[i] = i % 100
				}
				return arr
			}(),
			nums2: func() []int {
				arr := make([]int, 1000)
				for i := range arr {
					arr[i] = i % 100
				}
				return arr
			}(),
			want: 1000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := findLength(tc.nums1, tc.nums2)
			assert.Equal(t, tc.want, got)
		})
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
func findLength(nums1 []int, nums2 []int) int { // 时间复杂度O(n^2)
	n1, n2 := len(nums1), len(nums2)
	// nums1前i个数和nums2前j个数的最长公共子数组的长度
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}

	ans := 0
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				ans = max(ans, dp[i][j])
			}
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
