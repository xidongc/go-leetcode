package main

// 301 remove invalid parentheses: instead of traverse, using divide & conquer
func removeInvalidParentheses(s string) []string {
	if len(s) < 1 {
		return []string{""}
	}
	var result []string
	dict := removeHelper(s)
	length := 0
	for key, val := range dict {
		if val == 0 {
			if len(key) > length {
				result = []string{key}
				length = len(key)
			} else if len(key) == length {
				result = append(result, key)
			}
		}
	}
	if len(result) == 0 {
		return []string{""}
	}
	return result
}

func removeHelper(s string) map[string]int {
	if len(s) == 0 {
		return map[string]int{}
	} else if len(s) == 1 {
		if s[0] == ')' {
			return map[string]int{")": 1}
		} else if s[0] == '(' {
			return map[string]int{}
		} else {
			return map[string]int{string(s[0]):0}
		}
	}
	dict := removeHelper(s[1:])
	newDict := map[string]int{}
	copyMap(dict, newDict)

	if s[0] == '(' {
		for key, val := range dict {
			if val > 0 {
				newDict["(" + key] = val - 1
			}
		}
	} else if s[0] == ')' {
		for key, val := range dict {
			newDict[")" + key] = val + 1
		}
		newDict[")"] = 1
	} else {
		for key, val := range dict {
			newDict[string(s[0]) + key] = val
		}
		newDict[string(s[0])] = 0
	}
	return newDict
}

func copyMap(original, source map[string]int) {
	for key, value := range original {
		source[key] = value
	}
}
