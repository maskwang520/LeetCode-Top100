package main

func lengthOfLongestSubstring(s string) int {
	left, right, maxLen := 0, 0, 0
	m := map[byte]struct{}{}
	for right < len(s) {
		if _, ok := m[s[right]]; !ok {
			m[s[right]] = struct{}{}
			right++
		} else {
			delete(m, s[left])
			left++
		}
		maxLen = max(maxLen, right-left)
	}

	return maxLen
}
