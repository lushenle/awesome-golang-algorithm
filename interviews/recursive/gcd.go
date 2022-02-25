package recursive

// Gcd Find the greatest common divisor
func Gcd(m, n int) int {
	if m < n {
		return Gcd(n, m)
	}

	if m%n == 0 {
		return n
	}

	return Gcd(n, m%n)
}
