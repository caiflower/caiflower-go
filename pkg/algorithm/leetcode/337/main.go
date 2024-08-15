package main

//小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为
// root 。
//
// 除了
// root 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果 两个直接相连的
//房子在同一天晚上被打劫 ，房屋将自动报警。
//
// 给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 。
//
//
//
// 示例 1:
//
//
//
//
//输入: root = [3,2,3,null,3,null,1]
//输出: 7
//解释: 小偷一晚能够盗取的最高金额 3 + 3 + 1 = 7
//
// 示例 2:
//
//
//
//
//输入: root = [3,4,5,1,3,null,1]
//输出: 9
//解释: 小偷一晚能够盗取的最高金额 4 + 5 = 9
//
//
//
//
// 提示：
//
//
//
//
//
// 树的节点数在 [1, 10⁴] 范围内
// 0 <= Node.val <= 10⁴
//
//
// Related Topics 树 深度优先搜索 动态规划 二叉树 👍 1992 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) (res int) {
	m1, m2 := make(map[*TreeNode]int), make(map[*TreeNode]int)
	if root == nil {
		return
	}

	m1[root] += root.Val
	m2[root] = 0

	deep(m1, m2, root.Left)
	deep(m1, m2, root.Right)
	for _, v := range m2 {
		res = maxInt(res, v)
	}

	for _, v := range m1 {
		res = maxInt(res, v)
	}

	return
}

func deep(m1, m2 map[*TreeNode]int, node *TreeNode) {

}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
