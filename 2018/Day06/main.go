package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var inputPath = "./input.txt"

type vertex struct {
	X int
	Y int
}

func distance(x1 int, x2 int, y1 int, y2 int) int {
	x := int(math.Abs(float64(x2 - x1)))
	y := int(math.Abs(float64(y2 - y1)))
	return x + y
}

func main() {
	// initialisations
	//
	var input []vertex
	nullVertex := vertex{0, 0}
	x, y := 0, 0
	const gridDimension int = 350
	var grid [gridDimension][gridDimension]vertex
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = nullVertex
		}
	}
	tl := vertex{0, 0}
	//tr := vertex{0, 0}
	//bl := vertex{0, 0}
	br := vertex{0, 0}

	// Read in input coordinates and identify boundary line
	//
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var line string

	minX, minY := gridDimension-1, gridDimension-1
	maxX, maxY := 0, 0
	for scanner.Scan() {
		line = scanner.Text()
		fmt.Sscanf(line, "%d, %d", &x, &y, err)
		newVertex := vertex{x, y}
		// find min/max X and Y
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}

		input = append(input, newVertex)
		grid[x][y] = newVertex
	}

	tl.X = minX
	for _, v := range input {
		if v.X == tl.X && v.Y <= tl.Y {
			tl = v
		}
	}

	fmt.Printf("%d coordinates\n", len(input))
	fmt.Println("Top left ", tl)
	fmt.Println("Bottom right ", br)

	// use corners to identify boundary
	//
	var boundary []vertex
	for _, v := range input {
		if v.X == maxX || v.X == minX {
			boundary = append(boundary, v)
			continue
		}
		if v.Y == maxY || v.Y == minY {
			boundary = append(boundary, v)
			continue
		}
	}

	fmt.Println("Inputs: ", input)
	fmt.Println("Boundary: ", boundary)

	// Add nullVertex to list to be filtered
	boundary = append(boundary, nullVertex)

	for X := range grid {
		for Y := range grid[X] {
			canWrite := true
			// search through the inputs and see if this is a coordinate
			// that should be skipped
			for _, v := range input {
				if v.X == X && v.Y == Y {
					canWrite = false
				}
			}
			if canWrite {
				dCount := 0
				dFinal := gridDimension + 1
				vFinal := nullVertex
				d := 0
				for _, v := range input {
					d = distance(X, v.X, Y, v.Y)
					if d < dFinal {
						// This is our best candidate
						dFinal = d
						vFinal = v
						// reset the count
						dCount = 0
					} else if d == dFinal {
						// This is a duplicate but might not be
						// the shortest distance
						dCount++
						vFinal = nullVertex
					}
				}
				grid[X][Y] = vFinal
			} else {
				// fmt.Printf("Skipping (%d, %d)\n", X, Y)
			}
		}
	}

	// shrink the board down to the topleft/bottom right
	//
	domains := make(map[vertex]int)

	for yIndex := minY; yIndex <= maxY; yIndex++ {
		for xIndex := minX; xIndex <= maxX; xIndex++ {
			domains[grid[xIndex][yIndex]]++
			//	fmt.Print(grid[xIndex][yIndex])
		}
		//	fmt.Print("\n")
	}

	// Filter out coordinates not eligible for largest area
	//
	for _, i := range boundary {
		fmt.Println("Remove ", i)
		delete(domains, i)
	}

	fmt.Println("New area calculations: ", domains)
	fmt.Println("Finding largest area")
	maxArea := 0
	for k, v := range domains {
		if v >= maxArea {
			maxArea = v
			fmt.Println(k, v)
		}
	}

	// error check
	//
	for _, v := range input {
		if grid[v.X][v.Y] != v {
			fmt.Println("Bang ", v.Y, ",", v.X, v)
		}
	}
}
