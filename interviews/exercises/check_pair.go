package exercises

// CheckPairs Given a slice with 'n' elements & a value 'x'
//  find two elements in the list that sums to 'x'
func CheckPairs(nums []int, x int) []int {
	m := make(map[int]int)

	for k, v := range nums {
		if i, ok := m[x-v]; ok {
			return []int{i, k}
		} else {
			m[v] = k
		}
	}

	return nil
}
