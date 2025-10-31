package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给你两棵二叉树： root1 和 root2 。
//
// 想象一下，当你将其中一棵覆盖到另一棵之上时，两棵树上的一些节点将会重叠（而另一些不会）。你需要将这两棵树合并成一棵新二叉树。合并的规则是：如果两个节点重叠
//，那么将这两个节点的值相加作为合并后节点的新值；否则，不为 null 的节点将直接作为新二叉树的节点。
//
// 返回合并后的二叉树。
//
// 注意: 合并过程必须从两个树的根节点开始。
//
//
//
// 示例 1：
//
//
//输入：root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
//输出：[3,4,5,5,4,null,7]
//
//
// 示例 2：
//
//
//输入：root1 = [1], root2 = [1,2]
//输出：[2,2]
//
//
//
//
// 提示：
//
//
// 两棵树中的节点数目在范围 [0, 2000] 内
// -10⁴ <= Node.val <= 10⁴
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1484 👎 0

func Test617(t *testing.T) {
	// 测试用例1：示例1中的两棵树
	// root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	root2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	// 期望的合并结果 [3,4,5,5,4,null,7]
	expected := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	result := mergeTrees(root1, root2)
	assert.Equal(t, expected, result, "合并树结果应该符合预期")

	// 测试用例2：示例2中的两棵树
	// root1 = [1], root2 = [1,2]
	root3 := &TreeNode{Val: 1}
	root4 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}

	// 期望的合并结果 [2,2]
	expected2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 2,
		},
	}

	result2 := mergeTrees(root3, root4)
	assert.Equal(t, expected2, result2, "合并树结果应该符合预期")

	// 测试用例3：一棵树为空
	var root5 *TreeNode = nil
	root6 := &TreeNode{Val: 5}

	result3 := mergeTrees(root5, root6)
	assert.Equal(t, root6, result3, "当一棵树为空时，应返回另一棵树")

	// 测试用例4：两棵树都为空
	var root7, root8 *TreeNode = nil, nil
	result4 := mergeTrees(root7, root8)
	assert.Nil(t, result4, "当两棵树都为空时，应返回nil")
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
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	} else if root1 == nil {
		return root2
	} else if root2 == nil {
		return root1
	} else {
		root1.Val += root2.Val
	}

	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

//leetcode submit region end(Prohibit modification and deletion)
