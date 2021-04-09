package main

func main() {

}

// Input: s = "3[a]2[bc]"
// Output: "aaabcbc"
func decodeString(s string) string {
	stack := make([]rune, 0)
	multipler := 1
	for _, ele := range s {
		if ele == '[' {
			stack = append(stack, ele)
		} else if ele == ']' {
			tmp := ""
			for len(stack) > 0 {
				ele := stack[len(stack)-1]
				stack = stack[0: len(stack)-1]
				if ele == '[' {
					break
				} else {
					tmp = string(ele) + tmp
				}
			}
			tmp = multipler *
		}
	}

}
