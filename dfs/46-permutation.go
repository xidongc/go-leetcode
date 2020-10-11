package main

import (
	"sort"
)

// 46 permute I
func permute(nums []int) (results [][]int) {
	if len(nums) == 0 {
		return [][]int{}
	}
	visited := make([]bool, len(nums), len(nums))
	permuteHelper(nums, &results, []int{}, visited)
	return
}

func permuteHelper(nums []int, result *[][]int, current []int, used []bool) {
	if len(current) == len(nums) {
		tmp := make([]int, len(current), len(current))
		copy(tmp, current)
		*result = append(*result, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		current = append(current, nums[i])
		used[i] = true
		permuteHelper(nums, result, current, used)
		used[i] = false
		current = current[0: len(current)-1]
	}
}

// 47 permute II
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	sig := make([]bool, len(nums), len(nums))
	result := make([][]int, 0)
	permuteUniqueHelper(nums, &result, sig, []int{})
	return result
}

func permuteUniqueHelper(nums []int, result *[][]int, sig []bool, current []int) {
	if len(nums) == len(current) {
		tmp := make([]int, len(current), len(current))
		copy(tmp, current)
		*result = append(*result, tmp)
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && sig[i-1] == false {
			continue
		}
		if sig[i] == true{
			continue
		}
		current = append(current, nums[i])
		sig[i] = true
		permuteUniqueHelper(nums, result, sig, current)
		current = current[:len(current)-1]
		sig[i] = false
	}
}
