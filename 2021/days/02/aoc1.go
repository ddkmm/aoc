package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Hello world\n")
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()

	// Look for depth changes
	var depth = 0
	var hor = 0
	var aim = 0
	for _, each_ln := range text {
		line := strings.Fields(each_ln)
		test := line[0]
		val, _ := strconv.Atoi(line[1])
		fmt.Printf("%s\n", each_ln)
		switch test {
		case "forward":
			hor += val
			depth += aim * val
		case "down":
			aim += val
		case "up":
			aim -= val
		}
		fmt.Printf("depth: %d\n", depth)
		fmt.Printf("horizontal: %d\n", hor)
		fmt.Printf("answer: %d\n", hor*depth)

	}

	fmt.Printf("depth: %d\n", depth)
	fmt.Printf("horizontal: %d\n", hor)
	fmt.Printf("answer: %d\n", hor*depth)

}
