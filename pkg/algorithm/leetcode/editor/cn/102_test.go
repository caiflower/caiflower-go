package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
//
//
//
// 示例 1：
//
//
//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[9,20],[15,7]]
//
//
// 示例 2：
//
//
//输入：root = [1]
//输出：[[1]]
//
//
// 示例 3：
//
//
//输入：root = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 2000] 内
// -1000 <= Node.val <= 1000
//
//
// Related Topics 树 广度优先搜索 二叉树 👍 2198 👎 0

func Test102(t *testing.T) {
	// 用例1：root = [3,9,20,null,null,15,7]
	root1 := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	expected1 := [][]int{{3}, {9, 20}, {15, 7}}
	assert.Equal(t, expected1, levelOrder(root1), "case 1 failed")

	// 用例2：root = [1]
	root2 := &TreeNode{Val: 1}
	expected2 := [][]int{{1}}
	assert.Equal(t, expected2, levelOrder(root2), "case 2 failed")

	// 用例3：root = []
	var root3 *TreeNode = nil
	expected3 := [][]int{}
	assert.Equal(t, expected3, levelOrder(root3), "case 3 failed")

	// 用例4：完全二叉树 root = [1,2,3,4,5,6,7]
	root4 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7},
		},
	}
	expected4 := [][]int{{1}, {2, 3}, {4, 5, 6, 7}}
	assert.Equal(t, expected4, levelOrder(root4), "case 4 failed")

	// 用例5：只有左子树的情况 root = [1,2,3,4]
	root5 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 4},
			},
		},
	}
	expected5 := [][]int{{1}, {2}, {3}, {4}}
	assert.Equal(t, expected5, levelOrder(root5), "case 5 failed")

	// 用例6：只有右子树的情况 root = [1,null,2,null,3,null,4]
	root6 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val:   3,
				Right: &TreeNode{Val: 4},
			},
		},
	}
	expected6 := [][]int{{1}, {2}, {3}, {4}}
	assert.Equal(t, expected6, levelOrder(root6), "case 6 failed")

	// 用例7：负数值的情况 root = [-10,-3,0,5,null,-4,null,null,9]
	root7 := &TreeNode{
		Val: -10,
		Left: &TreeNode{
			Val:   -3,
			Left:  &TreeNode{Val: 5},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   0,
			Left:  &TreeNode{Val: -4},
			Right: nil,
		},
	}
	expected7 := [][]int{{-10}, {-3, 0}, {5, -4}}
	assert.Equal(t, expected7, levelOrder(root7), "case 7 failed")

	// 用例8：单层多节点 root = [1,2,3,4,5,6,7,8,9]
	root8 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 8},
				Right: &TreeNode{Val: 9},
			},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7},
		},
	}
	expected8 := [][]int{{1}, {2, 3}, {4, 5, 6, 7}, {8, 9}}
	assert.Equal(t, expected8, levelOrder(root8), "case 8 failed")
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
type treeQueue []*TreeNode

func (q *treeQueue) Enqueue(x *TreeNode) {
	*q = append(*q, x)
}

func (q *treeQueue) Dequeue() *TreeNode {
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func (q *treeQueue) Size() int {
	return len(*q)
}

func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	q := treeQueue{}
	q.Enqueue(root)

	for q.Size() > 0 {
		size := q.Size()
		tmp := make([]int, 0, size)
		for size > 0 {
			node := q.Dequeue()
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				q.Enqueue(node.Left)
			}
			if node.Right != nil {
				q.Enqueue(node.Right)
			}
			size--
		}
		ans = append(ans, tmp)
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
