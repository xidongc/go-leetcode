package main

import "fmt"

// dp[i][j] = OR((dp[i-1][j] && s1[i-1] == s3[i+j-1]), (dp[i][j-1] && s2[j-1] == s3[i+j-1]))
// dp[0][j] = s2[j] == s3[j]
// dp[i][0] = s1[i] == s3[i]
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1) + len(s2) {
		return false
	}

	dp := make([][]bool, 0)

	for i := 0; i <= len(s1); i ++ {
		tmp := make([]bool, len(s2)+1, len(s2)+1)
		dp = append(dp, tmp)
	}

	if len(s3) == 0 {
		dp[0][0] = true
	}

	for i := 1; i <= len(s1); i ++ {
		if s1[i-1] == s3[i-1] {
			dp[i][0] = true
		}
	}

	for j := 1; j <= len(s2); j ++ {
		if s2[j-1] == s3[j-1] {
			dp[0][j] = true
		}
	}

	for i := 1; i <= len(s1); i ++ {
		for j := 1; j <= len(s2); j ++ {
			iMatch := dp[i-1][j] && s1[i-1] == s3[i+j-1]
			jMatch := dp[i][j-1] && s2[j-1] == s3[i+j-1]
			dp[i][j] = iMatch || jMatch
		}
	}
	return dp[len(s1)][len(s2)]
}

func main() {
	s1 := "db"
	s2 := "b"
	s3 := "cbb"
	fmt.Println(isInterleave(s1, s2, s3))
}