package main

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	m := map[byte]struct{}{}
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		right = i
		_, ok := m[s[i]]
		if ok {
			for left < right {
				if s[left] != s[right] {
					delete(m, s[left])
					left++
				} else {
					left++
					break
				}
			}
		}
		m[s[i]] = struct{}{}
		maxLen = max(maxLen, len(m))

	}

	return maxLen
}
