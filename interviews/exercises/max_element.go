package exercises

import "math"

// MaxElement Find the largest element in the list
func MaxElement(nums []int) int {
	max := math.MinInt
	for i := 0; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}

	return max
}
