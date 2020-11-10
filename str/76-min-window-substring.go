package main

// 76 min window substring
// solution is like a sparse rope, first extend, then shrink O(n)
func minWindow(s string, t string) string {
	if t == "" || len(s) < len(t) {
		return ""
	}
	target := buildDict(t)
	dict := map[byte]int{}
	for i := 0; i < len(t) - 1; i++ {
		if _, exist := target[s[i]]; exist {
			if _, exist := dict[s[i]]; exist {
				dict[s[i]] += 1
			} else {
				dict[s[i]] = 1
			}
		}
	}
	found := false
	result := s
	i := 0
	j := len(t) - 1
	for j < len(s) {
		for j < len(s) && !equal(dict, target) {
			j += 1
			if _, exist := target[s[j-1]]; exist {
				if _, exist2 := dict[s[j-1]]; exist2 {
					dict[s[j-1]] += 1
				} else {
					dict[s[j-1]] = 1
				}
			}
		}
		for j <= len(s) && i < j && equal(dict, target) {
			found = true
			if len(result) > j - i {
				result = s[i:j]
			}
			if _, exist := target[s[i]]; exist {
				dict[s[i]] -= 1
				if dict[s[i]] == 0 {
					delete(dict, s[i])
				}
			}
			i += 1
		}
	}
	if found {
		return result
	} else {
		return ""
	}
}

func equal(a, b map[byte]int) bool {
	if len(a) == len(b) {
		for key, val := range a {
			if b[key] > val {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func buildDict(s string) map[byte]int {
	result := map[byte]int{}
	for _, ele := range s {
		if val, exist := result[byte(ele)]; exist {
			result[byte(ele)] = val + 1
		} else {
			result[byte(ele)] = 1
		}
	}
	return result
}
