package main

import (
	"fmt"
	"os"
)

func readGraph() map[int][]int {
	var n, m int
	fmt.Fscan(os.Stdin, &n, &m) // n - количество узлов, m - количество ребер

	graph := make(map[int][]int)

	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(os.Stdin, &u, &v)
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u) // Удалить эту строку, если граф ориентированный
	}

	return graph
}

func main() {
	fmt.Println("Введите количество узлов и ребер, затем сами ребра (u, v):")
	graph := readGraph()
	fmt.Println("Построенный граф:", graph)
}
