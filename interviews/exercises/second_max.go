package exercises

// SecondMax Find the second-largest number in the list
func SecondMax(nums []int) int {
	var first, second int

	if nums[0] > nums[1] {
		first, second = nums[0], nums[1]
	} else {
		first, second = nums[1], nums[0]
	}

	for i := 2; i < len(nums); i++ {
		if nums[i] > first {
			second = first
			first = nums[i]
		} else if nums[i] > second {
			second = nums[i]
		}
	}

	return second
}
