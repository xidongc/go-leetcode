package main

/*
44	wildcard matching
	Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*' where:
	'?' Matches any single character.
	'*' Matches any sequence of characters (including the empty sequence).
	The matching should cover the entire input string (not partial).

	double sequence dp dp[i][j] = OR
		(str1[i] == str2[j] and is not alpha) && (dp[i-1][j-1])
		(str2[j] == "?") && (dp[i-1][j-1])
		(str2[j] == "*") && (dp[i-1][j-1] && dp[i][j-1] && dp[i-1][j])

Example 1:
Input: s = "aa", p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".

Example 2:
Input: s = "aa", p = "*"
Output: true
Explanation: '*' matches any sequence.

Example 3:
Input: s = "cb", p = "?a"
Output: false
Explanation: '?' matches 'c', but the second letter is 'a', which does not match 'b'.

Example 4:
Input: s = "adceb", p = "*a*b"
Output: true
Explanation: The first '*' matches the empty sequence, while the second '*' matches the substring "dce".

Example 5:
Input: s = "acdcb", p = "a*c?b"
Output: false
*/
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
