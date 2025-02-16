package internal

import (
	"slices"
	"sync"
)

var directions = []Point{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func ParallelSearch(matrix [][]string, start, end Point) (int, []Point) {
	wg := sync.WaitGroup{}
	muVisited := sync.Mutex{}
	muParents := sync.Mutex{}

	visited := make(map[Point]bool)
	parents := make(map[Point]Point)
	distances := make(map[Point]int)

	wg.Add(1)
	go visitNewDir(start, 0, matrix, end, visited, parents, distances, &wg, &muVisited, &muParents)
	wg.Wait()

	if _, exists := distances[end]; !exists {
		return -1, nil
	}

	path := getPath(parents, end)
	return distances[end], path
}

func visitNewDir(current Point, distance int, matrix [][]string, end Point, visited map[Point]bool, parents map[Point]Point, distances map[Point]int, wg *sync.WaitGroup, muVisited *sync.Mutex, muParents *sync.Mutex) {
	defer wg.Done()

	if current == end {
		return
	}

	muVisited.Lock()
	if visited[current] {
		muVisited.Unlock()
		return
	}
	visited[current] = true
	muVisited.Unlock()

	for _, dir := range directions {
		next := NewPoint(current.X+dir.X, current.Y+dir.Y)

		muVisited.Lock()
		if isValidPoint(next, matrix, visited) {
			muParents.Lock()
			if _, exists := distances[next]; !exists || distances[next] > distance+1 {
				distances[next] = distance + 1
				parents[next] = current
			}
			muParents.Unlock()

			wg.Add(1)
			go visitNewDir(next, distance+1, matrix, end, visited, parents, distances, wg, muVisited, muParents)
		}
		muVisited.Unlock()
	}
}

func isValidPoint(p Point, matrix [][]string, visited map[Point]bool) bool {
	return p.X >= 0 && p.X < len(matrix) && p.Y >= 0 && p.Y < len(matrix[0]) && (matrix[p.X][p.Y] == "." && !visited[p] || matrix[p.X][p.Y] == "E")
}

func getPath(parents map[Point]Point, end Point) []Point {
	path := make([]Point, 0)
	for end != (Point{}) {
		path = append(path, end)
		end = parents[end]
	}

	slices.Reverse(path)
	return path
}
