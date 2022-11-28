package main

import (
	"fmt"
	"time"

	"github.com/dkim/aoc/2021/utils"
)

type Die struct {
	val   int
	count int
}

func (d *Die) roll() int {
	d.val = d.val % 100
	d.val++
	d.count++
	return (d.val)
}

func part1(p1 int, p2 int) {
	defer utils.TimeTrack(time.Now(), "Part 1")
	d := Die{0, 0}
	s1 := 0
	s2 := 0
	winner := 0
	for {
		// player 1
		p1 += d.roll()
		p1 += d.roll()
		p1 += d.roll()
		p1 = p1 % 10
		if p1 == 0 {
			p1 = 10
		}
		s1 += p1
		if s1 >= 1000 {
			fmt.Printf("Player 1 wins with %d\n", s1)
			winner = 1
			break
		}
		// player 2
		p2 += d.roll()
		p2 += d.roll()
		p2 += d.roll()
		p2 = p2 % 10
		if p2 == 0 {
			p2 = 10
		}
		s2 += p2
		if s2 >= 1000 {
			fmt.Printf("Player 2 wins with %d\n", s2)
			winner = 2
			break
		}
	}
	var answer int
	if winner == 1 {
		answer = s2
	} else {
		answer = s1
	}
	fmt.Printf("losing score %d x roll count %d = %d\n", answer, d.count, answer*d.count)

}

func part2(p1 int, p2 int) {
	defer utils.TimeTrack(time.Now(), "Part 2")
	/*
		1,1,1
		1,1,2
		1,2,1
		2,1,1
		1,2,2
		2,2,1
	*/

}

func main() {
	part1(4, 8)
	//	part1(10, 1)
	part2(4, 8)
}
