package main

import "fmt"

func isValid(s string) bool {
	var stack []int32
	leftMap := map[int32]struct{}{'{': {}, '(': {}, '[': {}}
	rightMap := map[int32]int32{'}': '{', ')': '(', ']': '['}
	for _, c := range s {
		if _, ok := leftMap[c]; ok {
			stack = append(stack, c)
		} else {
			if val, ok := rightMap[c]; ok && len(stack) > 0 && stack[len(stack)-1] == val {
				stack = stack[0 : len(stack)-1]
			} else {
				return false
			}
		}

	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))
}
