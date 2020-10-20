package main

import "fmt"

// dp[i] = dp[j] && s[j:i] in wordDict O(n*2)
// dp[0] = true
func wordBreak(s string, wordDict []string) bool {
	hashSet := make(map[string]struct{}, 0)
	for _, word := range wordDict {
		hashSet[word] = struct{}{}
	}

	dp := make([]bool, len(s)+1, len(s)+1)
	dp[0] = true

	for i := 1; i < len(dp); i ++ {
		for j := 0; j < i; j ++ {
			if dp[j] && contains(hashSet, s[j:i]) {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}

// contains s or not in hash set
func contains(hashSet map[string]struct{}, s string) bool {
	_, exist := hashSet[s]
	return exist
}

func main() {
	input := "leetcode"
	dict := []string{"leet","code"}
	fmt.Println(wordBreak(input, dict))
}