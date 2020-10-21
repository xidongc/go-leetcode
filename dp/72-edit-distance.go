package main

// dp[i][j] = min(dp[i-1][j-1], dp[i][j-1] + 1, dp[i-1][j] + 1 if word[i] == word2[j]
// dp[i][j] = min(dp[i-1][j-1] + 1, dp[i][j-1] + 1, dp[i-1][j] + 1 if word[i] != word2[j]
// double sequence dp
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
				dp[i][j] = min(dp[i-1][j-1], dp[i][j-1] + 1, dp[i-1][j] + 1)
			} else{
				dp[i][j] = min(dp[i-1][j-1] + 1, dp[i][j-1] + 1, dp[i-1][j] + 1)
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
