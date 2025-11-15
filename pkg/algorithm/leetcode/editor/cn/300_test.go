package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//
// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子
//序列。
//
// 示例 1：
//
//
//输入：nums = [10,9,2,5,3,7,101,18]
//输出：4
//解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
//
//
// 示例 2：
//
//
//输入：nums = [0,1,0,3,2,3]
//输出：4
//
//
// 示例 3：
//
//
//输入：nums = [7,7,7,7,7,7,7]
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 2500
// -10⁴ <= nums[i] <= 10⁴
//
//
//
//
// 进阶：
//
//
// 你能将算法的时间复杂度降低到 O(n log(n)) 吗?
//
//
// Related Topics 数组 二分查找 动态规划 👍 4108 👎 0

func Test300(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			want: 4,
		},
		{
			name: "example 2",
			nums: []int{0, 1, 0, 3, 2, 3},
			want: 4,
		},
		{
			name: "example 3",
			nums: []int{7, 7, 7, 7, 7, 7, 7},
			want: 1,
		},
		{
			name: "single element",
			nums: []int{1},
			want: 1,
		},
		{
			name: "strictly increasing",
			nums: []int{1, 2, 3, 4, 5},
			want: 5,
		},
		{
			name: "strictly decreasing",
			nums: []int{5, 4, 3, 2, 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLIS(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}

// O(n^2) 纯暴力的动态规划写法
func lengthOfLIS1(nums []int) int {
	dp := make(map[int]int)

	for _, v1 := range nums {
		dp[v1] = 1
		for k, v := range dp {
			if k < v1 {
				dp[v1] = max(dp[v1], v+1)
			}
		}
	}

	ans := 0
	for _, v := range dp {
		ans = max(ans, v)
	}

	return ans
}

// O(n^2) 更加巧妙的写法，dp是一个递增的子序列，二分法要基于这种写法来优化
func lengthOfLIS2(nums []int) int {
	dp := make([]int, 0)

	for _, v := range nums {
		idx := len(dp)
		for i1, v1 := range dp {
			if v <= v1 {
				idx = i1
				break
			}
		}

		if idx == len(dp) {
			dp = append(dp, v)
		} else {
			dp[idx] = v
		}
	}

	return len(dp)
}

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLIS(nums []int) int { // 基于lengthOfLIS2进行二分优化
	dp := make([]int, 0)

	for _, v := range nums {
		left, right := 0, len(dp)

		for left < right {
			m := (left + right) / 2
			if dp[m] < v {
				left = m + 1
			} else {
				right = m
			}
		}

		if left == len(dp) {
			dp = append(dp, v)
		} else {
			dp[left] = v
		}
	}

	return len(dp)
}

//leetcode submit region end(Prohibit modification and deletion)
