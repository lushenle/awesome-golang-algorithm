package sorts

import "github.com/lushenle/awesome-golang-algorithm/constraints"

func BubbleSort[T constraints.Ordered](arr []T) []T {
	swapped := true

	for swapped {
		swapped = false
		// `length -i - 1` is for ignoring comparisons of elements which have
		// already been compared in earlier iterations
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				// here swapping of positions is being done
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
	}

	return arr
}
