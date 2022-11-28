package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dkim/aoc/2021/utils"
)

type foldPair struct {
	direction string
	axis      int
}

func part1(text []string) {
	defer utils.TimeTrack(time.Now(), "Part 1")
	maxX := 0
	maxY := 0
	for _, line := range text {
		if !strings.Contains(line, "fold") && len(line) != 0 {
			x, _ := strconv.Atoi(string(strings.Split(line, ",")[0]))
			y, _ := strconv.Atoi(string(strings.Split(line, ",")[1]))
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
		}
	}
	var grid [][]string
	var row []string
	for j := 0; j < maxY+1; j++ {
		for i := 0; i < maxX+1; i++ {
			row = append(row, ".")
		}
		grid = append(grid, row)
		row = nil
	}

	for _, line := range text {
		if !strings.Contains(line, "fold") && len(line) != 0 {
			strings.Split(line, ",")
			x, _ := strconv.Atoi(string(strings.Split(line, ",")[0]))
			y, _ := strconv.Atoi(string(strings.Split(line, ",")[1]))
			grid[y][x] = string('#')
		}
	}
	//printGrid(grid)

	// Do the folding now
	var direction string
	var axis int
	var fp []foldPair
	for _, line := range text {
		if strings.Contains(line, "fold") {
			fold := strings.Fields(line)
			values := strings.Split(fold[len(fold)-1], "=")
			direction = values[0]
			axis, _ = strconv.Atoi(values[1])
			fp = append(fp, foldPair{direction, axis})
		}
	}
	for _, i := range fp {
		grid = fold(grid, i.direction, i.axis)
		fmt.Printf("%d dots\n", countDots(grid))

	}
	printGrid(grid)
}

func fold(grid [][]string, direction string, axis int) [][]string {
	var newGrid [][]string
	if direction == "y" {
		// vertical fold
		// Make new grid
		i := 0
		for y := 0; y < len(grid); y++ {
			var line []string
			for x := 0; x < len(grid[0]); x++ {
				if y < axis {
					line = append(line, grid[y][x])
				} else if y > axis {
					if grid[y][x] == string('#') {
						newGrid[i][x] = grid[y][x]
					}
				}
			}
			if y < axis {
				i++
				newGrid = append(newGrid, line)
			} else {
				i--
			}
		}
	}
	if direction == "x" {
		// horizontal fold
		// Make new grid
		i := 0
		for y := 0; y < len(grid); y++ {
			var line []string
			for x := 0; x < axis; x++ {
				line = append(line, ".")
			}
			newGrid = append(newGrid, line)
			line = nil
		}

		for x := 0; x < len(grid[0]); x++ {
			for y := 0; y < len(grid); y++ {
				if x < axis {
					newGrid[y][x] = grid[y][x]
				} else if x > axis {
					if grid[y][x] == string('#') {
						newGrid[y][i] = grid[y][x]
					}
				}
			}
			if x < axis {
				i++
			} else {
				i--
			}
		}
	}
	return newGrid

}

func countDots(grid [][]string) int {
	i := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == string("#") {
				i++
			}
		}
	}
	return i
}

func printGrid(grid [][]string) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			fmt.Printf("%s", grid[y][x])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func main() {
	text := utils.ReadInput(1)
	part1(text)
}
