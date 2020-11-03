package main

import "github.com/xidongc/go-leetcode/utils"

// 643 max avg subarray: using pre-sum array
func findMaxAverage(nums []int, k int) float64 {
	if len(nums) < k || k == 0 {
		return 0.0
	}
	presum := make([]int, len(nums)+1, len(nums)+1)

	for i := 1; i <= k; i ++ {
		presum[i] = presum[i-1] + nums[i-1]
	}
	sumVal := presum[k] - presum[0]

	for i := k+1; i <= len(nums); i ++ {
		presum[i] = presum[i-1] + nums[i-1]
		sumVal = utils.Max(sumVal, presum[i] - presum[i-k])
	}
	return float64(sumVal) / float64(k)
}
