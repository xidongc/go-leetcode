package main

// 214 shortest palindrome
func shortestPalindrome(s string) string {
	result := ""
	possibleEnd := make([]int, 0, 0)
	for i := 1; i < len(s); i ++ {
		if s[i] == s[0]	{
			possibleEnd = append(possibleEnd, i)
		}
	}

	largestEnd := 0
	for _, j := range possibleEnd {
		start, end :=  0, j
		for start < end {
			if s[start] == s[end] {
				start += 1
				end -= 1
			} else {
				break
			}
		}
		if start >= end {
			if j > largestEnd {
				largestEnd = j
			}
		}
	}

	for i := len(s) - 1; i > largestEnd; i -- {
		result += string(s[i])
	}
	result += s
	return result
}
