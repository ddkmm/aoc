package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dkim/aoc/2021/utils"
)

func processInput(text []string) (snails []Snail) {
	for _, line := range text {
		var ss SnailStack
		for _, c := range strings.Split(line, "") {
			switch c {
			case "[":
				snail := &Snail{nil, nil, nil, -1, -1}
				ss.push(snail)
			case ",":

			case "]":
				s := ss.pop()
				if ss.len() == 0 {
					snails = append(snails, *s)
				} else {
					s2 := ss.pop()
					if s2.sLeft == nil && s2.vLeft < 0 {
						s2.sLeft = s
						s.sUp = s2
					} else if s2.sRight == nil && s2.vRight < 0 {
						s2.sRight = s
						s.sUp = s2
					}
					ss.push(s2)
				}
			default:
				s := ss.pop()
				num, _ := strconv.Atoi(c)
				if s.sLeft == nil && s.vLeft < 0 {
					s.vLeft = num
				} else if s.sRight == nil && s.vRight < 0 {
					s.vRight = num
				}
				ss.push(s)
			}
		}
	}
	return
}

func add(s1 *Snail, s2 *Snail) (s3 *Snail) {
	s3 = &Snail{nil, s1, s2, -1, -1}
	s1.setUp(s3)
	s2.setUp(s3)
	s1.print()
	fmt.Print(" + ")
	s2.print()
	fmt.Print(" = ")
	s3.print()
	fmt.Println()
	return s3
}

func work(snails []Snail) (sum *Snail) {
	snails[0].reduce()
	sum = &snails[0]
	for i := 0; i < len(snails)-1; i++ {
		sum = add(sum, &snails[i+1])
		sum.print()
		sum.reduce()
	}
	sum.print()
	fmt.Println()

	return
}

func part1(text []string) {
	defer utils.TimeTrack(time.Now(), "Part 1")
	snails := processInput(text)
	for _, s := range snails {
		s.print()
		fmt.Println()
	}
	sum := work(snails)
	fmt.Println("Sum is ")
	sum.print()
	fmt.Printf("\nPart 1: %d\n", sum.magnitude())
}

func main() {
	text := utils.ReadInput(0)
	part1(text)
}