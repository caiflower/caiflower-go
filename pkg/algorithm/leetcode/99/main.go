package main

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

func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}
	var arr []*TreeNode
	preSearch(root, &arr)

	var n, m *TreeNode
	m = arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i].Val < arr[i-1].Val {
			if m.Val < arr[i].Val {
				m = arr[i-1]
			}
		}
	}

	for i := 1; i < len(arr); i++ {
		if arr[i].Val < arr[i-1].Val {
			if n == nil || n.Val > arr[i].Val {
				n = arr[i]
			}
		}
	}

	tmp := m.Val
	m.Val = n.Val
	n.Val = tmp
}

func preSearch(cur *TreeNode, array *[]*TreeNode) {
	if cur == nil {
		return
	}

	preSearch(cur.Left, array)
	*array = append(*array, cur)
	preSearch(cur.Right, array)
}

func main() {
	node1 := &TreeNode{Val: 2}
	node4 := &TreeNode{Val: 4, Left: node1}
	node2 := &TreeNode{Val: 1}
	node3 := &TreeNode{Val: 3, Left: node2, Right: node4}

	recoverTree(node3)
}
