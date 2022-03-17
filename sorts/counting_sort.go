package sorts

//func CountSort(nums []int) []int {
//	length := len(nums)
//	output := make([]int, length)
//
//	// find the largest element of the array
//	max := nums[0]
//	for i := 1; i < length; i++ {
//		if nums[i] > max {
//			max = nums[i]
//		}
//	}
//	count := make([]int, max+1)
//
//	// store the count of each element
//	// TODO: when the element of the array is negative, it will overflow
//	for i := 0; i < length; i++ {
//		count[nums[i]]++
//	}
//
//	// store the cumulative count of each array
//	for i := 1; i <= max; i++ {
//		count[i] += count[i-1]
//	}
//
//	// find the index of each element of the original array in count array, and
//	// place the elements in output array
//	for i := length - 1; i >= 0; i-- {
//		output[count[nums[i]]-1] = nums[i]
//		count[nums[i]]--
//	}
//
//	return output
//}

func CountSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	min, max := -1000, 1000
	count := make([]int, max-min+1)
	for _, x := range nums {
		count[x-min]++
	}

	z := 0
	for i, c := range count {
		for ; c > 0; c-- {
			nums[z] = i + min
			z++
		}
	}

	return nums
}
