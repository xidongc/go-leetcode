package main

// 42 trap rain water I: using two pointers
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	res := 0
	maxLeft, maxRight := 0, 0
	i, j := 0, len(height) - 1
	for i < j {
		if height[i] < height[j] {
			if height[i] > maxLeft {
				maxLeft = height[i]
			} else {
				res += maxLeft - height[i]
			}
			i += 1
		} else {
			if height[j] > maxRight {
				maxRight = height[j]
			} else {
				res += maxRight - height[j]
			}
			j -= 1
		}
	}
	return res
}
