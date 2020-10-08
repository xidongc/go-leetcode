package main

// search in first position >= target
func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start < end - 1 {
		mid := start + (end - start) / 2
		if nums[mid] == target {
			end = mid
		} else if nums[mid] > target {
			end = mid
		} else {
			start = mid
		}
	}
	if nums[start] >= target{
		return start
	}
	if nums[end] >= target {
		return end
	}
	return len(nums)
}
