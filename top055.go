package main

func canJump(nums []int) bool {
	flagSlice := make([]bool, len(nums))
	flagSlice[0] = true
	for i, flag := range flagSlice {
		if flag {
			for j := 1; j <= nums[i] && i+j < len(nums); j++ {
				flagSlice[i+j] = true
			}
		}
	}

	return flagSlice[len(flagSlice)-1]
}
