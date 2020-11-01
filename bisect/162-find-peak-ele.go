package main

// 163 find peak element: using bisect
func findPeakElement(nums []int) int {
	start, end := 0, len(nums)-1
	for start < end - 1 {
		mid := start + (end - start) / 2
		if nums[mid] < nums[mid-1] {
			end = mid - 1
		} else if nums[mid] > nums[mid-1] {
			start = mid
		}
	}
	if start > 0 && start < len(nums)-1 && nums[start] > nums[start-1] && nums[start] > nums[start+1] {
		return start
	} else if end > 0 && end < len(nums)-1 && nums[end] > nums[end - 1] && nums[end] > nums[end+1] {
		return end
	} else {
		if nums[start] > nums[end] {
			return start
		} else {
			return end
		}
	}
}
