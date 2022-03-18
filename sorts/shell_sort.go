package sorts

func step(nums []int, start, gap int) {
	length := len(nums)
	for i := start + gap; i < length; i += gap {
		backup := nums[i]
		j := i - gap
		for j >= 0 && backup < nums[j] {
			nums[j+gap] = nums[j]
			j -= gap
		}
		nums[j+gap] = backup
	}
}

func ShellSort(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	} else {
		gap := length / 2
		for gap > 0 {
			for i := 0; i < gap; i++ {
				step(nums, i, gap)
			}
			gap /= 2
		}
	}
	return nums
}
