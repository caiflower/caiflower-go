package main

func numberOfDice(n int) []int {
	dp := make([][]int, n+1)
	dp[0] = []int{1, 0, 0, 0, 0, 0, 0}

	// 枚举次数
	for i := 1; i <= n; i++ {
		t := make([]int, n*6+1)
		// 枚举点数
		for k := i; k <= i*6; k++ {
			// 枚举骰子
			for j := 1; j <= 6; j++ {
				if k-j >= 0 {
					t[k] += dp[i-1][k-j]
				}
			}
		}
		dp[i] = t
	}

	return dp[n][n:]
}
