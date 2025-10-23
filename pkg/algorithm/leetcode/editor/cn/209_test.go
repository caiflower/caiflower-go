package leetcode

import (
	"fmt"
	"math"
	"testing"
)

//Given an array of positive integers nums and a positive integer target,
//return the minimal length of a subarray whose sum is greater than or equal to target.
//If there is no such subarray, return 0 instead.
//
//
// Example 1:
//
//
//Input: target = 7, nums = [2,3,1,2,4,3]
//Output: 2
//Explanation: The subarray [4,3] has the minimal length under the problem
//constraint.
//
//
// Example 2:
//
//
//Input: target = 4, nums = [1,4,4]
//Output: 1
//
//
// Example 3:
//
//
//Input: target = 11, nums = [1,1,1,1,1,1,1,1]
//Output: 0
//
//
//
// Constraints:
//
//
// 1 <= target <= 10⁹
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁴
//
//
//
//Follow up: If you have figured out the
//O(n) solution, try coding another solution of which the time complexity is
//O(n log(n)).
//
// Related Topics 数组 二分查找 前缀和 滑动窗口 👍 2536 👎 0

func Test209(t *testing.T) {
	// Example 1
	target1 := 7
	nums1 := []int{2, 3, 1, 2, 4, 3}
	expected1 := 2
	result1 := minSubArrayLen(target1, nums1)
	fmt.Printf("Example 1: Input: target=%d, nums=%v\n", target1, nums1)
	fmt.Printf("Output: %d, Expected: %d\n", result1, expected1)
	fmt.Println()

	// Example 2
	target2 := 4
	nums2 := []int{1, 4, 4}
	expected2 := 1
	result2 := minSubArrayLen(target2, nums2)
	fmt.Printf("Example 2: Input: target=%d, nums=%v\n", target2, nums2)
	fmt.Printf("Output: %d, Expected: %d\n", result2, expected2)
	fmt.Println()

	// Example 3
	target3 := 11
	nums3 := []int{1, 1, 1, 1, 1, 1, 1, 1}
	expected3 := 0
	result3 := minSubArrayLen(target3, nums3)
	fmt.Printf("Example 3: Input: target=%d, nums=%v\n", target3, nums3)
	fmt.Printf("Output: %d, Expected: %d\n", result3, expected3)
	fmt.Println()
}

// leetcode submit region begin(Prohibit modification and deletion)
func minSubArrayLen(target int, nums []int) int {
	ans := math.MaxInt64
	l, r, sum := 0, 0, 0
	for l <= r {
		if sum >= target {
			ans = min(ans, r-l)
			sum -= nums[l]
			l++
		} else if r < len(nums) {
			sum += nums[r]
			r++
		} else {
			break
		}
	}

	if ans == math.MaxInt64 {
		return 0
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
