package main

import (
	"math"
	"sort"
	"strconv"
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

// 18 four sum: using hash map
type Index struct {
	i int
	j int
}

func fourSum(nums []int, target int) [][]int {
	hashmap := map[int][]Index{}
	for i := 0; i < len(nums) - 1; i ++ {
		for j := i+1; j < len(nums); j ++ {
			if _, exist := hashmap[nums[i]+nums[j]]; exist {
				hashmap[nums[i]+nums[j]] = append(hashmap[nums[i]+nums[j]], Index{
					i: i,
					j: j,
				})
			} else {
				hashmap[nums[i]+nums[j]] = []Index{
					{
						i: i,
						j: j,
					},
				}
			}
		}
	}

	result := make([][]int, 0)
	hashSet := map[string]struct{}{}

	for i := 0; i < len(nums) - 1; i ++ {
		for j := i+1; j < len(nums); j ++ {
			if val, exist := hashmap[target-nums[i]-nums[j]]; exist {
				for _, v := range val {
					if i != v.i && j != v.j && i != v.j && j != v.i {
						tmp := []int{nums[i], nums[j], nums[v.i], nums[v.j]}
						sort.Ints(tmp)
						hashCode := toStr(tmp)
						if _, exist := hashSet[hashCode]; !exist {
							result = append(result, tmp)
							hashSet[hashCode] = struct{}{}
						}
					}
				}
			}
		}
	}
	return result
}

func toStr(arr []int) string {
	s := ""
	for i, ele := range arr {
		s += strconv.Itoa(ele)
		for j := 0; j < i; j ++ {
			s += "a"
		}
	}
	return s
}