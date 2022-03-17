package sorts

import "github.com/lushenle/awesome-golang-algorithm/math/max"

// countSort using counting sort to sort the elements in the basis of significant places
func countSort(nums []int, place int) []int {
	var digits [10]int
	output := make([]int, len(nums))

	// calculate count of elements
	for _, n := range nums {
		digits[(n/place)%10]++
	}

	// calculate cumulative count
	for i := 1; i < 10; i++ {
		digits[i] += digits[i-1]
	}

	// place the elements in sorted order
	for i := len(nums) - 1; i >= 0; i-- {
		output[digits[(nums[i]/place)%10]-1] = nums[i]
		digits[(nums[i]/place)%10]--
	}

	return output
}

func unsignedRadixSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	// get maximum element
	maxElement := max.Int(nums...)

	// apply counting sort to sort elements based on place value
	for place := 1; maxElement/place > 0; place *= 10 {
		nums = countSort(nums, place)
	}

	return nums
}

func RadixSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	var negatives, positives []int

	for _, n := range nums {
		if n < 0 {
			negatives = append(negatives, -n)
		} else {
			positives = append(positives, n)
		}
	}
	negatives = unsignedRadixSort(negatives)

	// reverse the negative array and restore signs
	for i, j := 0, len(negatives)-1; i <= j; i, j = i+1, j-1 {
		negatives[i], negatives[j] = -negatives[j], -negatives[i]
	}

	return append(negatives, unsignedRadixSort(positives)...)
}
