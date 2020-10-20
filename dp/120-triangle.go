package main

import (
	"math"
)

// 120 triangle with dp
// dp[i][i] = Min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
// dp[i][0] = triangle[i-1][0] + dp[i][0]
// dp[i][i] = triangle[i][i] + dp[i-1][i-1]
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	dp := make([][]int, 0)
	for i := 0;  i < len(triangle); i ++ {
		tmp := make([]int, i+1, i+1)
		dp = append(dp, tmp)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i ++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}

	for i := 1; i < len(dp); i ++ {
		for j := 1; j < i; j ++ {
			dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
		}
	}
	max := math.MaxInt64
	for _, val := range dp[len(dp)-1] {
		max = min(val, max)
	}

	return max
}
