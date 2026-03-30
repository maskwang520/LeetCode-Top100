package main

import "fmt"

func longestPalindrome(s string) string {
	maxLen := 1
	result := s[0:1]
	for i := 0; i < len(s); i++ {
		k, j := i-1, i+1
		for j < len(s) && s[j:j+1] == s[i:i+1] {
			j++
		}
		for k >= 0 && j < len(s) {
			if s[k:k+1] == s[j:j+1] {
				k--
				j++
			} else {
				break
			}
		}
		if maxLen < j-k-1 {
			result = s[k+1 : j]
			maxLen = len(result)
		}

	}
	return result
}

func main() {
	fmt.Println(longestPalindrome("b"))
}
