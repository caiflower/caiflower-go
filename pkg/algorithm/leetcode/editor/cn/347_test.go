package leetcode

import (
	"container/heap"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

//给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
//
//
//
// 示例 1：
//
//
// 输入：nums = [1,1,1,2,2,3], k = 2
//
//
// 输出：[1,2]
//
// 示例 2：
//
//
// 输入：nums = [1], k = 1
//
//
// 输出：[1]
//
// 示例 3：
//
//
// 输入：nums = [1,2,1,2,1,2,3,1,3,2], k = 2
//
//
// 输出：[1,2]
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// k 的取值范围是 [1, 数组中不相同的元素的个数]
// 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的
//
//
//
//
// 进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。
//
// Related Topics 数组 哈希表 分治 桶排序 计数 快速选择 排序 堆（优先队列） 👍 2083 👎 0

func Test347(t *testing.T) {
	// 测试用例1
	nums1 := []int{1, 1, 1, 2, 2, 3}
	k1 := 2
	result1 := topKFrequent(nums1, k1)
	assert.ElementsMatch(t, []int{1, 2}, result1, "测试用例1失败")
	result1 = topKFrequent1(nums1, k1)
	assert.ElementsMatch(t, []int{1, 2}, result1, "测试用例1失败")

	// 测试用例2
	nums2 := []int{1}
	k2 := 1
	result2 := topKFrequent(nums2, k2)
	assert.ElementsMatch(t, []int{1}, result2, "测试用例2失败")
	result2 = topKFrequent1(nums2, k2)
	assert.ElementsMatch(t, []int{1}, result2, "测试用例2失败")

	// 测试用例3
	nums3 := []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}
	k3 := 2
	result3 := topKFrequent(nums3, k3)
	assert.ElementsMatch(t, []int{1, 2}, result3, "测试用例3失败")
	result3 = topKFrequent1(nums3, k3)
	assert.ElementsMatch(t, []int{1, 2}, result3, "测试用例3失败")
}

// leetcode submit region begin(Prohibit modification and deletion)
func topKFrequent1(nums []int, k int) []int {
	var ans []int
	mapNum := make(map[int]int)
	for _, v := range nums {
		mapNum[v]++
	}

	var nums1 []int
	for key, _ := range mapNum {
		nums1 = append(nums1, key)
	}

	sort.Slice(nums1, func(i, j int) bool {
		return mapNum[nums1[i]] > mapNum[nums1[j]]
	})

	for i := 0; i < k; i++ {
		ans = append(ans, nums1[i])
	}
	return ans
}

type IntHeap [][2]int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i][1] < (*h)[j][1]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IntHeap) Pop() interface{} {
	val := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return val
}

func topKFrequent(nums []int, k int) []int {
	map_num := map[int]int{}
	//记录每个元素出现的次数
	for _, item := range nums {
		map_num[item]++
	}
	h := &IntHeap{}
	heap.Init(h)
	//所有元素入堆，堆的长度为k
	for key, value := range map_num {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]int, k)
	//按顺序返回堆中的元素
	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
