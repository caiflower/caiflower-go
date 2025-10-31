package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给定二叉搜索树（BST）的根节点
// root 和一个整数值
// val。
//
// 你需要在 BST 中找到节点值等于 val 的节点。 返回以该节点为根的子树。 如果节点不存在，则返回
// null 。
//
//
//
// 示例 1:
//
//
//
//
//
//输入：root = [4,2,7,1,3], val = 2
//输出：[2,1,3]
//
//
// 示例 2:
//
//
//输入：root = [4,2,7,1,3], val = 5
//输出：[]
//
//
//
//
// 提示：
//
//
// 树中节点数在 [1, 5000] 范围内
// 1 <= Node.val <= 10⁷
// root 是二叉搜索树
// 1 <= val <= 10⁷
//
//
// Related Topics 树 二叉搜索树 二叉树 👍 515 👎 0

func Test700(t *testing.T) {
	// 构建测试用例
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  nil,
			Right: nil,
		},
	}

	// 测试用例1：存在的节点
	result1 := searchBST(root, 2)
	expected1 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	assert.Equal(t, expected1.Val, result1.Val)
	assert.Equal(t, expected1.Left.Val, result1.Left.Val)
	assert.Equal(t, expected1.Right.Val, result1.Right.Val)

	// 测试用例2：不存在的节点
	result2 := searchBST(root, 5)
	assert.Nil(t, result2)
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
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	} else if root.Val > val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
