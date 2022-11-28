package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/dkim/aoc/2021/utils"
)

const side = ">"
const down = "v"

type Coord struct {
	x int
	y int
}

type Cucumbers struct {
	area map[Coord]string
	xMax int
	yMax int
}

func (c *Cucumbers) print() {
	for y := 0; y <= c.yMax; y++ {
		for x := 0; x <= c.xMax; x++ {
			fmt.Printf("%s", c.area[Coord{x, y}])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func (c *Cucumbers) stepRight() (flag bool) {
	tempArea := make(map[Coord]string)
	flag = false
	for key, value := range c.area {
		if value == side {
			// Look to the right
			var tempCoord Coord
			if c.area[Coord{key.x + 1, key.y}] == "" {
				// Wrap around
				tempCoord = Coord{0, key.y}
			} else {
				tempCoord = Coord{key.x + 1, key.y}
			}
			if c.area[tempCoord] == "." {
				tempArea[tempCoord] = side
				tempArea[key] = "."
				flag = true
			}
		}
	}
	// Copy changes back over
	for key, value := range tempArea {
		c.area[key] = value
	}
	return
}

func (c *Cucumbers) stepDown() (flag bool) {
	tempArea := make(map[Coord]string)
	flag = false
	for key, value := range c.area {
		if value == down {
			// Look down
			var tempCoord Coord
			if c.area[Coord{key.x, key.y + 1}] == "" {
				// Wrap around
				tempCoord = Coord{key.x, 0}
			} else {
				tempCoord = Coord{key.x, key.y + 1}
			}
			if c.area[tempCoord] == "." {
				tempArea[tempCoord] = down
				tempArea[key] = "."
				flag = true
			}
		}
	}
	// Copy changes back over
	for key, value := range tempArea {
		c.area[key] = value
	}
	return
}

func (c *Cucumbers) step() {
	i := 0
	flag := true
	for flag {
		flag = c.stepRight()
		flag = c.stepDown() || flag
		i++
	}
	fmt.Printf("Final step: %d\n", i)
}

func readCucumbers(text []string) Cucumbers {
	var c Cucumbers
	c.area = make(map[Coord]string)
	for y, line := range text {
		c.yMax = y
		for x, val := range strings.Split(line, "") {
			c.xMax = x
			c.area[Coord{x, y}] = val
		}
	}
	return c
}

func part1(c Cucumbers) {
	defer utils.TimeTrack(time.Now(), "Part 1")
	c.step()
}

func main() {
	text := utils.ReadInput(1)
	c := readCucumbers(text)
	c.print()
	part1(c)
	c.print()
}
