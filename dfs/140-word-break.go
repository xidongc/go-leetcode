package main

// 140 word break II dfs with memorize
func wordBreak(s string, wordDict []string) []string {
	hashSet := make(map[string]struct{}, 0)
	for _, word := range wordDict {
		hashSet[word] = struct{}{}
	}
	return wordBreakHelper(hashSet, s, map[string][]string{})
}

func wordBreakHelper(hashSet map[string]struct{}, s string, memo map[string][]string) []string {
	if val, exist := memo[s]; exist {
		return val
	}

	partition := make([]string, 0)
	if s == "" {
		return partition
	}

	for i := 1; i <= len(s); i ++ {
		if contains(hashSet, s[0: i]) {
			sub := wordBreakHelper(hashSet, s[i:], memo)
			for _, subPartition := range sub {
				partition = append(partition, s[0:i] + " " + subPartition)
			}
		}
	}
	if _, exist := hashSet[s]; exist {
		partition = append(partition, s)
	}
	memo[s] = partition
	return partition
}

func contains(hashSet map[string]struct{}, s string) bool {
	_, exist := hashSet[s]
	return exist
}
