package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	var old_depth = 118
	var count = 0
	for _, each_ln := range text {
		depth, _ := strconv.Atoi(each_ln)
		if old_depth < depth {
			count++
		}
		old_depth = depth
	}
	fmt.Printf("depth: %d\n", count)
}
