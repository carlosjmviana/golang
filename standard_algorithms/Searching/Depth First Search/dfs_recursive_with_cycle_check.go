package main

import "fmt"

// Assuming that there are no cycles, a binary flag would be
// sufficient to do the job, ie. once we recurse into a node, we
// can mark it as visited. However, if we want to take care of cycles,
// we need a 3-valued flag: 0 refers to unvisited, 1 means that all
// neighbouring nodes including itself has been visited(fully visited),
// and we mark -1 for nodes that have been visited but not all neighbouring
// nodes have been visited(partially visited). This way, when we arrive at a
// node with a visited value of -1 in the midst of a dfs call stack, we know
// that a cycle exists.
const (
	UNVISITED_NODE         = 0
	FULLY_VISITED_NODE     = 1  //all neighbouring nodes including itself has been visited
	PARTIALLY_VISITED_NODE = -1 // node have been visited but not all neighbouring nodes have been visited.
)

func main() {
	adList := [][]int{
		{1, 2},
		{1, 7},
		{1, 8},
		{2, 3},
		{2, 6},
		{8, 9},
		{8, 12},
		{3, 4},
		{3, 5},
		{9, 10},
		{9, 11},
	}

	if IsNonCyclic(adList) {
		fmt.Println("Graph is NOT cyclic")
	} else {
		fmt.Println("Graph is cyclic")
	}
}

func IsNonCyclic(edges [][]int) bool {
	adList := make(map[int][]int)

	for _, edge := range edges {
		node1, node2 := edge[0], edge[1]

		_, exists := adList[node1]
		if !exists {
			adList[node1] = []int{node2}
		} else {
			adList[node1] = append(adList[node1], node2)
		}
	}

	fmt.Println("AdList structure:", adList)

	visited := make(map[int]int)
	fmt.Println("visited:", visited)

	for node := range adList {
		if visited[node] == 1 {
			continue
		}
		if !dfs(node, adList, &visited) {
			return false
		}
	}
	return true
}

func dfs(node int, adjList map[int][]int, visited *map[int]int) bool {
	neighbours, ok := adjList[node]
	if !ok {
		return true
	}

	if (*visited)[node] == PARTIALLY_VISITED_NODE {
		return false
	}

	if (*visited)[node] == FULLY_VISITED_NODE {
		return true
	}

	(*visited)[node] = PARTIALLY_VISITED_NODE
	fmt.Println("visited:", visited)

	for _, neighbour := range neighbours {
		if !dfs(neighbour, adjList, visited) {
			return false
		}
	}
	(*visited)[node] = FULLY_VISITED_NODE
	fmt.Println("visited:", visited)
	return true
}
