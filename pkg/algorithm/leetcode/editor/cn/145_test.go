package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。
//
//
//
// 示例 1：
//
//
// 输入：root = [1,null,2,3]
//
//
// 输出：[3,2,1]
//
// 解释：
//
//
//
// 示例 2：
//
//
// 输入：root = [1,2,3,4,5,null,8,null,null,6,7,9]
//
//
// 输出：[4,6,7,5,2,9,8,3,1]
//
// 解释：
//
//
//
// 示例 3：
//
//
// 输入：root = []
//
//
// 输出：[]
//
// 示例 4：
//
//
// 输入：root = [1]
//
//
// 输出：[1]
//
//
//
// 提示：
//
//
// 树中节点的数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶：递归算法很简单，你可以通过迭代算法完成吗？
//
// Related Topics 栈 树 深度优先搜索 二叉树 👍 1257 👎 0

func Test145(t *testing.T) {
	// 用例1：root = [1, nil, 2, 3]
	root1 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 3},
		},
	}
	expected1 := []int{3, 2, 1}
	assert.Equal(t, expected1, postorderTraversal(root1), "case 1 failed")

	// 用例2：root = []
	var root2 *TreeNode = nil
	expected2 := []int{}
	assert.Equal(t, expected2, postorderTraversal(root2), "case 2 failed")

	// 用例3：root = [1]
	root3 := &TreeNode{Val: 1}
	expected3 := []int{1}
	assert.Equal(t, expected3, postorderTraversal(root3), "case 3 failed")

	// 用例4：root = [2,1,3]
	root4 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}
	expected4 := []int{1, 3, 2}
	assert.Equal(t, expected4, postorderTraversal(root4), "case 4 failed")

	// 用例5：root = [5,3,6,2,4]
	root5 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{Val: 6},
	}
	expected5 := []int{2, 4, 3, 6, 5}
	assert.Equal(t, expected5, postorderTraversal(root5), "case 5 failed")
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

type treeStack2 []*TreeNode

func (s *treeStack2) Push(x *TreeNode) {
	*s = append(*s, x)
}

func (s *treeStack2) Pop() *TreeNode {
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

func (s *treeStack2) Size() int {
	return len(*s)
}

// 前序遍历 --- 根左右 稍做修改为 根右左，结果反转即是后续遍历的左右根
func postorderTraversal(root *TreeNode) []int {
	t := make([]int, 0)
	s := treeStack2{}
	s.Push(root)

	for s.Size() > 0 {
		node := s.Pop()
		if node != nil {
			t = append(t, node.Val)
			s.Push(node.Left)
			s.Push(node.Right)
		}
	}

	ans := make([]int, 0, len(t)*2)
	for i := len(t) - 1; i >= 0; i-- {
		ans = append(ans, t[i])
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
