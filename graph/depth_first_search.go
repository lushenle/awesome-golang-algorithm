package graph

// DepthFirstSearch is an algorithm for traversing and searching graph data structures
// reference:
//   https://en.wikipedia.org/wiki/Depth-first_search
//   https://www.geeksforgeeks.org/depth-first-search-or-dfs-for-a-graph/
func DepthFirstSearch(start, end int, nodes []int, edges [][]bool) ([]int, bool) {
	return depthFirstSearchHelper(start, end, nodes, edges, false)
}

func depthFirstSearchHelper(start, end int, nodes []int, edges [][]bool, showRoute bool) ([]int, bool) {
	var (
		route, stack []int
	)
	startIndex := getIndex(nodes, start)
	stack = append(stack, startIndex)

	for len(stack) > 0 {
		now := stack[len(stack)-1]
		route = append(route, nodes[now])
		stack = stack[:len(stack)-1]

		for i := 0; i < len(edges[now]); i++ {
			if edges[now][i] && notExist(stack, i) {
				stack = append(stack, i)
			}
			edges[now][i] = false
			edges[i][now] = false
		}

		if route[len(route)-1] == end {
			return route, true
		}
	}

	if showRoute {
		return route, false
	} else {
		return nil, false
	}
}

func getIndex(nodes []int, target int) int {
	for i := 0; i < len(nodes); i++ {
		if nodes[i] == target {
			return i
		}
	}

	return -1
}

func notExist(slice []int, target int) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			return false
		}
	}

	return true
}
