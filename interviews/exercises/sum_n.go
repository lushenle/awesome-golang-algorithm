package exercises

// SumN Write a method to compute Sum(N)=1+2+3+...+N
func SumN(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum
}

// SumN1 Gauss sum
func SumN1(n int) int {
	return n * (n + 1) / 2
}
