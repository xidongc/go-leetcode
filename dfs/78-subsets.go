package main

import "sort"

// 78 subsets I
func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	subsetsHelper(nums, 0, &result, []int{})
	return result
}

func subsetsHelper(nums []int, pos int, result *[][]int, current []int) {
	tmp := make([]int, len(current), len(current))
	copy(tmp, current)
	*result = append(*result, tmp)

	for i := pos; i < len(nums); i ++ {
		current = append(current, nums[i])
		subsetsHelper(nums, i+1, result, current)
		current = current[:len(current)-1]
	}
}

// 90 subsets II
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	subsetsUniqueHelper(nums, &result, []int{}, 0)
	return result
}

func subsetsUniqueHelper(nums []int, result *[][]int, current []int, start int) {
	tmp := make([]int, len(current), len(current))
	copy(tmp, current)
	*result = append(*result, tmp)

	for i := start; i < len(nums); i ++ {
		if i != start && nums[i] == nums[i-1] {
			continue
		}
		current = append(current, nums[i])
		subsetsUniqueHelper(nums, result, current, i+1)
		current = current[:len(current)-1]
	}
}