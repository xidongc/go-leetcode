package main

// two sequence dp
// dp[i][j] = max(dp[i-1][j-1] + 1, dp[i][j-1], dp[i-1][j]) if text1[i] == text2[j]
// dp[i][j] = max(dp[i][j-1], dp[i-1][j]) if text1[i] != text2[j]
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, 0)
	for i := 0; i <= len(text1); i ++ {
		tmp := make([]int, len(text2)+1)
		dp = append(dp, tmp)
	}

	for i := 1; i <= len(text1); i ++ {
		for j := 1; j <= len(text2); j ++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = max(dp[i-1][j-1]+1, dp[i-1][j], dp[i][j-1])
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

// 1092 Shortest Common Supersequence
func shortestCommonSupersequence(str1 string, str2 string) string {
	dp := make([][]string, 0)
	for i := 0; i <= len(str1); i ++ {
		tmp := make([]string, len(str2) + 1, len(str2) + 1)
		dp = append(dp, tmp)
	}
	for i := 0; i <= len(str1); i ++ {
		dp[i][0] = str1[0: i]
	}
	for j := 0; j <= len(str2); j ++ {
		dp[0][j] = str2[0:j]
	}

	for i := 1; i <= len(str1); i ++ {
		for j := 1; j <= len(str2); j ++ {
			if str1[i-1] == str2[j-1] {
				if len(dp[i-1][j]) < len(dp[i][j-1]) {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			} else {
				if len(dp[i-1][j]) < len(dp[i][j-1]) {
					dp[i][j] = dp[i-1][j] + string(str1[i-1])
				} else {
					dp[i][j] = dp[i][j-1] + string(str2[j-1])
				}
			}
		}
	}

	return dp[len(str1)][len(str2)]
}
