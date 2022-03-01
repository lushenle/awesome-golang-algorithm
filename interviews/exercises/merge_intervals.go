package exercises

import "sort"

// MergeIntervals Given a list of intervals, merge all overlapping intervals
func MergeIntervals(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var ans [][]int
	ans = append(ans, intervals[0])
	for i := 1; i < len(intervals); i++ {
		lastInterval := ans[len(ans)-1]
		currLeft := intervals[i][0]
		currRight := intervals[i][1]

		if lastInterval[1] < currLeft {
			ans = append(ans, intervals[i])
		} else {
			lastInterval[1] = max(lastInterval[1], currRight)
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
