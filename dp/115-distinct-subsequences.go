package main

// 115 distinct subsequences
// dp[i][j] = dp[i-1][j-1] + dp[i-1][j] if s[i] == t[j]
// dp[i][j] = dp[i-1][j] if s[i] != t[j]
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
