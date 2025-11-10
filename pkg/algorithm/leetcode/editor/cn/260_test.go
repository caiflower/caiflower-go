package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给你一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。
//
// 你必须设计并实现线性时间复杂度的算法且仅使用常量额外空间来解决此问题。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,1,3,2,5]
//输出：[3,5]
//解释：[5, 3] 也是有效的答案。
//
//
// 示例 2：
//
//
//输入：nums = [-1,0]
//输出：[-1,0]
//
//
// 示例 3：
//
//
//输入：nums = [0,1]
//输出：[1,0]
//
//
//
//
// 提示：
//
//
// 2 <= nums.length <= 3 * 10⁴
// -2³¹ <= nums[i] <= 2³¹ - 1
// 除两个只出现一次的整数外，nums 中的其他数字都出现两次
//
//
// Related Topics 位运算 数组 👍 964 👎 0

func Test260(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "示例1", nums: []int{1, 2, 1, 3, 2, 5}, want: []int{3, 5}},
		{name: "示例2", nums: []int{-1, 0}, want: []int{-1, 0}},
		{name: "示例3", nums: []int{0, 1}, want: []int{1, 0}},
		{name: "多个重复元素", nums: []int{1, 1, 2, 2, 3, 4, 5, 5}, want: []int{3, 4}},
		{name: "负数测试", nums: []int{-10, 10, -10, 20}, want: []int{10, 20}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := singleNumber(tt.nums)
			// 由于结果顺序不重要，需要先排序再比较
			assert.ElementsMatch(t, tt.want, result, tt.name)
		})
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
func singleNumber(nums []int) []int {
	xy := 0

	// xy = x ^ y
	for _, v := range nums {
		xy ^= v
	}

	// 找出x^y 异或中的特殊的1
	num := xy
	k := 0
	for num != 0 {
		if num&1 != 1 {
			k++
			num >>= 1
		} else {
			break
		}
	}

	num = xy
	// 算出其中一个特殊的数
	for _, v := range nums {
		if (v>>k)&1 == 1 {
			num ^= v
		}
	}

	return []int{num, num ^ xy}
}

//leetcode submit region end(Prohibit modification and deletion)
