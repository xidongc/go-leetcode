package main

import "math"

// 55 jump game I
// dp[i] = OR(nums[j] + j >= i && dp[j] && j < i)
// dp[0] = true
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	dp := make([]bool, len(nums), len(nums))
	dp[0] = true
	for i := 1; i < len(dp); i ++ {
		for j := 0; j < i; j ++ {
			if dp[j] && nums[j] + j >= i {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(dp)-1]
}

// 45 jump game II
// dp[i] = Min(dp[j]+1 && nums[j] + j >= i && j < i)
// dp[i] = Integer.MaxValue if can't reach pos i
// dp[0] = 0
func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	dp := make([]int, len(nums), len(nums))
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt64 // assume 64 bit server
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if dp[j] != math.MaxInt64 && nums[j] + j >= i {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}

	return dp[len(dp)-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}


