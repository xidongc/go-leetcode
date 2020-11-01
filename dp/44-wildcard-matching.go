package main

// two sequence dp
// dp[i][j] = dp[i-1][j-1] if s[i-1] == p[i-1]
// if p[i-1] = "?" dp[i][j] = dp[i-1][j-1]
// if p[i-1] = "*" dp[i][j] = OR(dp[i-1][j], dp[i-1][j-1], dp[i][j-1])
func isMatch(s string, p string) bool {
	dp := make([][]bool, 0)
	for i := 0; i <= len(s); i ++ {
		tmp := make([]bool, len(p)+1, len(p)+1)
		dp = append(dp, tmp)
	}
	dp[0][0] = true
	for i := 1; i <= len(s); i ++ {
		dp[i][0] = false
	}

	for i := 1; i <= len(p); i ++ {
		dp[0][i] = string(p[i-1])== "*" && dp[0][i-1]
	}

	for i := 1; i <= len(s); i ++ {
		for j := 1; j <= len(p); j ++ {
			if s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if string(p[j-1]) == "?" {
				dp[i][j] = dp[i-1][j-1]
			} else if string(p[j-1]) == "*" {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-1] || dp[i][j-1]
			}
		}
	}
	return dp[len(s)][len(p)]
}
