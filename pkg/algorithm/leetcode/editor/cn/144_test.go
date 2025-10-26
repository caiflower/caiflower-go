package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
//
//
//
// 示例 1：
//
//
// 输入：root = [1,null,2,3]
//
//
// 输出：[1,2,3]
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
// 输出：[1,2,4,5,6,7,3,8,9]
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
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶：递归算法很简单，你可以通过迭代算法完成吗？
//
// Related Topics 栈 树 深度优先搜索 二叉树 👍 1360 👎 0

func Test144(t *testing.T) {
	// 用例1：root = [1, nil, 2, 3]
	root1 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 3},
		},
	}
	expected1 := []int{1, 2, 3}
	assert.Equal(t, expected1, preorderTraversal(root1), "case 1 failed")

	// 用例2：root = []
	var root2 *TreeNode = nil
	expected2 := []int{}
	assert.Equal(t, expected2, preorderTraversal(root2), "case 2 failed")

	// 用例3：root = [1]
	root3 := &TreeNode{Val: 1}
	expected3 := []int{1}
	assert.Equal(t, expected3, preorderTraversal(root3), "case 3 failed")

	// 用例4：root = [2,1,3]
	root4 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}
	expected4 := []int{2, 1, 3}
	assert.Equal(t, expected4, preorderTraversal(root4), "case 4 failed")

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
	expected5 := []int{5, 3, 2, 4, 6}
	assert.Equal(t, expected5, preorderTraversal(root5), "case 5 failed")
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
type treeStack1 []*TreeNode

func (s *treeStack1) Push(x *TreeNode) {
	*s = append(*s, x)
}

func (s *treeStack1) Pop() *TreeNode {
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

func (s *treeStack1) Size() int {
	return len(*s)
}

func preorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	s := treeStack1{}
	s.Push(root)

	for s.Size() > 0 {
		node := s.Pop()
		if node == nil {
			continue
		}
		ans = append(ans, node.Val)
		s.Push(node.Right)
		s.Push(node.Left)
	}

	return ans
}

func preorderTraversal1(root *TreeNode) []int {
	ans := make([]int, 0)
	pre(root, &ans)
	return ans
}

func pre(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	*ans = append(*ans, root.Val)
	pre(root.Left, ans)
	pre(root.Right, ans)
}

//leetcode submit region end(Prohibit modification and deletion)
