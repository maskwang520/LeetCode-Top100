package main

import "strings"

func decodeString(s string) string {
	var strStack []string
	var numStack []int
	cur := ""
	k := 0

	for _, ch := range s {
		switch {
		case ch >= '0' && ch <= '9':
			k = k*10 + int(ch-'0')
		case ch == '[':
			strStack = append(strStack, cur)
			numStack = append(numStack, k)
			cur = ""
			k = 0
		case ch == ']':
			prev := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			num := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			cur = prev + strings.Repeat(cur, num)
		default:
			cur += string(ch)
		}
	}
	return cur
}
