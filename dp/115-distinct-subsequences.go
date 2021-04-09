package main

/*
	115. Distinct Subsequences
	Given two strings s and t, return the number of distinct subsequences of s which equals t.
	A string's subsequence is a new string formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ACE" is a subsequence of "ABCDE" while "AEC" is not).
	It's guaranteed the answer fits on a 32-bit signed integer.

	dp[i][j] = dp[i-1][j-1] + dp[i-1][j] if s[i] == t[j]
    dp[i][j] = dp[i-1][j] if s[i] != t[j]

Example 1:
Input: s = "rabbbit", t = "rabbit"
Output: 3
Explanation:
As shown below, there are 3 ways you can generate "rabbit" from S.
rabbbit
rabbbit
rabbbit

Example 2:
Input: s = "babgbag", t = "bag"
Output: 5
Explanation:
As shown below, there are 5 ways you can generate "bag" from S.
babgbag
babgbag
babgbag
babgbag
babgbag
*/
func numDistinct(s string, t string) int {
	dp := make([][]int, 0)
	for i := 0; i <= len(s); i ++ {
		tmp := make([]int, len(t)+1, len(t)+1)
		dp = append(dp, tmp)
	}

	for i := 0; i < len(s); i ++ {
		dp[i][0] = 1
	}

	for i := 1; i <= len(s); i ++ {
		for j := 1; j <= len(t); j ++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(s)][len(t)]
}
