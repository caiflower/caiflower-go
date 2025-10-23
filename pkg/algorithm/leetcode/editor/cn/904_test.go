package leetcode

import (
	"fmt"
	"testing"
)

//你正在探访一家农场，农场从左到右种植了一排果树。这些树用一个整数数组 fruits 表示，其中 fruits[i] 是第 i 棵树上的水果 种类 。
//
// 你想要尽可能多地收集水果。然而，农场的主人设定了一些严格的规矩，你必须按照要求采摘水果：
//
//
// 你只有 两个 篮子，并且每个篮子只能装 单一类型 的水果。每个篮子能够装的水果总量没有限制。
// 你可以选择任意一棵树开始采摘，你必须从 每棵 树（包括开始采摘的树）上 恰好摘一个水果 。采摘的水果应当符合篮子中的水果类型。每采摘一次，你将会向右移动到
//下一棵树，并继续采摘。
// 一旦你走到某棵树前，但水果不符合篮子的水果类型，那么就必须停止采摘。
//
//
// 给你一个整数数组 fruits ，返回你可以收集的水果的 最大 数目。
//
//
//
// 示例 1：
//
//
//输入：fruits = [1,2,1]
//输出：3
//解释：可以采摘全部 3 棵树。
//
//
// 示例 2：
//
//
//输入：fruits = [0,1,2,2]
//输出：3
//解释：可以采摘 [1,2,2] 这三棵树。
//如果从第一棵树开始采摘，则只能采摘 [0,1] 这两棵树。
//
//
// 示例 3：
//
//
//输入：fruits = [1,2,3,2,2]
//输出：4
//解释：可以采摘 [2,3,2,2] 这四棵树。
//如果从第一棵树开始采摘，则只能采摘 [1,2] 这两棵树。
//
//
// 示例 4：
//
//
//输入：fruits = [3,3,3,1,2,1,1,2,3,3,4]
//输出：5
//解释：可以采摘 [1,2,1,1,2] 这五棵树。
//
//
//
//
// 提示：
//
//
// 1 <= fruits.length <= 10⁵
// 0 <= fruits[i] < fruits.length
//
//
// Related Topics 数组 哈希表 滑动窗口 👍 805 👎 0

func Test904(t *testing.T) {
	type testCase struct {
		fruits   []int
		expected int
	}
	cases := []testCase{
		// 题目示例
		{[]int{1, 2, 1}, 3},
		{[]int{0, 1, 2, 2}, 3},
		{[]int{1, 2, 3, 2, 2}, 4},
		{[]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}, 5},
		// 边界测试
		{[]int{1}, 1},                         // 只有一棵树
		{[]int{1, 1, 1, 1}, 4},                // 全部同类水果
		{[]int{1, 2, 3, 4, 5}, 2},             // 每棵树不同
		{[]int{1, 2, 1, 2, 1, 2, 1, 2}, 8},    // 两类交替
		{[]int{1, 2, 3, 1, 2, 3, 1, 2, 3}, 2}, // 三类交替
		{[]int{2, 2, 2, 2, 2}, 5},             // 全部同类
		// 大量同类后突变
		{[]int{1, 1, 1, 2, 2, 3, 3, 3, 3, 3}, 7},
		// 只出现两种类型但间隔分布
		{[]int{1, 2, 1, 2, 1, 2, 1, 2, 1, 2}, 10},
		// 三种类型，但只有两种连续
		{[]int{1, 2, 3, 2, 1, 1, 1, 2, 2, 2}, 7},
	}

	for i, tc := range cases {
		result := totalFruit(tc.fruits)
		fmt.Printf("Case %d: fruits=%v, Output=%d, Expected=%d\n", i+1, tc.fruits, result, tc.expected)
		if result == tc.expected {
			fmt.Println("PASS")
		} else {
			fmt.Println("FAIL")
		}

		fmt.Println()
	}
}

//if isFull(fruitTypes) || (!isFull(fruitTypes) && r >= len(fruits)) {
//ans = max(ans, r-l)
//if r >= len(fruits) {
//break
//}
//}

// leetcode submit region begin(Prohibit modification and deletion)
func totalFruit(fruits []int) int {
	ans := 0
	fruitTypes := [2][2]int{{-1, 0}, {-1, 0}}
	l, r := 0, 0

	for l <= r {
		if r >= len(fruits) {
			break
		} else if containsFruitType(fruitTypes, fruits[r]) || !isFull(fruitTypes) {
			fill(&fruitTypes, fruits[r])
			r++
			ans = max(ans, r-l)
		} else {
			remove(&fruitTypes, fruits[l])
			l++
		}
	}

	return ans
}

func containsFruitType(fruitTypes [2][2]int, target int) bool {
	for _, v := range fruitTypes {
		if v[0] == target {
			return true
		}
	}
	return false
}

func isFull(fruitTypes [2][2]int) bool {
	return fruitTypes[0][0] != -1 && fruitTypes[1][0] != -1
}

func fill(fruitTypes *[2][2]int, target int) {
	idx := 0

	for i, v := range fruitTypes {
		if v[0] == target {
			idx = i
			break
		} else if v[0] == -1 {
			idx = i
		}
	}

	fruitTypes[idx][0] = target
	fruitTypes[idx][1]++
}

func remove(fruitTypes *[2][2]int, target int) {
	idx := 0

	for i, v := range fruitTypes {
		if v[0] == target {
			idx = i
			break
		} else if v[0] == -1 {
			idx = i
		}
	}

	fruitTypes[idx][1]--
	if fruitTypes[idx][1] == 0 {
		fruitTypes[idx][0] = -1
	}
}

//leetcode submit region end(Prohibit modification and deletion)
