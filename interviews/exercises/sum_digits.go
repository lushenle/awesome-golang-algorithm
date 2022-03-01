package exercises

// SumOfDigits return the sum of every number in an int number
// Example: input= 1984, output should be 22 (1+9+8+4)
func SumOfDigits(num int) int {
	var ans int
	for num > 0 {
		ans = ans + num%10
		num = num / 10
	}

	return ans
}
