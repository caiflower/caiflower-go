package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。
//
// 叶子节点 是指没有子节点的节点。
//
// 示例 1：
//
//
//输入：root = [1,2,3,null,5]
//输出：["1->2->5","1->3"]
//
//
// 示例 2：
//
//
//输入：root = [1]
//输出：["1"]
//
//
//
//
// 提示：
//
//
// 树中节点的数目在范围 [1, 100] 内
// -100 <= Node.val <= 100
//
// Related Topics 树 深度优先搜索 字符串 回溯 二叉树 👍 1257 👎 0
//

func Test257(t *testing.T) {
	// 测试用例1：[1,2,3,null,5]
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	expected1 := []string{"1->2->5", "1->3"}
	assert.ElementsMatch(t, expected1, binaryTreePaths(root1))

	// 测试用例2：[1]
	root2 := &TreeNode{
		Val: 1,
	}
	expected2 := []string{"1"}
	assert.ElementsMatch(t, expected2, binaryTreePaths(root2))

	// 测试用例3：空树
	var root3 *TreeNode = nil
	expected3 := []string{}
	assert.ElementsMatch(t, expected3, binaryTreePaths(root3))
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
func binaryTreePaths(root *TreeNode) []string {
	ans := []string{}
	treePaths(root, &ans, "")
	return ans
}

func treePaths(root *TreeNode, ans *[]string, cur string) {
	if root == nil {
		return
	}

	if cur != "" {
		cur += "->"
	}
	cur += fmt.Sprintf("%d", root.Val)

	if root.Left == nil && root.Right == nil {
		if cur != "" {
			*ans = append(*ans, cur)
		}
		return
	}
	treePaths(root.Left, ans, cur)
	treePaths(root.Right, ans, cur)
}

//leetcode submit region end(Prohibit modification and deletion)
