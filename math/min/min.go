package min

// Int returns the minimum of all the integers provided as arguments
func Int(nums ...int) int {
	min := nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
	}

	return min
}
