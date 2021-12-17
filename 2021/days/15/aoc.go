package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dkim/aoc/2021/utils"
)

type Vertex struct {
	x int
	y int
}

type Graph struct {
	grid [][]int
}

func (graph *Graph) getValue(v Vertex) int {
	return graph.grid[v.y][v.x]
}

func (graph *Graph) getNeighbours(v Vertex) []Vertex {
	var neighbours []Vertex
	// Left
	if v.x > 0 {
		neighbours = append(neighbours, Vertex{v.x - 1, v.y})
	}
	// Up
	if v.y > 0 {
		neighbours = append(neighbours, Vertex{v.x, v.y - 1})
	}
	// Right
	if v.x < len(graph.grid[0])-1 {
		neighbours = append(neighbours, Vertex{v.x + 1, v.y})
	}
	// Down
	if v.y < len(graph.grid)-1 {
		neighbours = append(neighbours, Vertex{v.x, v.y + 1})
	}
	return neighbours
}

func part1(grid [][]int) {
	defer utils.TimeTrack(time.Now(), "Part 1")
	fmt.Printf("Part 1 answer %d\n", dijkstra(grid))
}

func grow(grid [][]int) [][]int {
	var bigGrid [][]int
	for y := 0; y < len(grid)*5; y++ {
		var row []int
		for x := 0; x < len(grid[0])*5; x++ {
			row = append(row, 0)
		}
		bigGrid = append(bigGrid, row)
	}
	// now fill it

	xOffset := 0
	yOffset := 0
	for j := 0; j < 5; j++ {
		for k := 0; k < 5; k++ {
			for y := 0; y < len(grid); y++ {
				for x := 0; x < len(grid[0]); x++ {
					val := grid[y][x] + j + k
					if val > 9 {
						val = val - 9
					}
					bigGrid[y+yOffset][x+xOffset] = val
				}
			}
			xOffset += len(grid[0])
		}
		xOffset = 0
		yOffset += len(grid)
	}

	return bigGrid
}

func part2(grid [][]int) {
	defer utils.TimeTrack(time.Now(), "Part 1")
	bigGrid := grow(grid)
	fmt.Printf("Part 2 answer %d\n", dijkstra(bigGrid))
}

func dijkstra(grid [][]int) int {
	Q := make(map[Vertex]bool)
	dist := make(map[Vertex]int)
	prev := make(map[Vertex]*Vertex)
	g := Graph{grid}
	start := Vertex{0, 0}
	dest := Vertex{len(g.grid[0]) - 1, len(g.grid) - 1}
	score := 0

	// Setup
	for y := 0; y < len(g.grid); y++ {
		for x := 0; x < len(g.grid[0]); x++ {
			Q[Vertex{x, y}] = true
			dist[Vertex{x, y}] = 999
			prev[Vertex{x, y}] = nil
		}
	}
	dist[start] = 0

	for len(Q) != 0 {
		var minVertex Vertex
		min := 999
		for k, _ := range Q {
			if dist[k] < min {
				min = dist[k]
				minVertex = k
			}
		}
		//fmt.Println(minVertex)
		delete(Q, minVertex)
		if minVertex == dest {
			var path []Vertex
			var u *Vertex
			u = &dest
			if (prev[*u] != nil) || (*u == start) {
				for *u != start {
					path = append(path, *u)
					u = prev[*u]
					score += g.getValue(*u)
				}
			}
			/*
				fmt.Println(path)
				for s := len(path) - 1; s >= 0; s-- {
					fmt.Println(g.getValue(path[s]))
				}
			*/
			break

		}
		neighbours := g.getNeighbours(minVertex)
		for _, n := range neighbours {
			if Q[n] {
				alt := dist[minVertex] + g.getValue(n)
				if alt < dist[n] {
					dist[n] = alt
					prev[n] = &minVertex
				}

			}
		}
	}
	return score
}

func main() {
	text := utils.ReadInput(1)
	var numbers []int
	var grid [][]int
	for _, line := range text {
		for _, digit := range line {
			number, _ := strconv.Atoi(string(digit))
			numbers = append(numbers, number)
		}
		grid = append(grid, numbers)
		numbers = nil
	}
	part1(grid)
	part2(grid)
}
