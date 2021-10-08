package IntroductionToPath

import "math"

// 对于第一行其最小路径即为自身存储的值
// 第一行作为初始状态，转移方程为左上角，正上方，右上方
func minFallingPathSum(matrix [][]int) int {

	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	// 首行即为初始条件
	for i := 0; i < n; i++ {
		dp[0][i] = matrix[0][i]
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			val := matrix[i][j]
			dp[i][j] = dp[i-1][j] + val
			if j-1 >= 0 {
				dp[i][j] = min931(dp[i][j], dp[i-1][j-1]+val)
			}
			if j+1 < n {
				dp[i][j] = min931(dp[i][j], dp[i-1][j+1]+val)
			}
		}
	}
	res := math.MaxInt32
	for i := 0; i < n; i++ {
		res = min931(res, dp[n-1][i])
	}
	return res
}
func min931(j int, i int) int {
	if i > j {
		return j
	}
	return i
}
