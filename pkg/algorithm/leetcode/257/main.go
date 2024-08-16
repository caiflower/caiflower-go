package main

import (
	"fmt"
	"strconv"
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
//
// Related Topics 树 深度优先搜索 字符串 回溯 二叉树 👍 1156 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) (ans []string) {
	if root == nil {
		return
	}
	deep(root, "", &ans)
	return
}

func deep(root *TreeNode, cur string, ans *[]string) {
	if cur != "" {
		cur += "->"
	}
	cur += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		*ans = append(*ans, cur)
		return
	}
	if root.Left != nil {
		deep(root.Left, cur, ans)
	}
	if root.Right != nil {
		deep(root.Right, cur, ans)
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(binaryTreePaths(&TreeNode{Val: 1}))
}
