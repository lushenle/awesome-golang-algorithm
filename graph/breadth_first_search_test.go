package graph

import "testing"

func TestBreadthFirstSearch(t *testing.T) {
	var bfsTestData = []struct {
		name      string
		start     int
		end       int
		nodes     int
		edges     [][]int
		expected1 bool
		expected2 int
	}{
		{
			name:  "test 1 connected with distance 2",
			start: 0,
			end:   5,
			nodes: 6,
			edges: [][]int{
				{0, 1, 1, 0, 0, 0},
				{1, 0, 0, 1, 0, 1},
				{1, 0, 0, 1, 0, 0},
				{0, 1, 1, 0, 1, 0},
				{0, 0, 0, 1, 0, 0},
				{0, 1, 0, 0, 0, 0},
			},
			expected1: true,
			expected2: 2,
		},
		{
			name:  "test 2 connected with distance 4",
			start: 0,
			end:   5,
			nodes: 6,
			edges: [][]int{
				{0, 1, 1, 0, 0, 0},
				{1, 0, 0, 1, 0, 0},
				{1, 0, 0, 1, 0, 0},
				{0, 1, 1, 0, 1, 0},
				{0, 0, 0, 1, 0, 1},
				{0, 0, 0, 0, 1, 0},
			},
			expected1: true,
			expected2: 4,
		},
		{
			name:  "test 2 not connected",
			start: 0,
			end:   5,
			nodes: 6,
			edges: [][]int{
				{0, 1, 1, 0, 0, 0},
				{1, 0, 0, 1, 0, 0},
				{1, 0, 0, 1, 0, 0},
				{0, 1, 1, 0, 1, 0},
				{0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0},
			},
			expected1: false,
			expected2: 0,
		},
	}
	for _, test := range bfsTestData {
		t.Run(test.name, func(t *testing.T) {
			r1, r2 := BreadthFirstSearch(test.start, test.end,
				test.nodes, test.edges)
			if r1 != test.expected1 || r2 != test.expected2 {
				t.Logf("FAIL: %s", test.name)
				t.Errorf("Nodes '%v' and Edges '%v' start from '%d' and end in '%d' "+
					"expected '%v' with distance '%d' but result was '%v','%d'",
					test.nodes, test.edges, test.start, test.end, test.expected1, test.expected2, r1, r2)
			}
		})
	}
}
