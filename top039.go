package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	list := make([]int, 0, 10)
	result := make([][]int, 0, 10)
	recursiveFun(candidates, &list, 0, &result, target, 0)
	return result
}

func recursiveFun(candidates []int, list *[]int, sum int, result *[][]int, target int, index int) {
	if sum == target {
		temp := make([]int, len(*list))
		copy(temp, *list)
		*result = append(*result, temp)
	} else if sum < target {
		for i := index; i < len(candidates); i++ {
			sum += candidates[i]
			*list = append(*list, candidates[i])
			recursiveFun(candidates, list, sum, result, target, i)
			*list = (*list)[:len(*list)-1]
			sum -= candidates[i]
		}
	}
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}
