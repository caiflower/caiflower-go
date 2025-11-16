package leetcode

import "testing"

//给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target ，写一个函数搜索 nums 中的 target，如果 target 存在返
//回下标，否则返回 -1。
//
// 你必须编写一个具有 O(log n) 时间复杂度的算法。
//
// 示例 1:
//
//
//输入: nums = [-1,0,3,5,9,12], target = 9
//输出: 4
//解释: 9 出现在 nums 中并且下标为 4
//
//
// 示例 2:
//
//
//输入: nums = [-1,0,3,5,9,12], target = 2
//输出: -1
//解释: 2 不存在 nums 中因此返回 -1
//
//
//
//
// 提示：
//
//
// 你可以假设 nums 中的所有元素是不重复的。
// n 将在 [1, 10000]之间。
// nums 的每个元素都将在 [-9999, 9999]之间。
//
//
// Related Topics 数组 二分查找 👍 1815 👎 0

func Test704(t *testing.T) {
}

// leetcode submit region begin(Prohibit modification and deletion)
func search1(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) >> 1
		if nums[mid] < target {
			left = mid + 1 // left = mid + 1，那么mid 就是(left + right) >> 1 不用加1
		} else {
			right = mid
		}
	}

	if nums[left] != target {
		return -1
	}

	return left
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right + 1) >> 1
		if nums[mid] <= target {
			left = mid // left = mid，那么mid就是(left + right + 1) >> 1, 需要加1
		} else {
			right = mid - 1
		}
	}

	if nums[left] != target {
		return -1
	}

	return left
}

//leetcode submit region end(Prohibit modification and deletion)
