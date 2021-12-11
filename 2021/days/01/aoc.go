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
	var old_depth = 0
	var count = 0
	for i := 0; i < len(text)-2; i++ {
		depth1, _ := strconv.Atoi(text[i])
		depth2, _ := strconv.Atoi(text[i+1])
		depth3, _ := strconv.Atoi(text[i+2])
		current_depth := depth1 + depth2 + depth3
		if old_depth == 0 {
			old_depth = current_depth
		}
		if old_depth < current_depth {
			count++
		}
		old_depth = current_depth
	}
	fmt.Printf("depth: %d\n", count)
}
