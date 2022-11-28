package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dkim/aoc/2021/utils"
)

// Increment everything in the grid
// Find any 10s and increment all adjacent squares
// but skip if it's a 10
// keep repeating this until there are no more 10s
// Count up all 10s and set them to 0
// repeat the whole thing

func incrementAdjacent(grid [][]int, x int, y int) {

	// fmt.Printf("Incrementing adjecent to (%d, %d):\n", x, y)
	// up
	if y > 0 {
		if grid[y-1][x] < 10 {
			grid[y-1][x]++
		} else if grid[y-1][x] == 10 {
			grid[y-1][x]++
			incrementAdjacent(grid, x, y-1)
		}
	}
	// up-right
	if y > 0 && x < len(grid[0])-1 {
		if grid[y-1][x+1] < 10 {
			grid[y-1][x+1]++
		} else if grid[y-1][x+1] == 10 {
			grid[y-1][x+1]++
			incrementAdjacent(grid, x+1, y-1)
		}
	}
	// right
	if x < len(grid[0])-1 {
		if grid[y][x+1] < 10 {
			grid[y][x+1]++
		} else if grid[y][x+1] == 10 {
			grid[y][x+1]++
			incrementAdjacent(grid, x+1, y)
		}
	}
	// down-right
	if y < len(grid)-1 && x < len(grid[0])-1 {
		if grid[y+1][x+1] < 10 {
			grid[y+1][x+1]++
		} else if grid[y+1][x+1] == 10 {
			grid[y+1][x+1]++
			incrementAdjacent(grid, x+1, y+1)
		}
	}
	// down
	if y < len(grid)-1 {
		if grid[y+1][x] < 10 {
			grid[y+1][x]++
		} else if grid[y+1][x] == 10 {
			grid[y+1][x]++
			incrementAdjacent(grid, x, y+1)
		}
	}
	// down-left
	if y < len(grid)-1 && x > 0 {
		if grid[y+1][x-1] < 10 {
			grid[y+1][x-1]++
		} else if grid[y+1][x-1] == 10 {
			grid[y+1][x-1]++
			incrementAdjacent(grid, x-1, y+1)
		}
	}
	// left
	if x > 0 {
		if grid[y][x-1] < 10 {
			grid[y][x-1]++
		} else if grid[y][x-1] == 10 {
			grid[y][x-1]++
			incrementAdjacent(grid, x-1, y)
		}
	}
	// up-left
	if y > 0 && x > 0 {
		if grid[y-1][x-1] < 10 {
			grid[y-1][x-1]++
		} else if grid[y-1][x-1] == 10 {
			grid[y-1][x-1]++
			incrementAdjacent(grid, x-1, y-1)
		}
	}
}

func part1(grid [][]int, steps int) {
	defer utils.TimeTrack(time.Now(), "Part 1")
	sum := 0

	for i := 0; i < steps; i++ {
		// Increment everything
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[0]); x++ {
				grid[y][x]++
			}
		}

		// go through pulse cycle
		for j := 0; j < 10; j++ {
			for y := 0; y < len(grid); y++ {
				for x := 0; x < len(grid[0]); x++ {
					if grid[y][x] == 10 {
						grid[y][x]++
						incrementAdjacent(grid, x, y)
					}
				}
			}
		}

		// Count pulses and reset
		localSum := 0
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[0]); x++ {
				if grid[y][x] > 9 {
					grid[y][x] = 0
					localSum++
				}
			}
		}
		fmt.Printf("Round %d: %d\n", i+1, localSum)
		sum += localSum
		//printGrid(grid)
	}

	fmt.Printf("%d is the answer for part 1\n", sum)
}

func part2(grid [][]int) {
	defer utils.TimeTrack(time.Now(), "Part 1")
	sum := 0
	round := 0

	for sum != 100 {
		round++
		// Increment everything
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[0]); x++ {
				grid[y][x]++
			}
		}

		// go through pulse cycle
		for j := 0; j < 10; j++ {
			for y := 0; y < len(grid); y++ {
				for x := 0; x < len(grid[0]); x++ {
					if grid[y][x] == 10 {
						grid[y][x]++
						incrementAdjacent(grid, x, y)
					}
				}
			}
		}

		// Count pulses and reset
		sum = 0
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[0]); x++ {
				if grid[y][x] > 9 {
					grid[y][x] = 0
					sum++
				}
			}
		}
		fmt.Printf("Round %d: %d\n", round, sum)
		printGrid(grid)
	}

	fmt.Printf("%d is the answer for part 2\n", round)
}

func printGrid(grid [][]int) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			fmt.Printf("%3d", grid[y][x])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func main() {
	text := utils.ReadInput(1)
	var numbers []int
	var grid [][]int
	var grid2 [][]int
	for _, line := range text {
		for _, digit := range line {
			number, _ := strconv.Atoi(string(digit))
			numbers = append(numbers, number)
		}
		grid = append(grid, numbers)
		grid2 = append(grid, numbers)
		numbers = nil
	}
	//part1(grid, 100)
	part2(grid2)
}
