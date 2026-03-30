package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := make([]int, len(nums1)+len(nums2))
	flag := (len(nums1)+len(nums2))%2 == 0
	i, j := 0, 0
	count := 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			merged[count] = nums1[i]
			i++
		} else {
			merged[count] = nums2[j]
			j++
		}
		count++
	}

	for i < len(nums1) {
		merged[count] = nums1[i]
		count++
		i++
	}

	for j < len(nums2) {
		merged[count] = nums2[j]
		count++
		j++
	}

	if flag {
		return float64(merged[(len(nums1)+len(nums2))/2]+merged[(len(nums1)+len(nums2))/2-1]) / 2
	}

	return float64(merged[(len(nums1)+len(nums2))/2])

}
