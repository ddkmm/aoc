package main

import (
	"fmt"
	"time"

	"github.com/dkim/aoc/2021/utils"
)

type Beacon struct {
	x int
	y int
}

func (b *Beacon) print() {
	fmt.Printf("%d, %d\n", b.x, b.y)
}

func part1() {
	defer utils.TimeTrack(time.Now(), "Part 1")
}

func main() {
	fmt.Println("Hello")
	text := utils.ReadInput(0)
	fmt.Println(text[0])
	part1()
}
