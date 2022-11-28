package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dkim/aoc/2021/utils"
)

type coord struct {
	x int
	y int
}

// segment struct
type segment struct {
	start coord
	end   coord
}

// Process input into segment list
func createSegments(text []string) ([]segment, int) {
	var segments []segment
	maxVal := 0
	for _, each_ln := range text {
		pairs := strings.Split(each_ln, " -> ")
		start := strings.Split(pairs[0], ",")
		end := strings.Split(pairs[1], ",")
		var seg segment
		var point coord
		point.x, _ = strconv.Atoi(start[0])
		if point.x > maxVal {
			maxVal = point.x
		}
		point.y, _ = strconv.Atoi(start[1])
		if point.y > maxVal {
			maxVal = point.y
		}
		seg.start.x = point.x
		seg.start.y = point.y
		point.x, _ = strconv.Atoi(end[0])
		point.y, _ = strconv.Atoi(end[1])
		seg.end.x = point.x
		seg.end.y = point.y
		segments = append(segments, seg)
	}
	return segments, maxVal
}

// initialise gred
func makeGrid(maxVal int) [][]int {
	slice := make([][]int, maxVal+1)
	for i := range slice {
		slice[i] = make([]int, maxVal+1)
	}
	for x := 0; x < maxVal; x++ {
		for y := 0; y < maxVal; y++ {
			slice[x][y] = 0
		}
	}
	return slice
}

// plot
func plot(segments []segment, grid [][]int) {
	defer utils.TimeTrack(time.Now(), "plot")
	for _, seg := range segments {
		if seg.start.x == seg.end.x {
			// vertical line
			if seg.start.y < seg.end.y {
				for y := seg.start.y; y <= seg.end.y; y++ {
					grid[y][seg.start.x]++
				}
			} else {
				for y := seg.end.y; y <= seg.start.y; y++ {
					grid[y][seg.start.x]++
				}

			}
		} else if seg.start.y == seg.end.y {
			// horizontal line
			if seg.start.x < seg.end.x {
				for x := seg.start.x; x <= seg.end.x; x++ {
					grid[seg.start.y][x]++
				}

			} else {
				for x := seg.end.x; x <= seg.start.x; x++ {
					grid[seg.start.y][x]++
				}

			}
		} else {
			// Diagonal line
			if seg.start.x < seg.end.x && seg.start.y < seg.end.y {
				for x := 0; seg.start.x+x <= seg.end.x; x++ {
					grid[seg.start.y+x][seg.start.x+x]++
				}
			} else if seg.start.x > seg.end.x && seg.start.y > seg.end.y {
				for x := 0; seg.end.x+x <= seg.start.x; x++ {
					grid[seg.start.y-x][seg.start.x-x]++
				}
			} else if seg.start.x < seg.end.x && seg.start.y > seg.end.y {
				for x := 0; seg.start.x+x <= seg.end.x; x++ {
					grid[seg.start.y-x][seg.start.x+x]++
				}
			} else {
				for x := 0; seg.start.y+x <= seg.end.y; x++ {
					grid[seg.start.y+x][seg.start.x-x]++
				}
			}

		}
	}
	//printGrid(grid)
	// Look for points with at least 2
	count := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid); y++ {
			if grid[x][y] > 1 {
				count++
			}
		}
	}
	fmt.Printf("%d\n", count)

}

func printGrid(grid [][]int) {
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid); y++ {
			fmt.Printf("%d", grid[x][y])
		}
		fmt.Printf("\n")
	}
}

func main() {

	text := utils.ReadInput(1)
	segments, maxVal := createSegments(text)
	fmt.Printf("We have %d segments with max value %d\n", len(segments), maxVal)
	grid := makeGrid(maxVal)
	plot(segments, grid)
	fmt.Printf("Done")

}
