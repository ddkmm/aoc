package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/dkim/aoc/2021/utils"
)

type Snail struct {
	sUp    *Snail
	sLeft  *Snail
	sRight *Snail
	vLeft  int
	vRight int
	depth  int
}

func (s *Snail) print() {
	if s != nil {
		fmt.Printf("[")
		if s.sLeft != nil {
			s.sLeft.print()
		} else {
			fmt.Printf("%d", s.vLeft)
		}
		fmt.Printf(",")
		if s.sRight != nil {
			s.sRight.print()
		} else {
			fmt.Printf("%d", s.vRight)
		}
		fmt.Printf("]")
	} else {
		fmt.Printf("Empty snail\n")
	}
}

type SnailStack struct {
	Stack []*Snail
}

func (ss *SnailStack) push(snail *Snail) {
	ss.Stack = append(ss.Stack, snail)
}

func (ss *SnailStack) pop() (snail *Snail) {
	n := len(ss.Stack)
	if n > 0 {
		n--
		snail = ss.Stack[n]
		ss.Stack = ss.Stack[:n]
	}
	return
}

func (ss *SnailStack) len() int {
	return len(ss.Stack)
}

func processInput(text []string) (snails []Snail) {
	for _, line := range text {
		fmt.Printf("**** Processing %s\n", line)
		var ss SnailStack
		depth := 0
		for _, c := range strings.Split(line, "") {
			switch c {
			case "[":
				depth++
				snail := &Snail{nil, nil, nil, -1, -1, depth}
				fmt.Printf("New snail ")
				snail.print()
				fmt.Print(" ")
				fmt.Print(&snail)
				fmt.Printf(" at depth %d\n", depth)
				ss.push(snail)
			case ",":

			case "]":
				depth--
				s := ss.pop()
				if ss.len() == 0 {
					fmt.Printf("Completed snail ")
					s.print()
					fmt.Print(" ")
					fmt.Println(&s)
					snails = append(snails, *s)
					if s.findDepth(2) != nil {
						fmt.Println("Quick check:")
						if s.sLeft != nil {
							s.sLeft.sUp.print()
							fmt.Println()
						}
						if s.sRight != nil {
							s.sRight.sUp.print()
							fmt.Println()
						}
					}
				} else {
					s2 := ss.pop()
					fmt.Println(&s2)
					if s2.sLeft == nil && s2.vLeft < 0 {
						fmt.Printf("Left adding snail ")
						s.print()
						fmt.Printf(" with parent ")
						s2.print()
						fmt.Println(&s2)
						s2.sLeft = s
						s.sUp = s2
					} else if s2.sRight == nil && s2.vRight < 0 {
						fmt.Printf("Right adding snail ")
						s.print()
						fmt.Printf(" with parent ")
						s2.print()
						fmt.Println(&s2)
						s2.sRight = s
						s.sUp = s2
					}
					fmt.Printf("Pushing ")
					s2.print()
					fmt.Printf(" ")
					fmt.Print(&s2)
					fmt.Println(" onto stack")
					ss.push(s2)
				}
			default:
				s := ss.pop()
				num, _ := strconv.Atoi(c)
				if s.sLeft == nil && s.vLeft < 0 {
					fmt.Printf("Left adding %s to snail ", c)
					s.print()
					fmt.Println()
					s.vLeft = num
				} else if s.sRight == nil && s.vRight < 0 {
					fmt.Printf("Right adding %s to snail ", c)
					s.print()
					fmt.Println()
					s.vRight = num
				}
				ss.push(s)
			}
		}
	}
	return
}

func (s *Snail) setUp(up *Snail) {
	s.sUp = up
}

func (s *Snail) getUp() *Snail {
	return s.sUp
}

func (s *Snail) magnitude() int {
	mag := 0
	if s.sLeft != nil {
		mag += s.sLeft.magnitude() * 3
	} else if s.sLeft == nil {
		mag += s.vLeft * 3
	}
	if s.sRight != nil {
		mag += s.sRight.magnitude() * 2
	} else if s.sRight == nil {
		mag += s.vRight * 2
	}
	return mag
}

func (s *Snail) addDepth() {
	s.depth++
	if s.sLeft != nil {
		s.sLeft.addDepth()
	}
	if s.sRight != nil {
		s.sRight.addDepth()
	}
}

func (s *Snail) findDepth(depth int) *Snail {
	if s.depth == depth {
		return s
	}
	var retSnail *Snail = nil
	if s.sLeft != nil {
		retSnail = s.sLeft.findDepth(depth)
	}
	if retSnail == nil && s.sRight != nil {
		retSnail = s.sRight.findDepth(depth)
	}
	return retSnail
}

func add(s1 *Snail, s2 *Snail) *Snail {
	s1.sUp.print()
	s2.sUp.print()
	s := &Snail{nil, s1, s2, -1, -1, 1}
	s1.addDepth()
	s2.addDepth()
	fmt.Println(&(s1.sUp))
	fmt.Println(&(s2.sUp))
	s1.setUp(s)
	s2.setUp(s)
	s1.sUp.print()
	fmt.Println()
	s2.sUp.print()
	fmt.Println()
	s1.sUp = s
	s1.sUp.print()
	fmt.Println()
	s2.sUp.print()
	fmt.Println()
	return s
}

func (s *Snail) reduce() {
	s.explode()
	s.split()
}

func (s *Snail) split() {
	if s.sLeft == nil && s.vLeft > 9 {
		temp := float64(s.vLeft) / 2
		newSnail := Snail{s, nil, nil, int(math.Floor(temp)), int(math.Ceil(temp)), s.depth + 1}
		s.sLeft = &newSnail
	} else if s.sRight == nil && s.vRight > 9 {
		temp := float64(s.vLeft) / 2
		newSnail := Snail{s, nil, nil, int(math.Floor(temp)), int(math.Ceil(temp)), s.depth + 1}
		s.sRight = &newSnail
	}
}

/*
func (s *Snail) findLeft() (leftSnail *Snail) {
	if s.sUp == nil {
		leftSnail = nil
	} else if s.sUp.sLeft != nil && s.sUp.sLeft != s {
		leftSnail = s.sUp.sLeft
	} else {
		leftSnail = s.sUp.findLeft()
	}
	return
}
*/

func (s *Snail) findLeft(ss *SnailStack) {
	fmt.Printf("Find left for ")
	s.print()
	fmt.Printf(" at depth %d\n", s.depth)
	fmt.Println(s)

	if s.getUp() == nil {
		s.print()
		fmt.Println(" has no parents")
		fmt.Println(s)
		return
	}
	fmt.Printf("Parent is ")
	s.sUp.print()
	fmt.Print(" ")
	fmt.Println(s.sUp)
	if s.sUp.sLeft != nil && s.sUp.sLeft != s {
		s.sUp.sLeft.print()
		fmt.Println(" found")
		ss.push(s.getUp().sLeft)
	} else {
		fmt.Printf("Not found yet. Now find from ")
		s.sUp.print()
		fmt.Println()
		fmt.Println(s.sUp)
		s.getUp().findLeft(ss)
	}
}

func (s *Snail) findRight(ss *SnailStack) {
	if s.sUp == nil {
		return
	}
	if s.sUp.sRight != nil {
		ss.push(s.sUp.sRight)
	} else {
		s.sUp.findRight(ss)
	}
}

func (s *Snail) explode() {
	var ss SnailStack
	// Find the first snail with depth 5
	exploder := s.findDepth(3)
	if exploder != nil {
		s.print()
		fmt.Printf(" has snail ")
		exploder.print()
		fmt.Printf(" at depth %d\n", exploder.depth)
		exploder.findLeft(&ss)
		newS := ss.pop()
		if newS != nil {
			fmt.Printf("Left snail: ")
			fmt.Println()
			newS.print()
		} else {
			fmt.Printf("No snail on the left\n")
		}
		exploder.findRight(&ss)
		newS = ss.pop()
		if newS != nil {
			fmt.Printf("Right snail: ")
			newS.print()
			fmt.Println()
		} else {
			fmt.Printf("No snail on the right\n")
		}
	}
}

func work(snails []Snail) (sum Snail) {
	fmt.Printf("Add ")
	snails[0].print()
	fmt.Printf(" and ")
	snails[1].print()
	fmt.Println()
	snail := add(&snails[0], &snails[1])
	snail.print()
	fmt.Println("\nExplode")
	snail.explode()
	fmt.Printf("Now try again with ")
	fmt.Println(snails[2])
	snails[2].explode()
	fmt.Println("\nSplit")
	snail.split()
	snail.print()
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
	fmt.Printf("Part 1: %d\n", sum.magnitude())
}

func main() {
	text := utils.ReadInput(0)
	part1(text)
}
