package leetcode

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

//给你一份航线列表 tickets ，其中 tickets[i] = [fromi, toi] 表示飞机出发和降落的机场地点。请你对该行程进行重新规划排序。
//
//
// 所有这些机票都属于一个从 JFK（肯尼迪国际机场）出发的先生，所以该行程必须从 JFK 开始。如果存在多种有效的行程，请你按字典排序返回最小的行程组合。
//
//
//
// 例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"] 相比就更小，排序更靠前。
//
//
// 假定所有机票至少存在一种合理的行程。且所有的机票 必须都用一次 且 只能用一次。
//
//
//
// 示例 1：
//
//
//输入：tickets = [["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]
//输出：["JFK","MUC","LHR","SFO","SJC"]
//
//
// 示例 2：
//
//
//输入：tickets = [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL",
//"SFO"]]
//输出：["JFK","ATL","JFK","SFO","ATL","SFO"]
//解释：另一种有效的行程是 ["JFK","SFO","ATL","JFK","ATL","SFO"] ，但是它字典排序更大更靠后。
//
//
//
//
// 提示：
//
//
// 1 <= tickets.length <= 300
// tickets[i].length == 2
// fromi.length == 3
// toi.length == 3
// fromi 和 toi 由大写英文字母组成
// fromi != toi
//
//
// Related Topics 深度优先搜索 图 欧拉回路 👍 986 👎 0

func Test332(t *testing.T) {
	testCases := []struct {
		name    string
		tickets [][]string
		want    []string
	}{
		{
			name:    "case1",
			tickets: [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}},
			want:    []string{"JFK", "MUC", "LHR", "SFO", "SJC"},
		},
		{
			name:    "case2",
			tickets: [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}},
			want:    []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
		},
		{
			name:    "case3",
			tickets: [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "JFK"}, {"ATL", "AAA"}, {"AAA", "ATL"}, {"ATL", "BBB"}, {"BBB", "ATL"}, {"ATL", "CCC"}, {"CCC", "ATL"}, {"ATL", "DDD"}, {"DDD", "ATL"}, {"ATL", "EEE"}, {"EEE", "ATL"}, {"ATL", "FFF"}, {"FFF", "ATL"}, {"ATL", "GGG"}, {"GGG", "ATL"}, {"ATL", "HHH"}, {"HHH", "ATL"}, {"ATL", "III"}, {"III", "ATL"}, {"ATL", "JJJ"}, {"JJJ", "ATL"}, {"ATL", "KKK"}, {"KKK", "ATL"}, {"ATL", "LLL"}, {"LLL", "ATL"}, {"ATL", "MMM"}, {"MMM", "ATL"}, {"ATL", "NNN"}, {"NNN", "ATL"}},
			want:    []string{"JFK", "SFO", "JFK", "ATL", "AAA", "ATL", "BBB", "ATL", "CCC", "ATL", "DDD", "ATL", "EEE", "ATL", "FFF", "ATL", "GGG", "ATL", "HHH", "ATL", "III", "ATL", "JJJ", "ATL", "KKK", "ATL", "LLL", "ATL", "MMM", "ATL", "NNN", "ATL"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findItinerary(tc.tickets)
			assert.Equal(t, tc.want, result)
		})
	}
}

// leetcode submit region begin(Prohibit modification and deletion)

// N^N，超时
// 优化点在dfs函数中的for循环，可以优化为map，key为出发地，value为目的地，将时间复杂度降低到N
func findItinerary(tickets [][]string) []string {
	ans := []string{"JFK"}
	n := len(tickets) + 1
	used := make(map[int]bool, len(tickets))

	sort.Slice(tickets, func(i, j int) bool {
		if tickets[i][0] == tickets[j][0] {
			return tickets[i][1] < tickets[j][1]
		} else {
			return tickets[i][0] < tickets[j][0]
		}
	})

	var dfs func(location string) bool
	dfs = func(location string) bool {
		if len(ans) == n {
			return true
		}

		for i, v := range tickets {
			if !used[i] && location == v[0] {
				used[i] = true
				ans = append(ans, v[1])
				if dfs(v[1]) {
					return true
				}
				used[i] = false
				ans = ans[:len(ans)-1]
			}
		}

		return false
	}

	dfs("JFK")
	return ans
}

// 依旧超时
func findItinerary1(tickets [][]string) []string {
	ans := []string{"JFK"}
	n := len(tickets) + 1

	sort.Slice(tickets, func(i, j int) bool {
		if tickets[i][0] == tickets[j][0] {
			return tickets[i][1] < tickets[j][1]
		} else {
			return tickets[i][0] < tickets[j][0]
		}
	})
	graph := make(map[string][]string)
	graphUse := make(map[string][]bool)
	for _, v := range tickets {
		graph[v[0]] = append(graph[v[0]], v[1])
		graphUse[v[0]] = append(graphUse[v[0]], false)
	}

	var dfs func(location string) bool
	dfs = func(location string) bool {
		if len(ans) == n {
			return true
		}

		targets := graph[location]
		for i, v := range targets {
			if !graphUse[location][i] {
				graphUse[location][i] = true
				ans = append(ans, v)
				if dfs(v) {
					return true
				}
				graphUse[location][i] = false
				ans = ans[:len(ans)-1]
			}
		}

		return false
	}

	dfs("JFK")
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
