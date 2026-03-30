package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	p1, p2 := 0, 0
	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] < nums2[p2] {
			p1++
		} else {
			p2++
		}
	}
}
