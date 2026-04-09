package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0, len(strs))
	m := make(map[string][]string, len(strs))
	for _, str := range strs {
		orderStr := reOrderStr(str)
		if _, ok := m[orderStr]; !ok {
			m[orderStr] = []string{str}
		} else {
			m[orderStr] = append(m[orderStr], str)
		}
	}

	for _, v := range m {
		result = append(result, v)
	}

	return result
}

func reOrderStr(s string) string {
	chars := []rune(s)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}
