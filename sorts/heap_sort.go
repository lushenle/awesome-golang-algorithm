package sorts

func heapSortMax(nums []int, length int) []int {
	if length < 2 {
		return nums
	}

	depth := length/2 - 1
	for i := depth; i >= 0; i-- {
		// assuming that the most value is at the location of i
		topMax := i
		leftChild := 2*i + 1
		rightChild := 2*i + 2

		if leftChild <= length-1 && nums[leftChild] > nums[topMax] {
			topMax = leftChild
		}

		if rightChild <= length-1 && nums[rightChild] > nums[topMax] {
			topMax = rightChild
		}

		// make sure the i value is the maximum
		if topMax != i {
			nums[i], nums[topMax] = nums[topMax], nums[i]
		}
	}

	return nums
}

func HeapSort(nums []int) []int {
	length := len(nums)

	for i := 0; i < length; i++ {
		lastMessLen := length - i // intercept a section at a time
		heapSortMax(nums, lastMessLen)
		if i < length {
			nums[0], nums[lastMessLen-1] = nums[lastMessLen-1], nums[0]
		}
	}

	return nums
}
