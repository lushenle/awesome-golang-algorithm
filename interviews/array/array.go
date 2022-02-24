package array

//SumArray Return the sum of all the elements of the integer list
func SumArray(data []int) int {
	sum := 0
	for i := 0; i < len(data); i++ {
		sum += data[i]
	}

	return sum
}

// SequentialSearch search a list for some given value
func SequentialSearch(data []int, value int) bool {
	for i := 0; i < len(data); i++ {
		if value == data[i] {
			return true
		}
	}

	return false
}

// BinarySearch Binary search in a sorted list
func BinarySearch(data []int, value int) bool {
	size := len(data)
	mid, low := 0, 0
	high := size - 1

	for low <= high {
		mid = low + (high-low)/2
		if data[mid] == value {
			return true
		}

		if data[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}

	return false
}

// ReverseArray Reverse the list from the specified starting and ending position
func ReverseArray(data []int, start, end int) {
	for start < end {
		data[start], data[end] = data[end], data[start]
		start++
		end--
	}
}

// RotateArray Rotate the list elements K number of times
func RotateArray(data []int, k int) {
	n := len(data)
	ReverseArray(data, 0, k-1)
	ReverseArray(data, k, n-1)
	ReverseArray(data, 0, n-1)
}

// MaxSubArraySum find the contiguous subarray (containing at least one number) which has the largest sum
// and return its sum
func MaxSubArraySum(data []int) int {
	size := len(data)
	maxSoFar, maxEndingHere := 0, 0

	for i := 0; i < size; i++ {
		maxEndingHere = maxEndingHere + data[i]
		if maxEndingHere < 0 {
			maxEndingHere = 0
		}
		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}
	}
	return maxSoFar
}
