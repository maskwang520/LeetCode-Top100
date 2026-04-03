package main

import (
	"fmt"
)

var str = "()"

func generateParenthesis(n int) []string {
	result := make([]string, 0, 10)
	list := make([]int32, 0, 10)
	appendFun(n*2, &list, &result)
	return result
}

func appendFun(n int, list *[]int32, result *[]string) {
	if len(*list) < n {
		for _, c := range str {
			*list = append(*list, c)
			appendFun(n, list, result)
			*list = (*list)[:len(*list)-1]

		}
	} else {
		if validateParenthesis(*list) {
			*result = append(*result, string(*list))
		}
	}
}

func validateParenthesis(list []int32) bool {
	stack := make([]int32, 0, len(list))
	for _, v := range list {
		if v == '(' {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(3))
}
