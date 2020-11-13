package main

// count of small numbers after self: using bisect, standard solution should use binary index tree
// time complexity O(n * log(n)) eg: input: nums = [5,2,6,1]; output: [2,1,1,0]
func countSmaller(nums []int) []int {
	result := make([]int, len(nums), len(nums))
	newArray := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i -- {
		insertPos := biInsert(nums[i], newArray)
		newArray = append(newArray, 0)
		copy(newArray[insertPos+1:], newArray[insertPos:])
		newArray[insertPos] = nums[i]
		result[i] = insertPos
	}
	return result
}

// helper func, return index where x should be inserted into array
// first pos of ele in array >= x
func biInsert(x int, array []int) int {
	if len(array) == 0 {
		return 0
	}
	start, end := 0, len(array) - 1
	for start < end - 1 {
		mid := start + (end - start) / 2
		if array[mid] < x {
			start = mid
		} else {
			end = mid
		}
	}
	if array[start] >= x {
		return start
	} else if array[end] >= x {
		return end
	} else {
		return len(array)
	}
}
