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
	start := Vertex{0, 0}
	dest := Vertex{len(grid[0]) - 1, len(grid) - 1}
	fmt.Printf("Part 1a answer %d\n", dijkstra(grid, start, dest))
	fmt.Printf("Part 1b answer %d\n", dijkstra(grid, dest, start))
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
	defer utils.TimeTrack(time.Now(), "Part 2")
	bigGrid := grow(grid)
	start := Vertex{0, 0}
	dest := Vertex{len(bigGrid[0]) - 1, len(bigGrid) - 1}
	// Part 2 answer 3018
	// 2021/12/18 00:41:06 Part 2 took 49m59.607356208s
	// fmt.Printf("Part 2a answer %d\n", dijkstra(bigGrid, start, dest))
	/*
		a is forwards, b is backwards. This was a compiled run in terminal
		   	Part 1a answer 712
		   	Part 1b answer 720
		    2021/12/18 07:06:39 Part 1 took 4.479031s
		    Part 2b answer 3025
		    2021/12/18 07:37:03 Part 2 took 30m24.327768667s
	*/
	fmt.Printf("Part 2b answer %d\n", dijkstra(bigGrid, dest, start))

}

func dijkstra(grid [][]int, start Vertex, dest Vertex) int {
	Q := make(map[Vertex]bool)
	dist := make(map[Vertex]int)
	prev := make(map[Vertex]*Vertex)
	g := Graph{grid}
	score := 0

	// Setup
	for y := 0; y < len(g.grid); y++ {
		for x := 0; x < len(g.grid[0]); x++ {
			Q[Vertex{x, y}] = true
			dist[Vertex{x, y}] = 999999
			//			prev[Vertex{x, y}] = nil
		}
	}
	dist[start] = 0

	for len(Q) != 0 {
		//		fmt.Printf("Length of Q is %d\n", len(Q))
		var minVertex Vertex
		min := 999999
		for k, _ := range Q {
			if dist[k] < min {
				min = dist[k]
				minVertex = k
			}
		}
		//		fmt.Printf("New minVertex is ")
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
			//	fmt.Println(len(path))
			break

		}
		neighbours := g.getNeighbours(minVertex)
		for _, n := range neighbours {
			if Q[n] {
				alt := dist[minVertex] + g.getValue(n)
				//			fmt.Printf("%d ? %d\n", alt, dist[n])
				if alt < dist[n] {
					//					fmt.Printf("%d before, %d after\n", dist[n], alt)

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
