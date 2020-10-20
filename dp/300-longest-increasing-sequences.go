package main

// 300 lis using dp O(n*2)
// dp[i] = Max(dp[j] + 1) && j < i && nums[j] < nums[i] O(n*2)
// dp[0] = 1
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums), len(nums))
	for i := 0; i < len(dp); i ++ {
		dp[i] = 1
	}
	for i := 1; i < len(dp); i ++ {
		for j := 0; j < i; j ++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
	}

	lis := 0
	for i := 0; i < len(dp); i ++ {
		lis = max(lis, dp[i])
	}
	return lis
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

// 673 number of lis
// double dp
func findNumberOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums), len(nums))
	lisNum := make([]int, len(nums), len(nums))
	for i := 0; i < len(dp); i ++ {
		dp[i] = 1
		lisNum[i] = 1
	}
	for i := 1; i < len(dp); i ++ {
		for j := 0; j < i; j ++ {
			if nums[j] < nums[i] {
				if dp[j] + 1 > dp[i] {
					dp[i] = dp[j] + 1
					lisNum[i] = lisNum[j]
				} else if dp[j] + 1 == dp[i] {
					lisNum[i] = lisNum[i] + lisNum[j]
				}
			}
		}
	}

	count := 0
	val := 0
	for i := 0; i < len(dp); i ++ {
		if count < dp[i] {
			count = dp[i]
			val = lisNum[i]
		} else if count == dp[i] {
			val += lisNum[i]
		}
	}
	return val
}
