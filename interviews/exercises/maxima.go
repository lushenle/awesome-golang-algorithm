package exercises

// Maxima Print all the maxima’s in a list
// A value is a maximum if the value before and
// after its index are smaller than it is or does not exist
func Maxima(nums []int) []int {
	var ans []int

	length := len(nums)
	if nums[0] > nums[1] {
		ans = append(ans, nums[0])
	}

	for i := 2; i < length-1; i++ {
		if nums[i-1] < nums[i] && nums[i] > nums[i+1] {
			ans = append(ans, nums[i])
			i++
		}
	}

	if nums[length-1] > nums[length-2] {
		ans = append(ans, nums[length-1])
	}

	return ans
}
