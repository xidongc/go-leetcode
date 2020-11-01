package main

// 28 implement str-str
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	result := -1
	for i := 0 ; i <= len(haystack) - len(needle); i ++ {
		if equal(haystack[i: i+len(needle)], needle) {
			result = i
			break
		}
	}
	return result
}

func equal(str1, str2 string) bool {
	return str1 == str2
}
