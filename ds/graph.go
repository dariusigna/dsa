package ds

type Graph struct {
	adjList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		adjList: make(map[int][]int),
	}
}

func NewNGraph(n int) *Graph {
	g := NewGraph()
	for i := 0; i < n; i++ {
		g.AddVertex(i)
	}
	return g
}

func (g *Graph) AddVertex(v int) {
	if _, ok := g.adjList[v]; !ok {
		g.adjList[v] = []int{}
	}
}

func (g *Graph) AddEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
	g.adjList[v] = append(g.adjList[v], u)
}

func (g *Graph) Vertices() []int {
	vertices := make([]int, 0, len(g.adjList))
	for v := range g.adjList {
		vertices = append(vertices, v)
	}
	return vertices
}

func (g *Graph) RemoveVertex(v int) {
	for _, u := range g.adjList[v] {
		g.RemoveEdge(u, v)
	}
	delete(g.adjList, v)
}

func (g *Graph) RemoveEdge(u, v int) {
	for i, x := range g.adjList[u] {
		if x == v {
			g.adjList[u] = append(g.adjList[u][:i], g.adjList[u][i+1:]...)
			break
		}
	}

	for i, x := range g.adjList[v] {
		if x == u {
			g.adjList[v] = append(g.adjList[v][:i], g.adjList[v][i+1:]...)
			break
		}
	}
}

func (g *Graph) Neighbors(v int) []int {
	return g.adjList[v]
}
