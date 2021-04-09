package main

import "github.com/xidongc/go-leetcode/utils"

/*
	72 Edit Distance
	Given two strings word1 and word2, return the minimum number
	of operations required to convert word1 to word2.
	You have the following three operations permitted on a word:
	Insert a character
	Delete a character
	Replace a character

	dp[i][j] = min(dp[i-1][j-1], dp[i][j-1] + 1, dp[i-1][j] + 1) if word1[i] == word2[j]
	dp[i][j] = min(dp[i-1][j-1] + 1, dp[i][j-1] + 1, dp[i-1][j] + 1) if word1[i] != word2[j]
	double sequence dp

Example 1:
Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation:
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')

Example 2:
Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation:
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')
*/
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, 0)
	for i := 0; i <= len(word1); i ++ {
		tmp := make([]int, len(word2)+1, len(word2)+1)
		dp = append(dp, tmp)
	}

	for i := 0; i <= len(word1); i ++ {
		dp[i][0] = i
	}

	for j := 1; j <= len(word2); j ++ {
		dp[0][j] = j
	}

	for i := 1; i <= len(word1); i ++ {
		for j := 1; j <= len(word2); j ++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = utils.Min(dp[i-1][j-1], dp[i][j-1] + 1, dp[i-1][j] + 1)
			} else{
				dp[i][j] = utils.Min(dp[i-1][j-1] + 1, dp[i][j-1] + 1, dp[i-1][j] + 1)
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
