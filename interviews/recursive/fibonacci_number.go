package recursive

// Fibonacci Given N find the Nth number in the Fibonacci series
// 1, 1, 2, 3, 5, ...
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}
