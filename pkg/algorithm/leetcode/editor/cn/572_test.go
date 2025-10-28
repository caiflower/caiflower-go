package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//
//
// 给你两棵二叉树 root 和 subRoot 。检验 root 中是否包含和 subRoot 具有相同结构和节点值的子树。如果存在，返回 true ；否则
//，返回 false 。
//
//
//
// 二叉树 tree 的一棵子树包括 tree 的某个节点和这个节点的所有后代节点。tree 也可以看做它自身的一棵子树。
//
//
//
// 示例 1：
//
//
//输入：root = [3,4,5,1,2], subRoot = [4,1,2]
//输出：true
//
//
// 示例 2：
//
//
//输入：root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
//输出：false
//
//
//
//
// 提示：
//
//
// root 树上的节点数量范围是 [1, 2000]
// subRoot 树上的节点数量范围是 [1, 1000]
// -10⁴ <= root.val <= 10⁴
// -10⁴ <= subRoot.val <= 10⁴
//
// Related Topics 树 深度优先搜索 二叉树 字符串匹配 哈希函数 👍 1162 👎 0
//

func Test572(t *testing.T) {
	// 测试用例1：示例1
	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 2},
		},
		Right: &TreeNode{Val: 5},
	}

	subRoot1 := &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 2},
	}
	assert.True(t, isSubtree(root1, subRoot1), "示例1应返回true")

	// 测试用例2：示例2
	root2 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  4,
			Left: &TreeNode{Val: 1},
			Right: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 0},
			},
		},
		Right: &TreeNode{Val: 5},
	}

	subRoot2 := &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 2},
	}
	assert.False(t, isSubtree(root2, subRoot2), "示例2应返回false")

	// 测试用例3：空子树
	root3 := &TreeNode{Val: 1}
	var subRoot3 *TreeNode = nil

	assert.True(t, isSubtree(root3, subRoot3), "空子树应返回true")

	// 测试用例4：相同的树
	root4 := &TreeNode{Val: 1}
	subRoot4 := &TreeNode{Val: 1}

	assert.True(t, isSubtree(root4, subRoot4), "相同的树应返回true")

	// 测试用例5：子树在右侧
	root5 := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 1},
		Right: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 5},
		},
	}

	subRoot5 := &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 5},
	}
	assert.True(t, isSubtree(root5, subRoot5), "右侧子树应返回true")

	// 测试用例6：子树在深层次
	root6 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 5},
			},
		},
	}

	subRoot6 := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 4},
		Right: &TreeNode{Val: 5},
	}
	assert.True(t, isSubtree(root6, subRoot6), "深层子树应返回true")

	// 测试用例7：结构相同但值不同
	root7 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}

	subRoot7 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 4}, // 值不同
	}
	assert.False(t, isSubtree(root7, subRoot7), "结构相同但值不同应返回false")

	// 测试用例8：空主树
	var root8 *TreeNode = nil
	subRoot8 := &TreeNode{Val: 1}
	assert.False(t, isSubtree(root8, subRoot8), "空主树非空子树应返回false")

	// 测试用例9：两棵空树
	var root9 *TreeNode = nil
	var subRoot9 *TreeNode = nil
	assert.True(t, isSubtree(root9, subRoot9), "两棵空树应返回true")

	// [3,1,2,4,5,null,null,1,5,null,null,6,5,null,null]
	// [3,1,2,6,5]
	// 测试用例10：[3,1,2,4,5,null,null,1,5,null,null,6,5,null,null] 和 [3,1,2,6,5]
	root10 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 5}},
				Right: &TreeNode{Val: 5},
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{Val: 2},
	}

	subRoot10 := &TreeNode{
		Val: 3,
		Left: &TreeNode{Val: 1,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 5}},
		Right: &TreeNode{
			Val: 2,
		},
	}

	assert.False(t, isSubtree(root10, subRoot10), "用例10应返回false")
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

// 递归所有节点，判断以该节点为根的子树是否与subRoot相同
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root != nil && isSubtree(root.Left, subRoot) {
		return true
	} else if root != nil && isSubtree(root.Right, subRoot) {
		return true
	} else {
		return isSame(root, subRoot)
	}
}

func isSame(i, j *TreeNode) bool {
	if i == nil && j == nil {
		return true
	} else if i == nil || j == nil {
		return false
	} else {
		if i.Val != j.Val {
			return false
		} else if !isSame(i.Left, j.Left) {
			return false
		} else {
			return isSame(i.Right, j.Right)
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
