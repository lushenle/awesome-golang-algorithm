package exercises

// Average Find average of all the elements in a list
func Average(nums []int) int {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}

	return sum / n
}
