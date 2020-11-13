package main

import (
	"fmt"
	"strconv"
)

var (
	Priority = map[string]int{
		"+": 2,
		"-": 2,
		"*": 1,
		"/": 1,
		"(": 3,
		")": 3,
	}
)

// common implementation of calculator using rpn
func calculate(s string) int {
	input := cleanInput(s)
	rpn := toRpn(input)
	fmt.Println(rpn)
	fmt.Println(cal(rpn))
	return cal(rpn)
}

// should exist only num, operator
func cleanInput(s string) []string {
	num := 0
	isSet := false
	result := make([]string, 0)
	for i := 0; i < len(s); i ++ {
		if _, exist := Priority[string(s[i])]; exist {
			if isSet {
				value := strconv.Itoa(num)
				result = append(result, value)
			}
			result = append(result, string(s[i]))
			num = 0
			isSet = false
		} else {
			if value, err := strconv.Atoi(string(s[i])); err == nil {
				isSet = true
				num = num * 10 + value
			}
			if i == len(s) - 1 && isSet {
				result = append(result, strconv.Itoa(num))
			}
		}
	}
	return result
}

// regular to reverse rpn expression
func toRpn(input []string) []string {
	stack := make([]string, 0)
	result := make([]string, 0)
	for _, ele := range input {
		if val, exist := Priority[ele]; exist {
			if ele == ")" {
				for len(stack) > 0 && stack[len(stack) - 1] != "(" {
					result = append(result, stack[len(stack) - 1])
					stack = stack[0: len(stack) - 1]
				}
				if len(stack) == 0 {
					panic("please input ( before )")
				}
				stack = stack[0: len(stack) - 1]
			} else if ele == "(" {
				stack = append(stack, ele)
			} else if len(stack) == 0 || val < Priority[stack[len(stack) - 1]] {
				stack = append(stack, ele)
			} else {
				for len(stack) > 0 && val >= Priority[stack[len(stack) - 1]]{
					result = append(result, stack[len(stack) - 1])
					stack = stack[0: len(stack) - 1]
				}
				stack = append(stack, ele)
			}
		} else {
			result = append(result, ele)
		}
	}
	for len(stack) > 0 {
		result = append(result, stack[len(stack) - 1])
		stack = stack[0: len(stack) - 1]
	}
	return result
}

func cal(input []string) int {
	stack := make([]int, 0)
	for _, ele := range input {
		if _, exist := Priority[ele]; exist {
			if len(stack) < 2 {
				panic("expression illegal")
			}
			value := basicCal(stack[len(stack) - 1], stack[len(stack) - 2], ele)
			stack = stack[:len(stack) - 2]
			stack = append(stack, value)
		} else {
			value, err := strconv.Atoi(ele)
			if err != nil {
				panic("can't convert")
			}
			stack = append(stack, value)
		}
	}
	if len(stack) != 1 {
		panic("please check input")
	}
	return stack[0]
}

// basic cal
func basicCal(val1, val2 int, operator string) int {
	if operator == "+" {
		return val1 + val2
	} else if operator == "*" {
		return val1 * val2
	} else if operator == "/" {
		if val1 == 0 {
			panic("cannot divide zero")
		}
		return val2 / val1
	} else if operator == "-" {
		return val2 - val1
	} else {
		panic("operator currently not supported")
	}
}
