package main

import (
	"strings"
)

// 140 word break II TLE
func wordBreak(s string, wordDict []string) []string {
	hashSet := make(map[string]struct{}, 0)
	for _, word := range wordDict {
		hashSet[word] = struct{}{}
	}
	result := make([]string, 0)
	wordBreakHelper(hashSet, s, 0, []string{}, &result)
	return result
}

func wordBreakHelper(hashSet map[string]struct{}, s string, start int, tmp []string, result *[]string) {
	if start == len(s) {
		*result = append(*result, strings.Join(tmp, " "))
		return
	}

	for i := start+1; i <= len(s); i ++ {
		if contains(hashSet, s[start: i]) {
			tmp = append(tmp, s[start:i])
			wordBreakHelper(hashSet, s, i, tmp, result)
			tmp = tmp[:len(tmp)-1]
		}
	}
}

func contains(hashSet map[string]struct{}, s string) bool {
	_, exist := hashSet[s]
	return exist
}
