package main

func countSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if isPalindromic(s[i:j]) {
				count++
			}
		}
	}
	return count
}

func isPalindromic(s string) bool {
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
