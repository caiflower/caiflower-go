package leetcode

import (
	"testing"
)

//给你二叉搜索树的根节点 root ，同时给定最小边界low 和最大边界 high。通过修剪二叉搜索树，使得所有节点的值在[low, high]中。修剪树 不
//应该 改变保留在树中的元素的相对结构 (即，如果没有被移除，原有的父代子代关系都应当保留)。 可以证明，存在 唯一的答案 。
//
// 所以结果应当返回修剪好的二叉搜索树的新的根节点。注意，根节点可能会根据给定的边界发生改变。
//
//
//
// 示例 1：
//
//
//输入：root = [1,0,2], low = 1, high = 2
//输出：[1,null,2]
//
//
// 示例 2：
//
//
//输入：root = [3,0,4,null,2,null,null,1], low = 1, high = 3
//输出：[3,2,null,1]
//
//
//
//
// 提示：
//
//
// 树中节点数在范围 [1, 10⁴] 内
// 0 <= Node.val <= 10⁴
// 树中每个节点的值都是 唯一 的
// 题目数据保证输入是一棵有效的二叉搜索树
// 0 <= low <= high <= 10⁴
//
//
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 1020 👎 0

func Test669(t *testing.T) {
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

// 递归法
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val >= low && root.Val <= high {
		root.Left = trimBST(root.Left, low, high)
		root.Right = trimBST(root.Right, low, high)
		return root
	} else if root.Val > high {
		return trimBST(root.Left, low, high)
	} else {
		return trimBST(root.Right, low, high)
	}
}

//type trimTreeQueue []*TreeNode
//
//func (q *trimTreeQueue) EnQueue(x *TreeNode) {
//	*q = append(*q, x)
//}
//
//func (q *trimTreeQueue) DeQueue() *TreeNode {
//	v := (*q)[0]
//	*q = (*q)[1:]
//	return v
//}
//
//func (q *trimTreeQueue) Size() int {
//	return len(*q)
//}
//
//// 迭代法
//func trimBST(root *TreeNode, low int, high int) *TreeNode {
//	if root == nil {
//		return nil
//	}
//
//	var ans *TreeNode
//	q := &trimTreeQueue{}
//	q.EnQueue(root)
//
//	for q.Size() > 0 {
//		r := q.DeQueue()
//		if r == nil {
//			continue
//		}
//
//		if r.Val >= low && r.Val <= high {
//			if ans == nil {
//				ans = r
//			}
//			if r.Left != nil && r.Left.Val < low {
//				r.Left = nil
//			} else {
//				q.EnQueue(r.Left)
//			}
//			if r.Right != nil && r.Right.Val > high {
//				r.Right = nil
//			} else {
//				q.EnQueue(r.Right)
//			}
//		} else if r.Val > high {
//			q.EnQueue(r.Left)
//		} else {
//			q.EnQueue(r.Right)
//		}
//
//	}
//
//	return ans
//}

//leetcode submit region end(Prohibit modification and deletion)
