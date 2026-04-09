package main

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := make([][]int, 0, len(intervals))
	current := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if current[1] >= intervals[i][0] {
			current[1] = max(intervals[i][1], current[1])
		} else {
			result = append(result, current)
			current = intervals[i]
		}
	}
	result = append(result, current)
	return result
}
