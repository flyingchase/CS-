package IntroductionToPath

import (
	"math"
)

func HminFallingPathSum(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	copy(dp[0], grid[0])
	min := func(j int, i int) int {
		if i > j {
			return j
		}
		return i
	}
	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			dp[i][j] = math.MaxInt32
			val := grid[i][j]
			for k := 0; k < m; k++ {
				if j != k {
					dp[i][j] = min(dp[i][j], dp[i-1][k]+val)
				}
			}
		}
	}
	res := math.MaxInt32
	for i := 0; i < m; i++ {
		res = min(res, dp[n-1][i])
	}
	return res
}

