package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给定一个二叉树，判断它是否是 平衡二叉树
//
//
//
// 示例 1：
//
//
//输入：root = [3,9,20,null,null,15,7]
//输出：true
//
//
// 示例 2：
//
//
//输入：root = [1,2,2,3,3,null,null,4,4]
//输出：false
//
//
// 示例 3：
//
//
//输入：root = []
//输出：true
//
//
//
//
// 提示：
//
//
// 树中的节点数在范围 [0, 5000] 内
// -10⁴ <= Node.val <= 10⁴
//
//
// Related Topics 树 深度优先搜索 二叉树 👍 1646 👎 0

func Test110(t *testing.T) {
	// 测试用例1：平衡二叉树 [3,9,20,null,null,15,7]
	tree1 := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	assert.True(t, isBalanced(tree1), "示例1应该返回true")

	// 测试用例2：不平衡二叉树 [1,2,2,3,3,null,null,4,4]
	tree2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 4},
			},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 2},
	}
	assert.False(t, isBalanced(tree2), "示例2应该返回false")

	// 测试用例3：空树
	assert.True(t, isBalanced(nil), "空树应该返回true")
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
func isBalanced(root *TreeNode) bool {
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
