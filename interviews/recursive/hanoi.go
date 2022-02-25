package recursive

import "fmt"

// TOHUtil Use the peg to move the disk
func TOHUtil(num int, from, to, temp string) {
	if num < 1 {
		return
	}

	TOHUtil(num-1, from, temp, to)
	fmt.Println("Move disk", num, "from peg", from, "to peg", to)
	TOHUtil(num-1, temp, to, from)
}

// TowersOfHanoi Call TOHUtil function completion the operation of moving the disk
func TowersOfHanoi(num int) {
	fmt.Println("The sequence of moves involved in the Tower of Hanoi are:")
	TOHUtil(num, "A", "B", "C")
}
