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

func (s *Snail) findDepth(depth int, ss *SnailStack) bool {
	ret := false
	ss.push(s)
	if s.depth == depth {
		ret = true
	}
	if s.sLeft != nil {
		ret = s.sLeft.findDepth(depth, ss)
		if !ret {
			ss.pop()
		}
	}
	if !ret && s.sRight != nil {
		ret = s.sRight.findDepth(depth, ss)
		if !ret {
			ss.pop()
		}
	}
	return ret
}

func add(s1 *Snail, s2 *Snail) (s3 *Snail) {
	s3 = &Snail{nil, s1, s2, -1, -1, 1}
	s1.setUp(s3)
	s2.setUp(s3)
	s3.sLeft.addDepth()
	s3.sRight.addDepth()
	s1.print()
	fmt.Print(" + ")
	s2.print()
	fmt.Print(" = ")
	s3.print()
	fmt.Println()
	return s3
}

func (s *Snail) reduce() {
	res := true
	e := 0
	sp := 0
	for res {
		s.print()
		fmt.Println()
		res = s.explode()
		s.print()
		fmt.Println()
		if res {
			e++
		}
		if !res {
			res = s.split()
			if res {
				sp++
			}
		}
		fmt.Printf("%d explodes, %d splits\n", e, sp)
	}
}

func (s *Snail) split() (res bool) {
	res = false
	if s.sLeft == nil && s.vLeft > 9 {
		temp := float64(s.vLeft) / 2
		newSnail := Snail{s, nil, nil, int(math.Floor(temp)), int(math.Ceil(temp)), s.depth + 1}
		s.sLeft = &newSnail
		res = true
	} else if s.sRight == nil && s.vRight > 9 {
		temp := float64(s.vLeft) / 2
		newSnail := Snail{s, nil, nil, int(math.Floor(temp)), int(math.Ceil(temp)), s.depth + 1}
		s.sRight = &newSnail
		res = true
	}
	return
}

func (s *Snail) find(dir string, ss *SnailStack) (retSnail *Snail) {
	if ss.len() == 0 {
		s.print()
		fmt.Println(" has no parents")
		fmt.Println(s)
		retSnail = nil
	} else {
		snail := ss.pop()
		if dir == "left" {
			if snail.sLeft != nil && snail.sLeft != s {
				//ss.push(snail.sLeft)
				retSnail = snail.sLeft
			} else if snail.sLeft == nil && snail.sRight == s {
				//ss.push(snail)
				retSnail = snail
			} else if retSnail == nil {
				return snail.find(dir, ss)
			}
		} else if dir == "right" {
			if snail.sRight != nil && snail.sRight != s {
				//ss.push(snail.sRight)
				retSnail = snail.sRight
			} else if snail.sRight == nil && snail.sLeft == s {
				//ss.push(snail)
				retSnail = snail
			} else if retSnail == nil {
				return snail.find(dir, ss)
			}
		}
	}
	return
}

func (s *Snail) explode() (res bool) {
	var ss SnailStack
	// Find the first snail with depth 5
	if s.findDepth(5, &ss) {
		fmt.Printf("** Exploding ")
		s.print()
		fmt.Println()
		/*
			for i := ss.len(); i > 0; i-- {
				fmt.Printf("Depth %d: ", ss.Stack[i-1].depth)
				fmt.Println(ss.Stack[i-1])
			}
		*/
		leftStack := ss
		rightStack := ss
		exploder := leftStack.pop()
		s.print()
		fmt.Printf(" has snail ")
		exploder.print()
		fmt.Printf(" at depth %d\n", exploder.depth)
		newS := exploder.find("left", &leftStack)
		if newS != nil {
			fmt.Printf("Left snail: ")
			newS.print()
			fmt.Println()
			if newS.sRight == exploder {
				newS.vLeft += exploder.vLeft
				newS.vRight = 0
				newS.sRight = nil
				exploder = nil
			} else {
				newS.vRight += exploder.vLeft
			}
		} else {
			fmt.Printf("No snail on the left\n")
		}
		exploder = rightStack.pop()
		newS = exploder.find("right", &rightStack)
		if newS != nil {
			fmt.Printf("Right snail: ")
			newS.print()
			fmt.Println()
			if newS.sLeft == exploder {
				newS.vRight += exploder.vRight
				newS.vLeft = 0
				newS.sLeft = nil
				exploder = nil
			} else {
				newS.vLeft += exploder.vRight
			}
		} else {
			fmt.Printf("No snail on the right\n")
		}
		res = true
	} else {
		res = false
	}
	return
}

func work(snails []Snail) (sum Snail) {
	tmpSum := &snails[0]
	for i := 0; i < len(snails)-2; i++ {
		tmpSum = add(tmpSum, &snails[i+1])
		tmpSum.reduce()
		tmpSum.print()
		fmt.Println()
	}
	tmpSum.print()
	fmt.Println()

	fmt.Printf("Add ")
	snails[0].print()
	fmt.Printf(" and ")
	snails[1].print()
	fmt.Println()
	snail := add(&snails[0], &snails[1])
	snail.reduce()
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
