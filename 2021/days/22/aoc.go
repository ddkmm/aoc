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

func readLine(line string) (word string, minX, minY, minZ, maxX, maxY, maxZ int) {
	w := strings.FieldsFunc(line, func(r rune) bool { return strings.ContainsRune(" =,.", r) })
	word = w[0]
	minX, _ = strconv.Atoi(w[2])
	minY, _ = strconv.Atoi(w[5])
	minZ, _ = strconv.Atoi(w[8])
	maxX, _ = strconv.Atoi(w[3])
	maxY, _ = strconv.Atoi(w[6])
	maxZ, _ = strconv.Atoi(w[9])

	return
}

func initCube(text []string, min int, max int) Cube {
	var c Cube
	c.state = make(map[Coord]string)
	c.min = min
	c.max = max
	for _, line := range text {
		word, minX, minY, minZ, maxX, maxY, maxZ := readLine(line)
		if (minX >= c.min && maxX <= c.max) &&
			(minY >= c.min && maxY <= c.max) &&
			(minZ >= c.min && maxZ <= c.max) {
			c.toggle(Coord{minX, minY, minZ}, Coord{maxX, maxY, maxZ}, word)
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

type Segment struct {
	xMin  int
	xMax  int
	state string
}

func (s *Segment) add(a Segment) bool {
	var left, right Segment
	if s.xMin < a.xMin {
		left = *s
		right = a
	} else {
		right = *s
		left = a
	}

	// Is there overlap?
	if left.xMax >= right.xMin && left.xMin <= right.xMax {
		if left.xMin <= right.xMin {
			if left.xMax <= right.xMax {
				/*
					simple overlap
						LLLLL
						  RRRRR
				*/
				s.xMin = left.xMin
				s.xMax = right.xMax
				return true
			} else if left.xMax > right.xMax {
				/*
					 Right is contained by left
							LLLLL
							 RRR
				*/
				s.xMin = left.xMin
				s.xMax = left.xMax
				return true
			}
		}
	}

	// No overlap
	return false
}

// finish this one up
func (left *Segment) subtract(right *Segment) bool {
	if left.xMax >= right.xMin && left.xMin <= right.xMax {
		if left.xMin <= right.xMin {
			if left.xMax <= right.xMax {
				/*
					simple overlap
						LLLLL
						  RRRRR
				*/
				left.xMax = right.xMin
				return true
			} else if left.xMax > right.xMax {
				/*
					 Right is contained by left
							LLLLL
							 RRR
				*/
				tempSeg := Segment{right.xMax, left.xMax, ""}
				left.xMax = right.xMin
				right.xMin = tempSeg.xMin
				right.xMax = tempSeg.xMax
				right.state = "on"
				return true
			} else if left.xMin >= right.xMin && left.xMax >= right.xMax {
				/*
					  LLLLL
					RRRRR
				*/
				left.xMin = right.xMax
				return true
			}
		}
	}
	return false

}

func part2(text []string) {
	defer utils.TimeTrack(time.Now(), "Part 2")
	var segList []Segment
	for _, line := range text {
		word, minX, _, _, maxX, _, _ := readLine(line)
		on := Segment{minX, maxX, word}
		segList = append(segList, on)

	}
	// Make a new list of processed segments
	// If we subtract and get a disjointed group,
	// then add it to the list and start over again
	// We will naturally skip any processed segments because they will be stateless

	seg := segList[0]
	fmt.Printf("Starting with ")
	fmt.Println(seg)
	for i := 1; i < len(segList)-1; i++ {
		fmt.Printf("Processing ")
		fmt.Println(segList[i])
		if segList[i].state == "on" {
			if seg.add(segList[i]) {
				fmt.Println(seg)
				fmt.Printf("add success, deleting ")
				fmt.Println(segList[i])
				segList[i].state = ""
				//				segList = append(segList[:1], segList[2:]...)
			}
		} else if segList[i].state == "off" {
			if seg.subtract(&segList[i]) {
				fmt.Printf("subtract success\n")
			}
		}
	}
	fmt.Println(segList)

}

func main() {
	text := utils.ReadInput(0)
	part1(text)
	part2(text)
}
