package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
//
//
//
// 示例 1：
//
//
//输入：root = [1,null,2,3]
//输出：[1,3,2]
//
//
// 示例 2：
//
//
//输入：root = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：root = [1]
//输出：[1]
//
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
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
//
// Related Topics 栈 树 深度优先搜索 二叉树 👍 2320 👎 0

func Test94(t *testing.T) {
	// 用例1：root = [1, nil, 2, 3]
	root1 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 3},
		},
	}
	expected1 := []int{1, 3, 2}
	assert.Equal(t, expected1, inorderTraversal(root1), "case 1 failed")

	// 用例2：root = []
	var root2 *TreeNode = nil
	expected2 := []int{}
	assert.Equal(t, expected2, inorderTraversal(root2), "case 2 failed")

	// 用例3：root = [1]
	root3 := &TreeNode{Val: 1}
	expected3 := []int{1}
	assert.Equal(t, expected3, inorderTraversal(root3), "case 3 failed")

	// 用例4：root = [2,1,3]
	root4 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}
	expected4 := []int{1, 2, 3}
	assert.Equal(t, expected4, inorderTraversal(root4), "case 4 failed")

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
	expected5 := []int{2, 3, 4, 5, 6}
	assert.Equal(t, expected5, inorderTraversal(root5), "case 5 failed")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
type treeStack []*TreeNode

func (s *treeStack) Push(x *TreeNode) {
	*s = append(*s, x)
}

func (s *treeStack) Pop() *TreeNode {
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

func (s *treeStack) Size() int {
	return len(*s)
}

// 递归法
func inorderTraversal1(root *TreeNode) []int {
	ans := make([]int, 0)
	midleTravel(root, &ans)
	return ans
}

func midleTravel(root *TreeNode, ans *[]int) {
	if root != nil {
		midleTravel(root.Left, ans)
		*ans = append(*ans, root.Val)
		midleTravel(root.Right, ans)
	}
}

// 迭代法
func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	p := root
	s := treeStack{}

	for p != nil || s.Size() > 0 {
		if p != nil {
			s.Push(p)
			p = p.Left
		} else {
			node := s.Pop()
			ans = append(ans, node.Val)
			if node.Right != nil {
				p = node.Right
			}
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
