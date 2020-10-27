package main

// 1143 longest common subsequence
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := getLcsDp(text1, text2)
	return dp[len(text1)][len(text2)]
}

// 1092 shortest common super sequence based on 1143
// dp[ptr1][ptr2] == dp[ptr1][ptr2-1] means str2[ptr2-1] does not belong to lcs, add to return seq
// dp[ptr1][ptr2] == dp[ptr1-1][ptr2] means str2[ptr1-1] does not belong to lcs, add to return seq
func shortestCommonSupersequence(str1 string, str2 string) string {
	superSeq := ""
	dp := getLcsDp(str1, str2)
	ptr1, ptr2 := len(str1), len(str2)
	for ptr1 > 0 && ptr2 > 0 {
		if str1[ptr1-1] == str2[ptr2-1] {
			superSeq = string(str1[ptr1-1]) + superSeq
			ptr1 -= 1
			ptr2 -= 1
		} else {
			if dp[ptr1][ptr2] == dp[ptr1][ptr2-1] {
				superSeq = string(str2[ptr2-1]) + superSeq
				ptr2 -= 1
			} else if dp[ptr1][ptr2] == dp[ptr1-1][ptr2] {
				superSeq = string(str1[ptr1-1]) + superSeq
				ptr1 -= 1
			}
		}
	}
	for ptr1 > 0 {
		superSeq = string(str1[ptr1-1]) + superSeq
		ptr1 -= 1
	}
	for ptr2 > 0 {
		superSeq = string(str2[ptr2-1]) + superSeq
		ptr2 -= 1
	}
	return superSeq
}

// get lcs dp using two sequence dp
// dp[i][j] = max(dp[i-1][j-1] + 1, dp[i][j-1], dp[i-1][j]) if str1[i-1] == str2[j-1]
// dp[i][j] = max(dp[i][j-1], dp[i-1][j]) if str1[i-1] != str2[j-1]
// dp[i][0] = 0 dp[0][j] = 0
func getLcsDp(text1, text2 string) [][]int {
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
	return dp
}
