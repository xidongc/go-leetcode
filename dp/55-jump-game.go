package main

import (
	"github.com/xidongc/go-leetcode/utils"
	"math"
)

/*
	55 jump game
	Given an array of non-negative integers nums, you are initially positioned at the first index of the array.
	Each element in the array represents your maximum jump length at that position.
	Determine if you are able to reach the last index.

	co-ordinate dp, O(n*n) where dp[i] if you can reach i-th pos,
	dp[i] = dp[j] && OR(nums[j] + j >= i), dp[0] = True

Example 1:
Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:
Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
*/
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

/*
	45 jump game II
	Given an array of non-negative integers nums, you are initially positioned at the first index of the array.
	Each element in the array represents your maximum jump length at that position.
	Your goal is to reach the last index in the minimum number of jumps.
	You can assume that you can always reach the last index.

	dp[i] = Min(dp[j]+1 && nums[j] + j >= i && j < i)
	dp[i] = Integer.MaxValue if can't reach pos i
	dp[0] = 0

Example 1:
Input: nums = [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:
Input: nums = [2,3,0,1,4]
Output: 2
*/
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
				dp[i] = utils.Min(dp[i], dp[j]+1)
			}
		}
	}

	return dp[len(dp)-1]
}
