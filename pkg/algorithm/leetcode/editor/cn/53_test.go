package leetcode

import (
	"testing"
)

//给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
// 子数组 是数组中的一个连续部分。
//
//
//
// 示例 1：
//
//
//输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
//输出：6
//解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
//
//
// 示例 2：
//
//
//输入：nums = [1]
//输出：1
//
//
// 示例 3：
//
//
//输入：nums = [5,4,-1,7,8]
//输出：23
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
//
//
//
//
// 进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
//
// Related Topics 数组 分治 动态规划 👍 7192 👎 0

func Test53(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "示例1：包含负数的数组",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			name: "示例2：单个元素",
			nums: []int{1},
			want: 1,
		},
		{
			name: "示例3：多个正数",
			nums: []int{5, 4, -1, 7, 8},
			want: 23,
		},
		{
			name: "全部负数",
			nums: []int{-3, -2, -5, -1},
			want: -1,
		},
		{
			name: "全部正数",
			nums: []int{1, 2, 3, 4, 5},
			want: 15,
		},
		{
			name: "交替正负数",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

//leetcode submit region begin(Prohibit modification and deletion)

//	腾讯面试题
//
// 贪心
// O(n)
func maxSubArray(nums []int) int {
	ans := -(1<<63 - 1)

	cur := 0
	for _, num := range nums {
		// 前面的和为负数了，就没必要加上去了，贪心
		if cur < 0 {
			cur = 0
		}
		cur += num
		ans = max(ans, cur)
	}
	return ans
}

// 动态规划
// O(n)
func maxSubArray1(nums []int) int {
	n := len(nums)

	// 以i结尾的最大值
	dp := make([]int, n+1)
	ans := -(1<<63 - 1)

	for i, v := range nums {
		dp[i+1] = max(dp[i]+v, v)
		ans = max(ans, dp[i+1])
	}

	return ans
}

// 前缀和 O(n^2)
// 力扣超时
func maxSubArray2(nums []int) int {
	ans := -(1<<63 - 1)

	n := len(nums)
	sums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sums[i] = nums[i-1] + sums[i-1]
	}

	// 枚举子数组
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			// 前j个的和 减去前i-1个的和
			ans = max(ans, sums[j+1]-sums[i])
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
