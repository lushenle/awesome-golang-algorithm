package max

// Int returns the maximum of all the integers provided as arguments
func Int(nums ...int) int {
	ans := nums[0]
	for _, v := range nums {
		if v > ans {
			ans = v
		}
	}
	return ans
}
