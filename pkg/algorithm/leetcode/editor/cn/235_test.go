package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
//
// 最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（
//一个节点也可以是它自己的祖先）。”
//
// 例如，给定如下二叉搜索树: root = [6,2,8,0,4,7,9,null,null,3,5]
//
//
//
//
//
// 示例 1:
//
// 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
//输出: 6
//解释: 节点 2 和节点 8 的最近公共祖先是 6。
//
//
// 示例 2:
//
// 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
//输出: 2
//解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。
//
//
//
// 说明:
//
//
// 所有节点的值都是唯一的。
// p、q 为不同节点且均存在于给定的二叉搜索树中。
//
//
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 1341 👎 0

func Test235(t *testing.T) {
	// 测试用例1
	// 构建树 [6,2,8,0,4,7,9,null,null,3,5]
	root1 := &TreeNode{Val: 6}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 8}
	root1.Left.Left = &TreeNode{Val: 0}
	root1.Left.Right = &TreeNode{Val: 4}
	root1.Right.Left = &TreeNode{Val: 7}
	root1.Right.Right = &TreeNode{Val: 9}
	root1.Left.Right.Left = &TreeNode{Val: 3}
	root1.Left.Right.Right = &TreeNode{Val: 5}

	// 示例1: p=2, q=8, 期望输出=6
	p1 := root1.Left  // 值为2的节点
	q1 := root1.Right // 值为8的节点
	result1 := lowestCommonAncestor(root1, p1, q1)
	assert.Equal(t, 6, result1.Val, "示例1测试失败")

	// 示例2: p=2, q=4, 期望输出=2
	p2 := root1.Left       // 值为2的节点
	q2 := root1.Left.Right // 值为4的节点
	result2 := lowestCommonAncestor(root1, p2, q2)
	assert.Equal(t, 2, result2.Val, "示例2测试失败")

	// 额外测试: p=3, q=5, 期望输出=4
	p3 := root1.Left.Right.Left  // 值为3的节点
	q3 := root1.Left.Right.Right // 值为5的节点
	result3 := lowestCommonAncestor(root1, p3, q3)
	assert.Equal(t, 4, result3.Val, "额外测试失败")
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q || root == nil {
		return root
	}

	v1 := lowestCommonAncestor(root.Left, p, q)
	v2 := lowestCommonAncestor(root.Right, p, q)
	if v1 != nil && v2 != nil {
		return root
	} else if v1 != nil {
		return v1
	} else if v2 != nil {
		return v2
	} else {
		return nil
	}
}

//leetcode submit region end(Prohibit modification and deletion)
