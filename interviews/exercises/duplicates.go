package exercises

// FindDuplicates find the duplicate elements in a list of size n where each element is in the range 0 to n-1
func FindDuplicates(nums []int) []int {
	var result []int
	length := len(nums)

	for i := 0; i < length; i++ {
		x := nums[i] % length
		nums[x] = nums[x] + length
	}

	for i := 0; i < length; i++ {
		if nums[i] >= length*2 {
			result = append(result, i)
		}
	}

	return result
}

// RemoveDuplicates find the duplicate elements and remove it from list
func RemoveDuplicates(nums []int) []int {
	encountered := map[int]bool{}
	var result []int

	for k := range nums {
		if !encountered[nums[k]] {
			encountered[nums[k]] = true
			result = append(result, nums[k])
		}
	}

	return result
}
