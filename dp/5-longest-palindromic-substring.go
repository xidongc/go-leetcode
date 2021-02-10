package main

import (
	"math"
)

/*
5. Longest Palindromic Substring
   range dp O(n*2)
   loop sequence matters, first loop length, then loop start

Example 1:

Input: s = "babad"
Output: "bab"
Note: "aba" is also a valid answer.

Example 2:

Input: s = "cbbd"
Output: "bb"

Example 3:

Input: s = "a"
Output: "a"

Example 4:

Input: s = "ac"
Output: "a"
*/
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

/*
132 Palindrome Partitioning II
	Given a string s, partition s such that every substring of the partition is a palindrome.
	Return the minimum cuts needed for a palindrome partitioning of s.

	sequence which dp(i) represents given pre i element in str, how many cuts needed
	dp(i) = Min(dp(j) + 1) && j < i if s[j:i] is palindrome, O(n*2)

	when determine s[j:i], using above No.5 solution of range dp: O(n*2)

	tricky point is how to treat dp[0] as pre 0s ele could have -1 cuts

Example 1:
Input: s = "aab"
Output: 1
Explanation: The palindrome partitioning ["aa","b"] could be produced using 1 cut.

Example 2:
Input: s = "a"
Output: 0

Example 3:
Input: s = "ab"
Output: 1
*/
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
	return dp[len(s)]-1 // as cut == partition numbers - 1
}

/*
	Standard way to determine whether s[i: j+1] is a palindrome or not
	O(n*2), in loop, length goes first, start from 2
*/
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
