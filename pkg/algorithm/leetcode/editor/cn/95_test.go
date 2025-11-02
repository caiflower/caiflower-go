package leetcode

import "testing"

//给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
//
//
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
//
//
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：[[1]]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 8
//
//
// Related Topics 树 二叉搜索树 动态规划 回溯 二叉树 👍 1641 👎 0

func Test95(t *testing.T) {
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
func generateTrees(n int) []*TreeNode {
	dp := make([][]*TreeNode, n+1)
	dp[0] = []*TreeNode{nil}
	dp[1] = []*TreeNode{{Val: 1}}

	for i := 2; i <= n; i++ {
		// 左子树个数
		for j := 0; j < i; j++ {
			// 左子树所有组合
			for _, v := range dp[j] {
				// 右边子树所有组合
				for _, v1 := range dp[i-1-j] {
					r := &TreeNode{Val: j + 1, Left: v}
					r.Right = buildRight(v1, j+1)
					dp[i] = append(dp[i], r)
				}
			}

		}
	}

	return dp[n]
}

func buildRight(node *TreeNode, add int) *TreeNode {
	if node == nil {
		return nil
	}
	t := &TreeNode{Val: node.Val + add, Left: buildRight(node.Left, add), Right: buildRight(node.Right, add)}
	return t
}

//leetcode submit region end(Prohibit modification and deletion)
