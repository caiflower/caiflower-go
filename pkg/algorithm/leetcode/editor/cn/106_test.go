package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并
//返回这颗 二叉树 。
//
//
//
// 示例 1:
//
//
//输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
//输出：[3,9,20,null,null,15,7]
//
//
// 示例 2:
//
//
//输入：inorder = [-1], postorder = [-1]
//输出：[-1]
//
//
//
//
// 提示:
//
//
// 1 <= inorder.length <= 3000
// postorder.length == inorder.length
// -3000 <= inorder[i], postorder[i] <= 3000
// inorder 和 postorder 都由 不同 的值组成
// postorder 中每一个值都在 inorder 中
// inorder 保证是树的中序遍历
// postorder 保证是树的后序遍历
//
//
// Related Topics 树 数组 哈希表 分治 二叉树 👍 1366 👎 0

func Test106(t *testing.T) {
	//测试用例1
	inorder1 := []int{9, 3, 15, 20, 7}
	postorder1 := []int{9, 15, 7, 20, 3}
	result1 := buildTree(inorder1, postorder1)
	// 验证根节点
	assert.Equal(t, 3, result1.Val)
	// 验证左子树
	assert.Equal(t, 9, result1.Left.Val)
	assert.Nil(t, result1.Left.Left)
	assert.Nil(t, result1.Left.Right)
	// 验证右子树
	assert.Equal(t, 20, result1.Right.Val)
	assert.Equal(t, 15, result1.Right.Left.Val)
	assert.Equal(t, 7, result1.Right.Right.Val)

	// 测试用例2
	inorder2 := []int{-1}
	postorder2 := []int{-1}
	result2 := buildTree(inorder2, postorder2)
	assert.Equal(t, -1, result2.Val)
	assert.Nil(t, result2.Left)
	assert.Nil(t, result2.Right)

	// 测试用例3
	inorder3 := []int{2, 3, 1}
	postorder3 := []int{3, 2, 1}
	result3 := buildTree(inorder3, postorder3)
	// 验证根节点
	assert.Equal(t, 1, result3.Val)
	// 验证左子树
	assert.Equal(t, 2, result3.Left.Val)
	assert.Nil(t, result3.Left.Left)
	assert.Equal(t, 3, result3.Left.Right.Val)
	// 验证右子树
	assert.Nil(t, result3.Right)
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

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	pRootIndex, iRootIndex := len(postorder)-1, 0
	rootVal := postorder[pRootIndex]
	r := &TreeNode{Val: rootVal}

	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			iRootIndex = i
			break
		}
	}

	left := inorder[:iRootIndex]
	r.Left = buildTree(left, postorder[:len(left)])
	r.Right = buildTree(inorder[iRootIndex+1:], postorder[len(left):pRootIndex])

	return r
}

//leetcode submit region end(Prohibit modification and deletion)
