package algo

import "github.com/dariusigna/dsa/ds"

func ConnectedComponents(g ds.Graph) int {
	visited := make(map[int]bool)
	cc := 0

	for v := range g.Vertices() {
		if !visited[v] {
			cc++
			DFS(g, v, visited)
		}
	}

	return cc
}

func DFS(g ds.Graph, v int, visited map[int]bool) {
	visited[v] = true

	for _, u := range g.Neighbors(v) {
		if !visited[u] {
			DFS(g, u, visited)
		}
	}
}

func IterativeDFS(g ds.Graph, v int) {
	visited := make(map[int]bool)
	stack := ds.NewStack[int]()
	stack.Push(v)

	for !stack.IsEmpty() {
		v, _ = stack.Pop()
		if !visited[v] {
			visited[v] = true
			for _, u := range g.Neighbors(v) {
				stack.Push(u)
			}
		}
	}
}
