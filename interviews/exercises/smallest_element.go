package exercises

import "math"

// SmallestElement Find the largest element in the list
func SmallestElement(nums []int) int {
	min := math.MaxInt
	for i := 0; i < len(nums); i++ {
		if min > nums[i] {
			min = nums[i]
		}
	}

	return min
}
