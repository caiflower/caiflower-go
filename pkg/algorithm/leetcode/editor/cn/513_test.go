package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test513(t *testing.T) {
	// 构建测试用例1: [2,1,3]
	root1 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}

	// 构建测试用例2: [1,2,3,4,null,5,6,null,null,7]
	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 7},
			},
			Right: &TreeNode{Val: 6},
		},
	}

	// 测试用例
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "示例1",
			root: root1,
			want: 1,
		},
		{
			name: "示例2",
			root: root2,
			want: 7,
		},
	}

	// 执行测试
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findBottomLeftValue(tt.root)
			assert.Equal(t, tt.want, got, "findBottomLeftValue(%v) = %v, want %v", tt.root, got, tt.want)
		})
	}
}

//给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
//
// 假设二叉树中至少有一个节点。
//
//
//
// 示例 1:
//
//
//
//
//输入: root = [2,1,3]
//输出: 1
//
//
// 示例 2:
//
//
//
//
//输入: [1,2,3,4,null,5,6,null,null,7]
//输出: 7
//
//
//
//
// 提示:
//
//
// 二叉树的节点个数的范围是 [1,10⁴]
//
// -2³¹ <= Node.val <= 2³¹ - 1
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 645 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 使用递归解决，比层序遍历稍微难写一些。层序遍历非常容易写
func findBottomLeftValue(root *TreeNode) int {
	ans, maxDepth := 0, 0
	bottomLeftValue(root, 1, &maxDepth, &ans)
	return ans
}

func bottomLeftValue(root *TreeNode, curDepth int, maxDepth, ans *int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil && *maxDepth < curDepth {
		*maxDepth = curDepth
		*ans = root.Val
		return
	}

	bottomLeftValue(root.Left, curDepth+1, maxDepth, ans)
	bottomLeftValue(root.Right, curDepth+1, maxDepth, ans)
}

//leetcode submit region end(Prohibit modification and deletion)
