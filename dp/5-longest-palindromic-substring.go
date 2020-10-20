package main

import (
	"fmt"
	"math"
)

// range dp O(n*2) used in palindrome partition II as well
// loop sequence matters, first loop length, then loop start
func longestPalindrome(s string) string {
	dp := make([][]bool, 0)
	for i := 0 ; i < len(s); i ++ {
		tmp := make([]bool, len(s), len(s))
		dp = append(dp, tmp)
	}
	dp[0][0] = true
	for i := 1; i < len(s); i ++ {
		dp[i][i] = true
		if s[i-1] == s[i] {
			dp[i-1][i] = true
		}
	}
	for length := 2; length < len(s); length ++ {
		for start := 0; start < len(s) - length; start ++ {
			if s[start] == s[start+length] && dp[start+1][start+length-1] == true{
				dp[start][start+length] = true
			}
		}
	}

	longestVal := 0
	start := 0
	end := 0
	for i := 0; i < len(s); i ++ {
		for j := i+1; j < len(s); j ++ {
			if dp[i][j] ==true && j - i > longestVal {
				longestVal = j - i
				start = i
				end = j
			}
		}
	}
	return s[start:end+1]
}

// 132 palindrome partitioning II with dp O(n*2)
// O(i) = Min(O(j) + 1) && j < i if s[j:i] == palindrome
func minCut(s string) int {
	dp := make([]int, len(s)+1, len(s)+1)
	for i := 1; i < len(s)+1; i ++ {
		dp[i] = math.MaxInt64
	}
	palMatrix := isPalindrome(s)
	for i := 1; i <= len(s); i ++ {
		for j := 0; j < i; j ++ {
			if palMatrix[j][i-1] {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	fmt.Println(dp)
	return dp[len(s)]-1 // as cut == partition numbers - 1
}

// O(n*2) dp[i][j] represent s[i:j+1] is a palindrome or not
func isPalindrome(s string) [][]bool {
	dp := make([][]bool, 0)
	for i := 0; i < len(s); i ++ {
		tmp := make([]bool, len(s), len(s))
		dp = append(dp, tmp)
	}
	dp[0][0] = true
	for i := 1; i < len(s); i ++ {
		dp[i][i] = true
		if s[i-1] == s[i] {
			dp[i-1][i] = true
		}
	}

	for length := 2; length < len(s); length ++ {
		for start := 0; start < len(s) - length; start ++ {
			if dp[start+1][start+length-1] && s[start] == s[start+length] {
				dp[start][start+length] = true
			}
		}
	}
	return dp
}
