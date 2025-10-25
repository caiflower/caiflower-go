package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
//
//
// 示例 1：
//
//
//
//
//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6
//解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//
//
// 示例 2：
//
//
//输入：height = [4,2,0,3,2,5]
//输出：9
//
//
//
//
// 提示：
//
//
// n == height.length
// 1 <= n <= 2 * 10⁴
// 0 <= height[i] <= 10⁵
//
//
// Related Topics 栈 数组 双指针 动态规划 单调栈 👍 5938 👎 0

func Test42(t *testing.T) {
	// 测试用例1：示例1中的数据
	height1 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	assert.Equal(t, 6, trap(height1), "接雨水测试用例1失败")

	// 测试用例2：示例2中的数据
	height2 := []int{4, 2, 0, 3, 2, 5}
	assert.Equal(t, 9, trap(height2), "接雨水测试用例2失败")

	// 测试用例3：空数组
	height3 := []int{}
	assert.Equal(t, 0, trap(height3), "接雨水空数组测试失败")

	// 测试用例4：单个元素数组
	height4 := []int{5}
	assert.Equal(t, 0, trap(height4), "接雨水单元素数组测试失败")

	// 测试用例5：所有元素相等
	height5 := []int{2, 2, 2, 2, 2}
	assert.Equal(t, 0, trap(height5), "接雨水所有元素相等测试失败")

	// 测试用例6：递增序列
	height6 := []int{1, 2, 3, 4, 5}
	assert.Equal(t, 0, trap(height6), "接雨水递增序列测试失败")

	// 测试用例7：递减序列
	height7 := []int{5, 4, 3, 2, 1}
	assert.Equal(t, 0, trap(height7), "接雨水递减序列测试失败")

	// 测试用例8：两侧高中间低
	height8 := []int{5, 1, 2, 1, 5}
	assert.Equal(t, 11, trap(height8), "接雨水两侧高中间低测试失败")

	// 测试用例9：两个元素 - 边界情况
	height9 := []int{3, 1}
	assert.Equal(t, 0, trap(height9), "接雨水两个元素测试失败")

	// 测试用例10：两个元素相等 - 边界情况
	height10 := []int{2, 2}
	assert.Equal(t, 0, trap(height10), "接雨水两个相等元素测试失败")

	// 测试用例11：全零数组 - 边界情况
	height11 := []int{0, 0, 0, 0}
	assert.Equal(t, 0, trap(height11), "接雨水全零数组测试失败")

	// 测试用例12：单峰形状 - 特殊模式
	height12 := []int{1, 2, 3, 2, 1}
	assert.Equal(t, 0, trap(height12), "接雨水单峰形状测试失败")

	// 测试用例13：V型 - 特殊模式
	height13 := []int{3, 0, 2}
	assert.Equal(t, 2, trap(height13), "接雨水V型测试失败")

	// 测试用例14：深V型 - 特殊模式
	height14 := []int{5, 0, 0, 0, 5}
	assert.Equal(t, 15, trap(height14), "接雨水深V型测试失败")

	// 测试用例15：W型 - 特殊模式
	height15 := []int{3, 0, 2, 0, 4}
	assert.Equal(t, 7, trap(height15), "接雨水W型测试失败")

	// 测试用例16：复杂W型 - 特殊模式
	height16 := []int{4, 1, 3, 1, 5}
	assert.Equal(t, 7, trap(height16), "接雨水复杂W型测试失败")

	// 测试用例17：阶梯型上升 - 特殊模式
	height17 := []int{1, 1, 2, 2, 3, 3}
	assert.Equal(t, 0, trap(height17), "接雨水阶梯型上升测试失败")

	// 测试用例18：阶梯型下降 - 特殊模式
	height18 := []int{3, 3, 2, 2, 1, 1}
	assert.Equal(t, 0, trap(height18), "接雨水阶梯型下降测试失败")

	// 测试用例19：波浪型 - 特殊模式
	height19 := []int{2, 1, 2, 1, 2}
	assert.Equal(t, 2, trap(height19), "接雨水波浪型测试失败")

	// 测试用例20：复杂波浪型 - 特殊模式
	height20 := []int{3, 1, 3, 1, 3, 1, 3}
	assert.Equal(t, 6, trap(height20), "接雨水复杂波浪型测试失败")

	// 测试用例21：左侧高右侧低 - 边界情况
	height21 := []int{5, 1, 2, 3}
	assert.Equal(t, 3, trap(height21), "接雨水左侧高右侧低测试失败")

	// 测试用例22：右侧高左侧低 - 边界情况
	height22 := []int{1, 2, 3, 5}
	assert.Equal(t, 0, trap(height22), "接雨水右侧高左侧低测试失败")

	// 测试用例23：中间高两侧低 - 特殊模式
	height23 := []int{1, 5, 1}
	assert.Equal(t, 0, trap(height23), "接雨水中间高两侧低测试失败")

	// 测试用例24：多个小坑 - 复杂情况
	height24 := []int{2, 1, 2, 1, 2, 1, 2}
	assert.Equal(t, 3, trap(height24), "接雨水多个小坑测试失败")

	// 测试用例25：大坑小坑组合 - 复杂情况
	height25 := []int{5, 1, 3, 1, 5, 2, 4}
	assert.Equal(t, 12, trap(height25), "接雨水大坑小坑组合测试失败")

	// 测试用例26：不规则高度分布 - 复杂情况
	height26 := []int{6, 4, 2, 0, 3, 2, 0, 3, 1, 4, 5, 3, 2, 7, 5, 3, 0, 2}
	assert.Equal(t, 45, trap(height26), "接雨水不规则高度分布测试失败")

	// 测试用例27：只有开头和结尾高 - 边界情况
	height27 := []int{4, 0, 0, 0, 0, 4}
	assert.Equal(t, 16, trap(height27), "接雨水只有开头和结尾高测试失败")

	// 测试用例28：三个元素简单情况 - 边界情况
	height28 := []int{2, 1, 3}
	assert.Equal(t, 1, trap(height28), "接雨水三个元素简单情况测试失败")

	// 测试用例29：三个元素无积水 - 边界情况
	height29 := []int{1, 2, 3}
	assert.Equal(t, 0, trap(height29), "接雨水三个元素无积水测试失败")

	// 测试用例30：多峰情况 - 复杂情况
	height30 := []int{3, 1, 3, 1, 2, 1, 4}
	assert.Equal(t, 7, trap(height30), "接雨水多峰情况测试失败")

	// 测试用例31：大数值测试 - 极端情况
	height31 := []int{100000, 0, 100000}
	assert.Equal(t, 100000, trap(height31), "接雨水大数值测试失败")

	// 测试用例32：大数值复杂情况 - 极端情况
	height32 := []int{50000, 25000, 75000, 10000, 90000}
	assert.Equal(t, 90000, trap(height32), "接雨水大数值复杂情况测试失败")

	// 测试用例33：最小值边界 - 极端情况
	height33 := []int{0, 1, 0}
	assert.Equal(t, 0, trap(height33), "接雨水最小值边界测试失败")

	// 测试用例34：长数组递增 - 复杂情况
	height34 := make([]int, 100)
	for i := 0; i < 100; i++ {
		height34[i] = i + 1
	}
	assert.Equal(t, 0, trap(height34), "接雨水长数组递增测试失败")

	// 测试用例35：长数组递减 - 复杂情况
	height35 := make([]int, 100)
	for i := 0; i < 100; i++ {
		height35[i] = 100 - i
	}
	assert.Equal(t, 0, trap(height35), "接雨水长数组递减测试失败")

	// 测试用例36：长数组V型 - 复杂情况
	height36 := []int{100, 1, 1, 1, 1, 1, 1, 1, 1, 1, 100}
	assert.Equal(t, 891, trap(height36), "接雨水长数组V型测试失败")

	// 测试用例37：锯齿状 - 特殊模式
	height37 := []int{1, 0, 1, 0, 1, 0, 1}
	assert.Equal(t, 3, trap(height37), "接雨水锯齿状测试失败")

	// 测试用例38：平台型 - 特殊模式
	height38 := []int{3, 3, 1, 1, 3, 3}
	assert.Equal(t, 4, trap(height38), "接雨水平台型测试失败")

	// 测试用例39：复杂平台型 - 特殊模式
	height39 := []int{5, 5, 1, 3, 3, 1, 5, 5}
	assert.Equal(t, 12, trap(height39), "接雨水复杂平台型测试失败")

	// 测试用例40：山谷型 - 特殊模式
	height40 := []int{4, 3, 2, 1, 2, 3, 4}
	assert.Equal(t, 9, trap(height40), "接雨水山谷型测试失败")

	// 测试用例41：不对称山谷 - 特殊模式
	height41 := []int{6, 2, 1, 0, 1, 3}
	assert.Equal(t, 8, trap(height41), "接雨水不对称山谷测试失败")

	// 测试用例42：多层阶梯 - 复杂情况
	height42 := []int{5, 1, 2, 1, 3, 1, 4, 1, 5}
	assert.Equal(t, 22, trap(height42), "接雨水多层阶梯测试失败")

	// 测试用例43：倒U型 - 特殊模式
	height43 := []int{1, 3, 5, 3, 1}
	assert.Equal(t, 0, trap(height43), "接雨水倒U型测试失败")

	// 测试用例44：复合型 - 复杂情况
	height44 := []int{3, 0, 2, 0, 4, 0, 3, 2, 1, 0, 1, 2, 3}
	assert.Equal(t, 19, trap(height44), "接雨水复合型测试失败")

	// 测试用例45：边缘凹陷 - 边界情况
	height45 := []int{0, 2, 0}
	assert.Equal(t, 0, trap(height45), "接雨水边缘凹陷测试失败")

	// 测试用例46：深坑浅坑 - 复杂情况
	height46 := []int{4, 0, 0, 2, 0, 0, 4}
	assert.Equal(t, 18, trap(height46), "接雨水深坑浅坑测试失败")

	// 测试用例47：起伏不定 - 复杂情况
	height47 := []int{2, 1, 0, 1, 3, 0, 1, 2, 0, 3, 1, 0, 2}
	assert.Equal(t, 16, trap(height47), "接雨水起伏不定测试失败")

	// 测试用例48：高低交替 - 特殊模式
	height48 := []int{5, 0, 5, 0, 5, 0, 5}
	assert.Equal(t, 15, trap(height48), "接雨水高低交替测试失败")

	// 测试用例49：渐进式 - 复杂情况
	height49 := []int{1, 0, 2, 0, 3, 0, 4, 0, 5}
	assert.Equal(t, 10, trap(height49), "接雨水渐进式测试失败")

	// 测试用例50：综合测试 - 极端情况
	height50 := []int{10, 5, 0, 7, 3, 0, 8, 2, 1, 0, 6, 4, 0, 9, 1, 0, 5, 3, 2, 10}
	assert.Equal(t, 124, trap(height50), "接雨水综合测试失败")
}

// leetcode submit region begin(Prohibit modification and deletion)
type intStack [][2]int

func (s *intStack) Push(x [2]int) {
	*s = append(*s, x)
}

func (s *intStack) Pop() [2]int {
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

func (s *intStack) Peek() [2]int {
	return (*s)[len(*s)-1]
}

func (s *intStack) Size() int {
	return len(*s)
}

// 最容易理解的方法
// 暴力法的优化，找出每个柱子两边的高柱，用小的高柱减去这个柱子的高度，求和就是结果
// 时间复杂度：O(n)，空间复杂度：O(n)
func trap1(height []int) int {
	if 0 == len(height) {
		return 0
	}

	ans := 0
	var l, r [20000]int
	lv, rv := 0, 0
	for i := 0; i < len(height); i++ {
		l[i] = lv
		lv = max(height[i], lv)
	}
	for i := len(height) - 1; i >= 0; i-- {
		r[i] = rv
		rv = max(height[i], rv)
	}

	for i := 1; i < len(height); i++ {
		if r[i] > height[i] && l[i] > height[i] {
			ans += min(r[i], l[i]) - height[i]
		}
	}

	return ans
}

// 双指针法，从左右两边开始遍历，每次移动高度小的指针（因为接水的时候关心的是矮的柱子），并更新最大高度，如果当前高度小于最大高度，则可以接雨水，累加到结果中。
// 时间复杂度：O(n)，空间复杂度：O(1)
func trap2(height []int) int {
	if 0 == len(height) {
		return 0
	}

	ans := 0

	l, r := 1, len(height)-2
	lv, rv := height[l-1], height[r+1]
	for l <= r {
		if lv < rv {
			if lv > height[l] {
				ans += lv - height[l]
			}
			lv = max(lv, height[l])
			l++
		} else {
			if rv > height[r] {
				ans += rv - height[r]
			}
			rv = max(rv, height[r])
			r--
		}
	}

	return ans
}

// 单调栈写法
// 单调递减栈求雨水，相对复杂，不太好想出来，建议用示例2来想着写
// 时间复杂度O(n), 空间复杂度O(n)
func trap(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}

	s := intStack{}
	ans := 0

	for i := 0; i < len(height); i++ {
		tmp := 0
		for s.Size() > 0 && height[i] >= s.Peek()[1] {
			t := s.Pop()
			if s.Size() > 0 {
				tmp += (i - s.Peek()[0] - 1) * (min(s.Peek()[1], height[i]) - t[1])
			}
		}

		ans += tmp
		s.Push([2]int{i, height[i]})
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
