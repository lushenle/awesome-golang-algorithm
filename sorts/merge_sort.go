package sorts

import (
	"github.com/lushenle/awesome-golang-algorithm/constraints"
	"github.com/lushenle/awesome-golang-algorithm/math/min"
)

func MergeSort[T constraints.Ordered](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	// defines the current array in 2 parts
	var mid = len(arr) / 2

	// sort the 1st part of array
	var a = MergeSort(arr[:mid])

	// sort the 2nd part of array
	var b = MergeSort(arr[mid:])

	// merge the both parts by comparing elements of both the parts
	return merge(a, b)
}

func merge[T constraints.Ordered](a, b []T) []T {
	r, i, j := make([]T, len(a)+len(b)), 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}
	}

	for i < len(a) {
		r[i+j] = a[i]
		i++
	}

	for j < len(b) {
		r[i+j] = b[j]
		j++
	}

	return r
}

func MergeIter[T constraints.Ordered](items []T) []T {
	for steps := 1; steps < len(items); steps += steps {
		for i := 0; i+steps < len(items); i += 2 * steps {
			temp := merge(items[i:i+steps], items[i+steps:min.Int(i+2*steps, len(items))])
			copy(items[i:], temp)
		}
	}

	return items
}
