package gcd

// Iterative Faster iterative version of Gcd Recursive
// without holding up too much of the stack
func Iterative(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
