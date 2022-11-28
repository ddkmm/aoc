package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dkim/aoc/2021/utils"
)

const dp = "."
const lp = "#"

type Coord struct {
	x int
	y int
}

type TrenchMap struct {
	algorithm []string
	image     map[Coord]string
	startX    int
	endX      int
	startY    int
	endY      int
}

func (tm *TrenchMap) print() {
	for y := tm.startY; y < tm.endY; y++ {
		for x := tm.startX; x < tm.endY; x++ {
			if tm.image[Coord{x, y}] == "" {
				fmt.Printf(".")
			} else {
				fmt.Printf("%s", tm.image[Coord{x, y}])
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func (tm *TrenchMap) processLocation(pos Coord, step int) string {
	// Read the 3x3 locations around the argument coordinate
	var binString []string
	offset := []int{-1, 0, 1}
	for _, y := range offset {
		for _, x := range offset {
			if tm.image[Coord{pos.x + x, pos.y + y}] == "#" {
				binString = append(binString, "1")
			} else if tm.image[Coord{pos.x + x, pos.y + y}] == "." {
				binString = append(binString, "0")
			} else {
				if step%2 == 0 && tm.algorithm[0] == lp {
					binString = append(binString, "1")
				} else {
					binString = append(binString, "0")
				}
			}
		}
	}

	// Convert to binary and then decimal
	i, _ := strconv.ParseInt(strings.Join(binString, ""), 2, 0)

	// Return the character stored in the algorithm for that decimal value
	return tm.algorithm[i]
}

func (tm *TrenchMap) processImage(step int) {
	newImage := make(map[Coord]string)
	tm.startX--
	tm.startY--
	tm.endX++
	tm.endY++
	for y := tm.startY; y < tm.endY; y++ {
		for x := tm.startX; x < tm.endX; x++ {
			newImage[Coord{x, y}] = tm.processLocation(Coord{x, y}, step)
		}
	}
	tm.image = newImage
}

func (tm *TrenchMap) count() {
	count := 0
	for _, val := range tm.image {
		if val == lp {
			count++
		}
	}
	fmt.Printf("Count is %d\n", count)
}

func readTrenchMap(text []string) TrenchMap {
	var tm TrenchMap
	tm.image = make(map[Coord]string)
	tm.startX = 0
	tm.startY = 0
	tm.endX = 0
	tm.endY = 0
	for y, line := range text {
		if y == 0 {
			tm.algorithm = append(tm.algorithm, strings.Split(line, "")...)
		} else {
			if line != "" {
				tm.endY++
				for x, val := range strings.Split(line, "") {
					tm.image[Coord{x, y - 2}] = val
					if y == 2 {
						tm.endX++
					}
				}
			}
		}
	}

	return tm
}

func part1(tm TrenchMap, steps int) {
	defer utils.TimeTrack(time.Now(), "Part 1")
	for i := 1; i <= steps; i++ {
		tm.processImage(i)
		//	tm.print()
	}
	tm.count()
}

func main() {
	text := utils.ReadInput(1)
	tm := readTrenchMap(text)
	tm.print()
	part1(tm, 50)
}
