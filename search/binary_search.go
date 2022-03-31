package search

import "errors"

// Binary search for target within a sorted array by repeatedly dividing the array in half
// and comparing the midpoint with the target
// if a target is found, the index of the target is returned, else the function return -1 and error
func Binary(nums []int, target, lowIndex, highIndex int) (int, error) {
	if highIndex < lowIndex || len(nums) == 0 {
		return -1, errors.New("target not found")
	}

	mid := lowIndex + (highIndex-lowIndex)/2
	if nums[mid] > target {
		return Binary(nums, target, lowIndex, mid-1)
	} else if nums[mid] < target {
		return Binary(nums, target, mid+1, highIndex)
	} else {
		return mid, nil
	}
}
