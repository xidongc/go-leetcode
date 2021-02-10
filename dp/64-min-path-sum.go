package main

import "github.com/xidongc/go-leetcode/utils"

/*
	64 Minimum Path Sum
	Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right,
	which minimizes the sum of all numbers along its path.
	Note: You can only move either down or right at any point in time.

	co-ordinate dp, dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j], O(n)

Example 1:
Input: grid = [[1,3,1],[1,5,1],[4,2,1]]
Output: 7
Explanation: Because the path 1 → 3 → 1 → 1 → 1 minimizes the sum.

Example 2:
Input: grid = [[1,2,3],[4,5,6]]
Output: 12
*/
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	dp := make([][]int, 0)
	for i := 0; i < len(grid); i ++ {
		tmp := make([]int, len(grid[0]), len(grid[0]))
		dp = append(dp, tmp)
		if i == 0 {
			dp[i][0] = grid[i][0]
		} else {
			dp[i][0] = grid[i][0] + dp[i-1][0]
		}
	}

	for j := 1; j < len(grid[0]); j ++ {
		dp[0][j] = grid[0][j] + dp[0][j-1]
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			dp[i][j] = utils.Min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

/*
	62 unique path
	A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
	The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
	How many possible unique paths are there?

	dp[i][j] = dp[i-1][j] + dp[i][j-1]
*/
func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	dp := make([][]int, 0)
	for i := 0; i < m; i ++ {
		tmp := make([]int, n, n)
		tmp[0] = 1
		dp = append(dp, tmp)
	}

	for j := 1; j < n; j ++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i ++ {
		for j := 1; j < n; j ++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

// 63 unique paths II
// dp[i][j] = dp[i-1][j] + dp[i][j-1] if obstacleGrid[i][j] == 0 else 0
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0])==0 {
		return 0
	}
	dp := make([][]int, 0)
	for i := 0; i < len(obstacleGrid); i ++ {
		tmp := make([]int, len(obstacleGrid[0]), len(obstacleGrid[0]))
		dp = append(dp, tmp)
	}

	dp[0][0] = 1
	if obstacleGrid[0][0] == 1 {
		dp[0][0] = 0
	}
	for i := 1; i < len(dp); i ++ {
		dp[i][0] = dp[i-1][0]
		if obstacleGrid[i][0] == 1 {
			dp[i][0] = 0
		}
	}

	for j := 1; j < len(dp[0]); j ++ {
		dp[0][j] = dp[0][j-1]
		if obstacleGrid[0][j] == 1 {
			dp[0][j] = 0
		}
	}

	for i := 1; i < len(dp); i ++ {
		for j := 1; j < len(dp[0]); j ++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}
