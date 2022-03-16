package sorts

//func QuickSort(nums []int) []int {
//	if len(nums) < 2 {
//		return nums
//	}
//
//	pivot := nums[0]
//	low, mid, high := make([]int, 0), make([]int, 0), make([]int, 0)
//
//	mid = append(mid, pivot)
//	for i := 0; i < len(nums); i++ {
//		if nums[i] < pivot {
//			low = append(low, nums[i])
//		} else if nums[i] > pivot {
//			high = append(high, nums[i])
//		} else {
//			mid = append(mid, nums[i])
//		}
//	}
//
//	low, high = QuickSort(low), QuickSort(high)
//	return append(append(low, mid...), high...)
//}

//func QuickSort(nums []int) []int {
//	if len(nums) < 2 {
//		return nums
//	}
//
//	pivot := nums[0]
//	var left, right []int
//
//	for _, element := range nums[1:] {
//		if element <= pivot {
//			left = append(left, element)
//		} else {
//			right = append(right, element)
//		}
//	}
//
//	return append(QuickSort(left),
//		append([]int{pivot}, QuickSort(right)...)...)
//}

func partition(nums []int, low, high int) int {
	index := low - 1
	// make the end element as pivot element
	pivotElement := nums[high]
	for i := low; i < high; i++ {
		/* rearrange the array by putting elements which are less than pivot
		on one side and which are greater that on others */
		if nums[i] <= pivotElement {
			index++
			nums[index], nums[i] = nums[i], nums[index]
		}
	}
	// put the pivot element in its proper place
	nums[index+1], nums[high] = nums[high], nums[index+1]
	// return the position of the pivot
	return index + 1
}

// quickSortRange Sorts the specified range within the array
func quickSortRange(nums []int, low, high int) {
	if len(nums) < 2 {
		return
	}

	if low < high {
		pivot := partition(nums, low, high)
		quickSortRange(nums, low, pivot-1)
		quickSortRange(nums, pivot+1, high)
	}
}

func QuickSort(nums []int) []int {
	quickSortRange(nums, 0, len(nums)-1)
	return nums
}
