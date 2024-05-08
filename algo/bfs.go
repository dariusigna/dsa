package algo

import "github.com/dariusigna/dsa/ds"

func BFS(g ds.Graph, v int) {
	visited := make(map[int]bool)
	queue := ds.NewQueue[int]()
	queue.Add(v)

	for !queue.IsEmpty() {
		v, _ = queue.Pop()
		for _, u := range g.Neighbors(v) {
			if !visited[u] {
				visited[u] = true
				queue.Add(u)
			}
		}
	}
}

func BFSWithLevel(g ds.Graph, v int) map[int]int {
	visited := make(map[int]bool)
	queue := ds.NewQueue[int]()
	queue.Add(v)
	level := make(map[int]int)
	level[v] = 0

	for !queue.IsEmpty() {
		v, _ = queue.Pop()
		for _, u := range g.Neighbors(v) {
			if !visited[u] {
				visited[u] = true
				queue.Add(u)
				level[u] = level[v] + 1
			}
		}
	}

	return level
}
