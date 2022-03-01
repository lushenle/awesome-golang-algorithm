package exercises

// Reverse a list in-place
// You cannot use any additional list in other wards Space Complexity should be O(1)
func Reverse(nums []int) []int {
	//for start, end := 0, len(nums)-1; start < end; start, end = start+1, end-1 {
	//	nums[start], nums[end] = nums[end], nums[start]
	//}
	//
	//return nums

	start, end := 0, len(nums)-1

	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}

	return nums
}
