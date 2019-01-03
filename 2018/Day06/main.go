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
	nullVertex := Vertex{-1, -1}

	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var line string
	x, y := 0, 0
	xMax, yMax := 0, 0
	xMin, yMin := 10, 10
	var grid [10][10]Vertex
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = nullVertex
		}
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
		grid[v.X][v.Y] = v
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

	for i := range grid {
		for j := range grid[i] {
			canWrite := true
			for _, v := range input {
				if v.X == i && v.Y == j {
					canWrite = false
				}
			}
			if canWrite {
				dCount := 0
				dFinal := 700
				vFinal := nullVertex
				for _, v := range input {
					d := distance(i, v.X, j, v.Y)
					if d < dFinal {
						dFinal = d
						dCount = 0
						vFinal = v
					} else if d == dFinal {
						dCount++
						vFinal = nullVertex
					}
				}
				grid[i][j] = vFinal
			}
		}
	}

	for i := range grid {
		for j := range grid[i] {
			fmt.Print("  ", grid[i][j], "  ")
		}
		fmt.Print("\n")
	}

	domains := make(map[Vertex]int)

	for i := range grid {
		for j := range grid[i] {
			domains[grid[i][j]]++
		}
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

}
