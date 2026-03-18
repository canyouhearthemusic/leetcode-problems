package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	out := make([]int, totalLen)

	i, j, k := 0, 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			out[k] = nums1[i]
			i++
		} else {
			out[k] = nums2[j]
			j++
		}

		k++
	}

	for ; i < len(nums1); i++ {
		out[k] = nums1[i]
		k++
	}

	for ; j < len(nums2); j++ {
		out[k] = nums2[j]
		k++
	}

	if totalLen%2 == 0 {
		return (float64(out[(totalLen-1)/2]) + float64(out[(totalLen-1)/2+1])) / 2.0
	}

	return float64(out[(totalLen-1)/2])
}

func main() {
	nums1 := []int{1, 3, 4}
	nums2 := []int{2}

	result := findMedianSortedArrays(nums1, nums2)

	fmt.Println(result)
}
