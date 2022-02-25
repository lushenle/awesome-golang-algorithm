package recursive

//Factorial calculation N!= N*(N-1)...2*1
func Factorial(n int) int {
	// termination condition
	if n <= 1 {
		return 1
	}

	// body, recursive expansion
	return n * Factorial(n-1)
}
