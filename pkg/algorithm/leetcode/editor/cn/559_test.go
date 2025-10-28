package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给定一个 N 叉树，找到其最大深度。
//
// 最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。
//
// N 叉树输入按层序遍历序列化表示，每组子节点由空值分隔（请参见示例）。
//
//
//
// 示例 1：
//
//
//
//
//输入：root = [1,null,3,2,4,null,5,6]
//输出：3
//
//
// 示例 2：
//
//
//
//
//输入：root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,
//null,13,null,null,14]
//输出：5
//
//
//
//
// 提示：
//
// 树的深度不会超过 1000 。
// 树的节点数目位于 [0, 10⁴] 之间。
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 👍 409 👎 0

func Test559(t *testing.T) {
	// 测试用例1：示例1中的树
	node5 := &Node{Val: 5}
	node6 := &Node{Val: 6}
	node3 := &Node{Val: 3, Children: []*Node{node5, node6}}
	node2 := &Node{Val: 2}
	node4 := &Node{Val: 4}
	root1 := &Node{Val: 1, Children: []*Node{node3, node2, node4}}

	assert.Equal(t, 3, maxDepth(root1), "示例1的最大深度应为3")

	assert.Equal(t, 0, maxDepth(nil), "空树的最大深度应为0")
	// 测试用例2：空树

	// 测试用例3：只有根节点的树
	root3 := &Node{Val: 1}
	assert.Equal(t, 1, maxDepth(root3), "只有根节点的树深度应为1")
}

type Node struct {
	Val      int
	Children []*Node
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
type NodeQueue []*Node

func (q *NodeQueue) Enqueue(x *Node) {
	*q = append(*q, x)
}

func (q *NodeQueue) Dequeue() *Node {
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func (q *NodeQueue) Size() int {
	return len(*q)
}

// 使用层序遍历
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	q := NodeQueue{}
	q.Enqueue(root)
	ans := 0

	for q.Size() > 0 {
		ans++
		size := q.Size()
		for size > 0 {
			n := q.Dequeue()
			for _, v := range n.Children {
				if v != nil {
					q.Enqueue(v)
				}
			}
			size--
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
