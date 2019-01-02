package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputPath = "./input.txt"

// Vertex contains coordinates from the input
type Vertex struct {
	X int
	Y int
}

func generateNames(n int) string {
	p := make([]byte, n)
	for i := range p {
		p[i] = 'a' + byte(i)
	}
	return string(p)
}

func main() {
	var input []Vertex
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var line string
	x, y := 0, 0
	xMax, yMax := 0, 0
	var grid [][]int
	for scanner.Scan() {
		line = scanner.Text()
		fmt.Sscanf(line, "%d, %d", &x, &y, err)
		if x > xMax {
			xMax = x
		}
		if y > yMax {
			yMax = y
		}
		input = append(input, Vertex{x, y})
	}
	var rows = make([]int, xMax)
	for k := range rows {
		var col = make([int],yMax)
		rows
	}
	fmt.Println(generateNames(len(input)))
	fmt.Printf("%d coordinates\n", len(input))
	fmt.Printf("Max x value is %d, Max y value is %d\n", xMax, yMax)
	fmt.Println(input)
}
