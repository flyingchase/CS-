package IntroductionToPath

import "math"

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	res := math.MaxInt32
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		for j := 0; j < i+1; j++ {
			val := triangle[i][j]
			dp[i][j] = math.MaxInt32
			if j != 0 {
				dp[i][j] = min120(dp[i][j], dp[i-1][j-1]+val)
			}
			if i != j {
				dp[i][j] = min120(dp[i][j], dp[i-1][j]+val)
			}
		}
	}
	for i := 0; i < n; i++ {
		res = min120(res, dp[n-1][i])
	}
	return res
}
func min120(i, j int) int {
	if i < j {
		return i
	}
	return j
}
