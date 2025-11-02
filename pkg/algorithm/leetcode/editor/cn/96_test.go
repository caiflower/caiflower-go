package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：5
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= n <= 19
//
//
// Related Topics 树 二叉搜索树 数学 动态规划 二叉树 👍 2692 👎 0

func Test96(t *testing.T) {
	assert.Equal(t, 1, numTrees(1), "numTrees(1) should be 1")
	assert.Equal(t, 2, numTrees(2), "numTrees(2) should be 0")
	assert.Equal(t, 5, numTrees(3), "numTrees(3) should be 3")
}

// leetcode submit region begin(Prohibit modification and deletion)
func numTrees(n int) int {
	// dp代表n个节点组成不相同的二叉搜索树的个数
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	// dp[n] = 左子树的二叉搜索树个数 * 右子树的二叉树个数，因为左右子树节点个数可能不同，所以要for循环枚举，然后加和
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}

	return dp[n]
}

//leetcode submit region end(Prohibit modification and deletion)
