package main

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		if val, ok := numMap[target-num]; ok {
			return []int{i, val}
		}
		numMap[num] = i
	}

	return nil
}
