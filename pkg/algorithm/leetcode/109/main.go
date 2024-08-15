package main

//给定一个单链表的头节点 head ，其中的元素 按升序排序 ，将其转换为 平衡 二叉搜索树。
//
//
//
// 示例 1:
//
//
//
//
//输入: head = [-10,-3,0,5,9]
//输出: [0,-3,9,-10,null,5]
//解释: 一个可能的答案是[0，-3,9，-10,null,5]，它表示所示的高度平衡的二叉搜索树。
//
//
// 示例 2:
//
//
//输入: head = []
//输出: []
//
//
//
//
// 提示:
//
//
// head 中的节点数在[0, 2 * 10⁴] 范围内
// -10⁵ <= Node.val <= 10⁵
//
//
// Related Topics 树 二叉搜索树 链表 分治 二叉树 👍 909 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func sortedListToBST(head *ListNode) (ans *TreeNode) {
	ans = new(TreeNode)
	nums := make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	if len(nums) == 0 {
		return nil
	}
	buildTree(nums, 0, len(nums), ans, nil)
	return
}

func buildTree(nums []int, i, j int, node, p *TreeNode) {
	if i >= j {
		if p == nil {
			return
		} else if p.Left == node {
			p.Left = nil
		} else {
			p.Right = nil
		}
		return
	}

	m := i + (j-i)/2
	node.Val = nums[m]
	node.Left = new(TreeNode)
	node.Right = new(TreeNode)

	buildTree(nums, i, m, node.Left, node)
	buildTree(nums, m+1, j, node.Right, node)
	return
}
