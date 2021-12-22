package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dkim/aoc/2021/utils"
)

type Coord struct {
	x int
	y int
	z int
}

type Cube struct {
	state map[Coord]string
	min   int
	max   int
}

func (c *Cube) print() {
	for z := c.min; z < c.max; z++ {

		for y := c.min; y < c.max; y++ {
			for x := c.min; x < c.max; x++ {
				if c.state[Coord{x, y, z}] == "on" {
					fmt.Printf("%d, %d, %d", x, y, z)
					fmt.Printf("\n")
				}
			}
		}
		fmt.Printf("\n")
	}
}

func (c *Cube) count() {
	count := 0
	for _, val := range c.state {
		if val == "on" {
			count++
		}
	}
	fmt.Printf("Count is %d\n", count)
}

func initCube(text []string, min int, max int) Cube {
	var c Cube
	c.state = make(map[Coord]string)
	c.min = min
	c.max = max
	for _, line := range text {
		w := strings.FieldsFunc(line, func(r rune) bool { return strings.ContainsRune(" =,.", r) })
		minX, _ := strconv.Atoi(w[2])
		minY, _ := strconv.Atoi(w[5])
		minZ, _ := strconv.Atoi(w[8])
		maxX, _ := strconv.Atoi(w[3])
		maxY, _ := strconv.Atoi(w[6])
		maxZ, _ := strconv.Atoi(w[9])
		if (minX >= c.min && maxX <= c.max) &&
			(minY >= c.min && maxY <= c.max) &&
			(minZ >= c.min && maxZ <= c.max) {
			c.toggle(Coord{minX, minY, minZ}, Coord{maxX, maxY, maxZ}, w[0])
		}
	}
	return c
}

func (c *Cube) toggle(min Coord, max Coord, state string) {
	for z := min.z; z <= max.z; z++ {
		for y := min.y; y <= max.y; y++ {
			for x := min.x; x <= max.x; x++ {
				c.state[Coord{x, y, z}] = state
			}
		}
	}
}

func part1(text []string) {
	defer utils.TimeTrack(time.Now(), "Part 1")
	c := initCube(text, -50, 50)
	c.count()
}

func part2(text []string) {
	defer utils.TimeTrack(time.Now(), "Part 2")
}

func main() {
	text := utils.ReadInput(1)
	part1(text)
	part2(text)
}
