package main

import "fmt"

const (
	VISITED_NODE     = true
	NOT_VISITED_NODE = false
)

func main() {
	adList := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
	}

	numberOfNodes := 4
	Run(numberOfNodes, adList)
}

func Run(n int, edges [][]int) {
	// Make adjacent list
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
	fmt.Println("AdList:  ", adList)

	// Initialize visited array
	visited := make([]bool, n+1)
	for i := 0; i <= n; i++ {
		visited[i] = false
	}
	fmt.Println("visisted:", visited)
	fmt.Println("-------------------------------------------")
	for i := 1; i <= n; i++ {
		if visited[i] == NOT_VISITED_NODE {
			dfs(i, adList, &visited)
		}
	}
}

func dfs(node int, adjList map[int][]int, visited *[]bool) {
	// mark current node as visisted
	(*visited)[node] = VISITED_NODE

	// print current node or visited list
	fmt.Println("visisted:", node, *visited)

	// do for every edge (v, u)
	edges, _ := adjList[node]
	for _, edge := range edges {
		if (*visited)[edge] == NOT_VISITED_NODE {
			dfs(edge, adjList, visited)
		}
	}
}
