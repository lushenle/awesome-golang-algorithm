package sorts

func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	// defines the current array in 2 parts
	var mid = len(nums) / 2

	// sort the 1st part of array
	var a = MergeSort(nums[:mid])

	// sort the 2nd part of array
	var b = MergeSort(nums[mid:])

	// merge the both parts by comparing elements of both the parts
	return merge(a, b)
}

func merge(a, b []int) []int {
	r, i, j := make([]int, len(a)+len(b)), 0, 0
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
