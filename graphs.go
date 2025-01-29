package main

import (
	"container/list"
	"fmt"
)

func dfs(graph map[int][]int, visited map[int]bool, node int) {
	visited[node] = true
	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			dfs(graph, visited, neighbor)
		}
	}
}

func bfs(graph map[int][]int, start int) []int {
	queue := list.New()
	queue.PushBack(start)
	visited := map[int]bool{start: true}
	result := []int{}

	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(int)
		result = append(result, node)
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue.PushBack(neighbor)
			}
		}
	}
	return result
}

func main() {
	graph := map[int][]int{
		1: {2, 3},
		2: {4, 5},
		3: {6, 7},
		4: {},
		5: {},
		6: {},
		7: {},
	}
	startNode := 1
	result := bfs(graph, startNode)
	fmt.Printf("BFS traversal starting from node %d: %v\n", startNode, result)
}

