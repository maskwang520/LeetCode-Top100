package main

func subsets(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return [][]int{}
	}
	result := make([][]int, 0)
	result = append(result, []int{})
	list := make([]int, 0)
	deepSearch(nums, &result, list, 0)
	return result
}

func deepSearch(nums []int, result *[][]int, list []int, index int) {
	for i := index; i < len(nums); i++ {
		list = append(list, nums[i])
		copyList := make([]int, len(list))
		copy(copyList, list)
		*result = append(*result, copyList)
		deepSearch(nums, result, list, i+1)
		list = list[:len(list)-1]
	}
}
