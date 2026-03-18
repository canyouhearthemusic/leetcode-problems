package main

import "fmt"

func searchRange(nums []int, target int) []int {
	return []int{findFirst(nums, target), findLast(nums, target)}
}

func findFirst(nums []int, target int) int {
	lo, hi, result := 0, len(nums)-1, -1
	for lo <= hi {
		mid := (lo + hi) / 2
		if nums[mid] == target {
			result = mid
			hi = mid - 1
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return result
}

func findLast(nums []int, target int) int {
	lo, hi, result := 0, len(nums)-1, -1
	for lo <= hi {
		mid := (lo + hi) / 2
		if nums[mid] == target {
			result = mid
			lo = mid + 1
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return result
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}

	result := searchRange(nums, 8)

	fmt.Println(result)
}
