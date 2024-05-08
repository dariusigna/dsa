package algo

import "github.com/dariusigna/dsa/ds"

func ConnectedComponentsGrid(grid [][]int) int {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	cc := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 && !visited[i][j] {
				cc++
				DFSGrid(grid, i, j, visited)
			}
		}
	}

	return cc
}

var directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func outOfBounds(i, j, m, n int) bool {
	return i < 0 || i >= m || j < 0 || j >= n
}

func DFSGrid(grid [][]int, i, j int, visited [][]bool) {
	visited[i][j] = true
	for _, d := range directions {
		x, y := i+d[0], j+d[1]
		if !outOfBounds(x, y, len(grid), len(grid[0])) && !visited[x][y] {
			DFSGrid(grid, x, y, visited)
		}
	}
}

func GridShortestPath(grid [][]int, start, end []int) int {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	queue := ds.NewQueue[[2]int]()
	queue.Add([2]int{start[0], start[1]})
	visited[start[0]][start[1]] = true
	dist := make([][]int, len(grid))
	for i := range dist {
		dist[i] = make([]int, len(grid[0]))
	}

	for !queue.IsEmpty() {
		current, _ := queue.Pop()
		if current == [2]int{end[0], end[1]} {
			return dist[current[0]][current[1]]
		}

		for _, d := range directions {
			x, y := current[0]+d[0], current[1]+d[1]
			if !outOfBounds(x, y, len(grid), len(grid[0])) && grid[x][y] == 1 && !visited[x][y] {
				visited[x][y] = true
				dist[x][y] = dist[current[0]][current[1]] + 1
				queue.Add([2]int{x, y})
			}
		}
	}

	return -1
}
