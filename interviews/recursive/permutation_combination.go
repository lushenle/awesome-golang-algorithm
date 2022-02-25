package recursive

import "fmt"

// printSlice Format print the slice, length, capacity
func printSlice(data []int) {
	fmt.Printf("%v: len=%d cap=%d\n", data, len(data), cap(data))
}

// swap the positions of two numbers
func swap(data []int, a, b int) {
	data[a], data[b] = data[b], data[a]
}

// Permutation Generate all permutations of an integer list
func Permutation(data []int, n, length int) {
	if length == n {
		printSlice(data)
		return
	}

	for m := n; m < length; m++ {
		swap(data, n, m)
		Permutation(data, n+1, length)
		swap(data, n, m)
	}
}
