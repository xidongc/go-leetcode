package main

import (
	"github.com/xidongc/go-leetcode/utils"
	"math"
)

// 53 max sub array using pre sum array O(n)
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		panic("length input equals zero")
	}
	presum := make([]int, len(nums)+1, len(nums)+1)
	valley := 0
	maxVal := math.MinInt64 // assume 64 bit server
	for i := 1; i <= len(nums); i ++ {
		presum[i] = presum[i-1] + nums[i-1]
		maxVal = utils.Max(maxVal, presum[i]-valley)
		valley = utils.Min(valley, presum[i])
	}
	return maxVal
}

// 1031 maximum sum of two non-overlapping sub-arrays
// separate given array into two parts, and for each,
// calculate maximum of subarray (53 above). O(n)
func maxSumTwoNoOverlap(A []int, L int, M int) int {
	return utils.Max(maxSumTwo(A, L, M), maxSumTwo(A, M, L))
}

func maxSumTwo(A []int, L int, M int) int {
	if len(A) < M + L || (L == 0 && M == 0){
		return 0 // not found
	}
	preSumLeft := make([]int, len(A)+1)
	// first L ele pre sum
	for i := 1; i <= L; i ++ {
		preSumLeft[i] = preSumLeft[i-1] + A[i-1]
	}
	left := make([]int, len(A)-L-M+1, len(A)-L-M+1)
	left[0] = preSumLeft[L] - preSumLeft[0]

	for i := L; i < len(A)-M; i ++ {
		preSumLeft[i+1] = preSumLeft[i] + A[i]
		left[i-L+1] = utils.Max(left[i-L], preSumLeft[i+1]-preSumLeft[i+1-L])
	}
	right := make([]int, len(A)-L-M+1, len(A)-L-M+1)
	preSumRight := make([]int, len(A)+1, len(A)+1)

	for i := len(A); i > len(A)-M; i -- {
		preSumRight[i-1] = preSumRight[i] + A[i-1]
	}
	right[len(A)-M-L] = preSumRight[len(A)-M] - preSumRight[len(A)]
	for i := len(A)-M; i > L; i -- {
		preSumRight[i-1] = preSumRight[i] + A[i-1]
		right[i-L-1] = utils.Max(right[i-L], preSumRight[i-1] - preSumRight[i+M-1])
	}

	maxValue := math.MinInt64
	for i := 0; i <= len(A)-M-L; i ++ {
		maxValue = utils.Max(maxValue, left[i]+right[i])
	}
	return maxValue
}
