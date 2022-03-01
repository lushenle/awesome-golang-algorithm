package exercises

// SumOfMatrixElement Find the sum of all the elements of a two-dimensional list
func SumOfMatrixElement(nums [][]int) int {
	length := len(nums)
	var sum int
	for i := 0; i < length; i++ {
		for j := 0; j < len(nums[i]); j++ {
			sum += nums[i][j]
		}
	}

	return sum
}
