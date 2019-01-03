package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

var inputPath = "./input.txt"

// Vertex contains coordinates from the input
type Vertex struct {
	X int
	Y int
}

func distance(x1 int, x2 int, y1 int, y2 int) int {
	x := int(math.Abs(float64(x2 - x1)))
	y := int(math.Abs(float64(y2 - y1)))
	return x + y
}

func main() {
	var input []Vertex
	nullVertex := Vertex{0, 0}

	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var line string
	x, y := 0, 0
	gridDimension := 350
	xMax, yMax := 0, 0
	xMin, yMin := gridDimension, gridDimension
	var grid [122500]Vertex
	for i := range grid {
		grid[i] = nullVertex
	}
	for scanner.Scan() {
		line = scanner.Text()
		fmt.Sscanf(line, "%d, %d", &x, &y, err)
		if x > xMax {
			xMax = x
		}
		if x <= xMin {
			xMin = x
		}
		if y > yMax {
			yMax = y
		}
		if y <= yMin {
			yMin = y
		}
		input = append(input, Vertex{x, y})
	}
	for _, v := range input {
		grid[v.Y*gridDimension+v.X] = v
	}

	fmt.Printf("%d coordinates\n", len(input))
	fmt.Printf("Max x value is %d, Max y value is %d\n", xMax, yMax)
	fmt.Printf("Min x value is %d, Min y value is %d\n", xMin, yMin)

	var borders []Vertex
	for _, v := range input {
		if v.Y == yMin || v.Y == yMax || v.X == xMin || v.X == xMax {
			borders = append(borders, v)
		}
	}
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

	domains := make(map[Vertex]int)

	for _, v := range grid {
		domains[v]++
	}

	hack := map[int]Vertex{}
	hackkeys := []int{}
	for k, v := range domains {
		hack[v] = k
		hackkeys = append(hackkeys, v)
	}
	sort.Ints(hackkeys)
	for _, v := range hackkeys {
		fmt.Println(hack[v], v)
	}

	for _, v := range input {
		if grid[v.Y*gridDimension+v.X] != v {
			fmt.Println("Bang ", v)
		}
	}
}
