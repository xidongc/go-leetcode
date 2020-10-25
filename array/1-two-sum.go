package main

import (
	"math"
	"sort"
)

// 1 two sum: O(n) with hashMap
func twoSum(nums []int, target int) []int {
	hashMap := map[int]int{}
	for i := 0 ; i < len(nums); i ++ {
		if val, exist := hashMap[target - nums[i]]; exist {
			return []int{val, i}
		}
		hashMap[nums[i]] = i
	}
	return []int{}
}

// 15 three sum with duplicates: O(n*n) using three pointer
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	results := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums) - 2; i ++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums) - 1
		for j < k {
			tmp := nums[i] + nums[j] + nums[k]
			if tmp < 0 {
				j ++
			} else if tmp > 0 {
				k --
			} else {
				results = append(results, []int{nums[i], nums[j], nums[k]})
				j ++
				k --
				for j < k && nums[j] == nums[j-1] {
					j ++
				}
				for j < k && nums[k] == nums[k+1] {
					k --
				}
			}
		}
	}
	return results
}

// 16 three sum closest
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closest := math.MaxInt64
	result := 0
	i := 0
	for i < len(nums) - 2 {
		if i > 0 && nums[i] == nums[i-1] {
			i ++
		}
		j := i + 1
		k := len(nums)-1
		for j < k {
			tmp := nums[i] + nums[j] + nums[k]
			if tmp < target {
				if target - tmp < closest {
					result = tmp
					closest = target - tmp
				}
				j ++
			} else if tmp > target {
				if tmp - target < closest {
					result = tmp
					closest = tmp - target
				}
				k --
			} else {
				return tmp
			}
			for j < k && j > i+1 && nums[j] == nums[j-1] {
				j ++
			}
			for j < k && k < len(nums)-1 && nums[k] == nums[k+1] {
				k --
			}
		}
		i ++
	}
	return result
}
