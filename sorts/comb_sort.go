package sorts

func getNextGap(gap int) int {
	gap = gap * 10 / 13
	if gap < 1 {
		return 1
	}

	return gap
}

func CombSort(nums []int) []int {
	swapped := true
	length := len(nums)
	gap := length
	for gap != 1 || swapped {
		gap = getNextGap(gap)
		swapped = false
		for i := 0; i < length-gap; i++ {
			if nums[i] > nums[i+gap] {
				nums[i], nums[i+gap] = nums[i+gap], nums[i]
				swapped = true
			}
		}
	}

	return nums
}
