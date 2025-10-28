package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给你一个二叉树的根节点 root ， 检查它是否轴对称。
//
//
//
// 示例 1：
//
//
//输入：root = [1,2,2,3,4,4,3]
//输出：true
//
//
// 示例 2：
//
//
//输入：root = [1,2,2,null,3,null,3]
//输出：false
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [1, 1000] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶：你可以运用递归和迭代两种方法解决这个问题吗？
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 3021 👎 0

func Test101(t *testing.T) {
	// 测试用例1：对称的二叉树
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	assert.True(t, isSymmetric(root1), "示例1应该返回true")

	// 测试用例2：非对称的二叉树
	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	assert.False(t, isSymmetric(root2), "示例2应该返回false")

	// 测试用例3：空树
	var root3 *TreeNode
	assert.True(t, isSymmetric(root3), "空树应该被认为是对称的")

	// 测试用例4：只有根节点的树
	root4 := &TreeNode{Val: 1}
	assert.True(t, isSymmetric(root4), "只有根节点的树应该是对称的")
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func symemetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	} else if left.Val != right.Val {
		return false
	} else {
		if !symemetric(left.Left, right.Right) {
			return false
		} else {
			return symemetric(left.Right, right.Left)
		}
	}
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return symemetric(root.Left, root.Right)
}

//leetcode submit region end(Prohibit modification and deletion)
