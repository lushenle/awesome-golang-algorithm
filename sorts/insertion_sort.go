package sorts

import "github.com/lushenle/awesome-golang-algorithm/constraints"

func InsertionSort[T constraints.Ordered](nums []T) []T {
	for currentIndex := 1; currentIndex < len(nums); currentIndex++ {
		// storing current element whose left side is checked for
		// its correct position
		temp := nums[currentIndex]
		iterator := currentIndex

		// check whether the adjacent element in left side is greater or
		// less than the current element
		for ; iterator > 0 && nums[iterator-1] >= temp; iterator-- {

			// moving the left side element to one position forward
			nums[iterator] = nums[iterator-1]
		}

		// moving current element to its  correct position
		nums[iterator] = temp
	}

	return nums
}
