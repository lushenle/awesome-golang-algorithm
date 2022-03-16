package sorts

func SelectionSort(nums []int) []int {
	// reduces the effective size of the array by one in  each iteration
	for i := 0; i < len(nums); i++ {
		// temporary variable to store the position of minimum element
		// and assuming the first element to be the minimum of the unsorted array
		min := i

		// gives the effective size of the unsorted  array
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] { //finds the minimum element
				min = j
			}
		}
		// putting minimum element on its proper position
		nums[i], nums[min] = nums[min], nums[i]
	}

	return nums
}
