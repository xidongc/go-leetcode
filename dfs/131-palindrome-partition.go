package main

import "math"

// 131 palindrome partition I with dfs
func partition(s string) [][]string {
	result := make([][]string, 0)
	PartitionHelper([]byte(s), 0, &result, []string{})
	return result
}

func PartitionHelper(input []byte, start int, result *[][]string, current []string) {
	if start == len(input) {
		tmp := make([]string, len(current), len(current))
		copy(tmp, current)
		*result = append(*result, tmp)
	}

	for i := start+1; i <= len(input); i ++ {
		if !isPalindrome(input[start: i]) {
			continue
		}
		current = append(current, string(input[start: i]))
		PartitionHelper(input, i, result, current)
		current = current[:len(current)-1]
	}
}

// determine if input is palindrome
func isPalindrome(input []byte) bool {
	i := 0
	j := len(input) - 1
	for i < j {
		if input[i] == input[j] {
			i ++
			j --
		} else {
			return false
		}
	}
	return true
}

// 132 palindrome partition II return dfs with memory
func minCut(s string) int {
	return minCutHelper([]byte(s), 0, map[int]int{})
}

func minCutHelper(input []byte, start int, memo map[int]int) int {
	if start == len(input) || isPalindrome(input[start:]) {
		return 0
	}
	if val, exist := memo[start]; exist {
		return val
	}
	result := math.MaxInt64 // assume int is running on 64 bit system
	for i := start+1; i <= len(input); i++ {
		if !isPalindrome(input[start:i]) {
			continue
		}
		tmpVal := minCutHelper(input, i, memo)
		if tmpVal + 1 < result {
			result = tmpVal + 1
		}
	}
	memo[start] = result
	return result
}
