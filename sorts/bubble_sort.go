package sorts

func BubbleSort(nums []int) []int {
	swapped := true

	for swapped {
		swapped = false
		// `length -i - 1` is for ignoring comparisons of elements which have
		// already been compared in earlier iterations
		for i := 0; i < len(nums)-1; i++ {
			if nums[i+1] < nums[i] {
				// here swapping of positions is being done
				nums[i+1], nums[i] = nums[i], nums[i+1]
				swapped = true
			}
		}
	}

	return nums
}
