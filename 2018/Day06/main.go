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
	var input []vertex
	nullVertex := vertex{0, 0}

	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var line string
	x, y := 0, 0
	const gridDimension int = 500
	xMax, yMax := 0, 0
	xMin, yMin := 1000, 1000
	var grid [gridDimension * gridDimension]vertex
	for i := range grid {
		grid[i] = nullVertex
	}
	name := 1
	tl := vertex{1000, 1000}
	tr := vertex{0, 1000}
	bl := vertex{1000, 0}
	br := vertex{0, 0}

	for scanner.Scan() {
		line = scanner.Text()
		fmt.Sscanf(line, "%d, %d", &x, &y, err)
		newVertex := vertex{x, y}
		// find top left
		if x <= tl.X && y <= tl.Y {
			tl = newVertex
		}
		// find top right
		if x >= tr.X && y <= tr.Y {
			tr = newVertex
		}
		// find bottom left
		if x <= bl.X && y >= bl.Y {
			bl = newVertex
		}
		// find bottom right
		if x >= br.X && y >= br.Y {
			br = newVertex
		}

		input = append(input, newVertex)
		name++
	}
	for _, v := range input {
		grid[v.Y*gridDimension+v.X] = v
	}

	fmt.Printf("%d coordinates\n", len(input))
	xMax = br.X
	yMax = br.Y
	xMin = tl.X
	yMin = tl.Y
	fmt.Printf("Top left (%d,%d)\n", xMin, yMin)
	fmt.Printf("Bottom right (%d,%d)\n", xMax, yMax)

	var borders []vertex
	borders = append(borders, tl)
	//	borders = append(borders, tr)
	//	borders = append(borders, bl)
	borders = append(borders, br)
	fmt.Println("Inputs: ", input)
	fmt.Println("Borders: ", borders)

	for j := range grid {
		canWrite := true
		testX := j % gridDimension
		testY := int(j / gridDimension)
		for _, v := range input {
			if v.X == testX && v.Y == testY {
				canWrite = false
			}
		}
		if canWrite {
			dCount := 0
			dFinal := gridDimension * 2
			vFinal := nullVertex
			for _, v := range input {
				d := distance(testX, v.X, testY, v.Y)
				if d < dFinal {
					dFinal = d
					dCount = 0
					vFinal = v
				} else if d == dFinal {
					dCount++
					vFinal = nullVertex
				}
			}
			grid[j] = vFinal
		}
	}
	/*
		for k, v := range grid {
			if k%gridDimension == 0 {
				fmt.Print("\n  ", v, "  ")
			} else {
				fmt.Print("  ", v, "  ")
			}
		}
		fmt.Print("\n")
	*/

	// shrink the board down to the topleft/bottom right
	// and add up only those

	domains := make(map[vertex]int)

	// new dimensions of the grid are based on top left and bottom right
	newGridLength := br.X - tl.X
	newGridHeight := br.Y - tl.Y

	// starting index is tl.X + tl.Y*gridDimension
	for yy := 0; yy < newGridHeight; yy++ {
		for xx := 0; xx < newGridLength; xx++ {
			// offset the index
			index := xx + tl.X + (yy+tl.Y)*gridDimension
			domains[grid[index]]++
		}
	}
	fmt.Println(len(grid))
	borders = append(borders, nullVertex)
	for _, i := range borders {
		fmt.Println("deleting ", i)
		delete(domains, i)
	}

	fmt.Println(domains)
	maxArea := 0
	for k, v := range domains {
		if v >= maxArea {
			maxArea = v
			fmt.Println(v, k)
		}
	}

	for _, v := range input {
		if grid[v.Y*gridDimension+v.X] != v {
			fmt.Println("Bang ", v.Y, ",", v.X, v)
		}
	}
}
