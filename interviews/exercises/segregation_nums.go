package exercises

// Segregate0and1 given a list of 0s and 1s. We need to sort it so that all the 0s are before all the 1s
func Segregate0and1(nums []int) []int {
	start, end := 0, len(nums)-1
	for start < end {
		if nums[start] == 1 {
			nums[start], nums[end] = nums[end], nums[start]
			end--
		} else {
			start++
		}
	}

	return nums
}

// Segregate0and1and2 Given an array of 0s, 1s and 2s
// We need to sort it so that all the 0s are before all the 1s and all the 1s are before 2s
func Segregate0and1and2(nums []int) []int {
	low, mid, high := 0, 0, len(nums)-1

	for mid <= high {
		if nums[mid] == 0 {
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		} else if nums[mid] == 1 {
			mid++
		} else {
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}

	return nums
}
