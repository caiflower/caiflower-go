package leetcode

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

//给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
//
//
// 示例 1：
//
//
//输入：nums = [3,2,1,5,6,4], k = 2
//输出：5
//
//
// 示例 2：
//
//
//输入：nums = [3,2,3,1,2,4,5,5,6], k = 4
//输出：4
//
//
//
// 提示：
//
//
// 1 <= k <= nums.length <= 10⁴
// -10⁴ <= nums[i] <= 10⁴
//
//
//
//
//
// 注意：本题与主站 215 题相同： https://leetcode-cn.com/problems/kth-largest-element-in-an-
//array/
//
// Related Topics 数组 分治 快速选择 排序 堆（优先队列） 👍 109 👎 0

func TestLCR_076(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		k      int
		expect int
	}{
		{
			name:   "示例1",
			nums:   []int{3, 2, 1, 5, 6, 4},
			k:      2,
			expect: 5,
		},
		{
			name:   "示例2",
			nums:   []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:      4,
			expect: 4,
		},
		{
			name:   "单元素数组",
			nums:   []int{1},
			k:      1,
			expect: 1,
		},
		{
			name:   "所有元素相同",
			nums:   []int{3, 3, 3, 3},
			k:      2,
			expect: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findKthLargest(tt.nums, tt.k)
			assert.Equal(t, tt.expect, got)
		})
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	h := &KHeap{}

	heap.Init(h)

	for i := 0; i < n; i++ {
		heap.Push(h, nums[i])
	}

	t := 0
	for i := 0; i < k; i++ {
		t = heap.Pop(h).(int)
	}

	return t
}

type KHeap []int

func (h *KHeap) Len() int {
	return len(*h)
}

func (h *KHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *KHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *KHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *KHeap) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

//leetcode submit region end(Prohibit modification and deletion)
