package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dkim/aoc/2021/utils"
)

func main() {
	text := utils.ReadInput(1)
	// Look for depth changes
	part1(text)
	part2(text)

}

func part1(depths []string) {
	defer utils.TimeTrack(time.Now(), "Part 1")
	var old_depth = 0
	var count = 0
	for i := 0; i < len(depths)-2; i++ {
		depth1, _ := strconv.Atoi(depths[i])
		depth2, _ := strconv.Atoi(depths[i+1])
		depth3, _ := strconv.Atoi(depths[i+2])
		current_depth := depth1 + depth2 + depth3
		if old_depth == 0 {
			old_depth = current_depth
		}
		if old_depth < current_depth {
			count++
		}
		old_depth = current_depth
	}
	fmt.Printf("Part 1 depth: %d\n", count)

}

func part2(depths []string) {
	defer utils.TimeTrack(time.Now(), "Part 2")
	// Look for depth changes
	var old_depth = 118
	var count = 0
	for _, each_ln := range depths {
		depth, _ := strconv.Atoi(each_ln)
		if old_depth < depth {
			count++
		}
		old_depth = depth
	}
	fmt.Printf("Part 2 depth: %d\n", count)
}
