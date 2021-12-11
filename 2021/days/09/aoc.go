package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/dkim/aoc/2021/utils"
)

type coord struct {
	y int
	x int
}

type basin struct {
	start coord
	body  []coord
}

func part1(grid [][]int) []basin {
	defer utils.TimeTrack(time.Now(), "Part 1")
	sum := 0
	var basins []basin

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			//fmt.Printf("Inspecting %d\n", grid[y][x])
			isLowest := true
			// Left side check
			if x > 0 {
				if grid[y][x-1] <= grid[y][x] {
					isLowest = false
				}
			}
			// right side check
			if x < len(grid[0])-1 {
				if grid[y][x+1] <= grid[y][x] {
					isLowest = false
				}
			}
			// top check
			if y > 0 {
				if grid[y-1][x] <= grid[y][x] {
					isLowest = false
				}
			}
			// bottom check
			if y < len(grid)-1 {
				if grid[y+1][x] <= grid[y][x] {
					isLowest = false
				}
			}

			if isLowest {
				//				fmt.Printf("(%d, %d) %d\n", y, x, grid[y][x])
				var start coord
				start.x = x
				start.y = y
				var body []coord
				body = append(body, coord{y, x})
				var basin basin
				basin.body = body
				basin.start = start
				basins = append(basins, basin)

				sum += grid[y][x] + 1
			}
		}
	}
	fmt.Printf("%d is the answer for part 1\n", sum)
	return basins
}

func findPoint(body []coord, pt coord) bool {
	for _, i := range body {
		if i == pt {
			return true
		}
	}
	return false
}

func part2(grid [][]int, basins []basin) {
	defer utils.TimeTrack(time.Now(), "Part 2")
	var scores []int
	// From each basin start, travel out until you hit a 9, increasing size as you go
	for _, b := range basins {
		//fmt.Printf("Looking at basin (%d, %d): %d\n", b.start.y, b.start.x, grid[b.start.y][b.start.x])
		size := 1
		for i := 0; i < size; i++ {
			x := b.body[i].x
			y := b.body[i].y
			// up
			if y > 0 && grid[y-1][x] != 9 {
				testPoint := coord{y - 1, x}
				if !findPoint(b.body, testPoint) {
					b.body = append(b.body, testPoint)
					size++
				}
			}
			// right
			if x < len(grid[0])-1 && grid[y][x+1] != 9 {
				testPoint := coord{y, x + 1}
				if !findPoint(b.body, testPoint) {
					b.body = append(b.body, testPoint)
					size++
				}
			}
			// down
			if y < len(grid)-1 && grid[y+1][x] != 9 {
				testPoint := coord{y + 1, x}
				if !findPoint(b.body, testPoint) {
					b.body = append(b.body, testPoint)
					size++
				}
			}
			// left
			if x > 0 && grid[y][x-1] != 9 {
				testPoint := coord{y, x - 1}
				if !findPoint(b.body, testPoint) {
					b.body = append(b.body, testPoint)
					size++
				}
			}
		}

		//		fmt.Printf("Basin (%d, %d) with size %d\n", b.start.y, b.start.x, len(b.body))
		scores = append(scores, len(b.body))
	}
	sort.Ints(scores)
	score := scores[len(scores)-1] * scores[len(scores)-2] * scores[len(scores)-3]
	fmt.Printf("%d is the answer for part 2\n", score)

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
	basins := part1(grid)
	part2(grid, basins)

}
