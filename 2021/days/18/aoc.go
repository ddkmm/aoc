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

func (ss *SnailStack) clear() {
	ss.Stack = nil
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
		//		fmt.Printf("**** Processing %s\n", line)
		var ss SnailStack
		depth := 0
		for _, c := range strings.Split(line, "") {
			switch c {
			case "[":
				depth++
				snail := &Snail{nil, nil, nil, -1, -1, depth}
				/*
					fmt.Printf("New snail ")
					snail.print()
					fmt.Print(" ")
					fmt.Print(&snail)
					fmt.Printf(" at depth %d\n", depth)
				*/
				ss.push(snail)
			case ",":

			case "]":
				depth--
				s := ss.pop()
				if ss.len() == 0 {
					/*
						fmt.Printf("Completed snail ")
						s.print()
						fmt.Print(" ")
						fmt.Println(&s)
					*/
					snails = append(snails, *s)
				} else {
					s2 := ss.pop()
					//					fmt.Println(&s2)
					if s2.sLeft == nil && s2.vLeft < 0 {
						/*
							fmt.Printf("Left adding snail ")
							s.print()
							fmt.Printf(" with parent ")
							s2.print()
							fmt.Println(&s2)
						*/
						s2.sLeft = s
						s.sUp = s2
					} else if s2.sRight == nil && s2.vRight < 0 {
						/*
							fmt.Printf("Right adding snail ")
							s.print()
							fmt.Printf(" with parent ")
							s2.print()
							fmt.Println(&s2)
						*/
						s2.sRight = s
						s.sUp = s2
					}
					/*
						fmt.Printf("Pushing ")
						s2.print()
						fmt.Printf(" ")
						fmt.Print(&s2)
						fmt.Println(" onto stack")
					*/
					ss.push(s2)
				}
			default:
				s := ss.pop()
				num, _ := strconv.Atoi(c)
				if s.sLeft == nil && s.vLeft < 0 {
					/*
						fmt.Printf("Left adding %s to snail ", c)
						s.print()
						fmt.Println()
					*/
					s.vLeft = num
				} else if s.sRight == nil && s.vRight < 0 {
					/*
						fmt.Printf("Right adding %s to snail ", c)
						s.print()
						fmt.Println()
					*/
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

func (s *Snail) reduceDepth() {
	s.depth--
	if s.sLeft != nil {
		s.sLeft.reduceDepth()
	}
	if s.sRight != nil {
		s.sRight.reduceDepth()
	}
}

func (s *Snail) findDepth(depth int, ss *SnailStack) bool {
	ret := false
	ss.push(s)
	if s.sLeft == nil && s.sRight == nil {
		if ss.len() == depth {
			ret = true
		}
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

func (s *Snail) findSplit(ss *SnailStack) bool {
	ret := false
	ss.push(s)
	if s.vLeft > 9 || s.vRight > 9 {
		ret = true
	}
	if s.sLeft != nil {
		ret = s.sLeft.findSplit(ss)
		if !ret {
			ss.pop()
		}
	}
	if !ret && s.sRight != nil {
		ret = s.sRight.findSplit(ss)
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
		if s.explode() {
			res = true
			e++
			fmt.Print("EXPLODE: ")
		} else if s.split() {
			res = true
			sp++
			fmt.Print("SPLIT: ")
		} else {
			res = false
		}
		s.print()
		fmt.Println()
		fmt.Printf("%d explodes, %d splits\n", e, sp)
	}
}

func (s *Snail) split() (res bool) {
	var ss SnailStack
	res = false
	if s.findSplit(&ss) {
		sp := ss.pop()
		if sp.sLeft == nil && sp.vLeft > 9 {
			temp := float64(sp.vLeft) / 2
			newSnail := Snail{sp, nil, nil, int(math.Floor(temp)), int(math.Ceil(temp)), sp.depth + 1}
			sp.sLeft = &newSnail
			sp.vLeft = -1
			res = true
		} else if sp.sRight == nil && sp.vRight > 9 {
			temp := float64(sp.vRight) / 2
			newSnail := Snail{sp, nil, nil, int(math.Floor(temp)), int(math.Ceil(temp)), sp.depth + 1}
			sp.sRight = &newSnail
			sp.vRight = -1
			res = true
		}
	}
	return
}

func (s *Snail) find(dir string, ss *SnailStack) (retSnail *Snail) {
	if ss.len() == 0 {
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

func (ss *SnailStack) print() {
	for i := ss.len(); i > 0; i-- {
		fmt.Printf("Depth %d: ", ss.Stack[i-1].depth)
		ss.Stack[i-1].print()
		fmt.Print(" ")
		fmt.Println(ss.Stack[i-1])
	}
}

func (s *Snail) explode_old(ss SnailStack) {
	leftEdge := false
	rightEdge := false
	fmt.Printf("** Exploding\n")
	s.print()
	fmt.Println()
	ss.print()
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
			// Left neighbor is in same pair
			if newS.sLeft == nil {
				// Just a value [A, [X, B]]
				// [A+X, 0]
				newS.vLeft += exploder.vLeft
				newS.vRight = 0
				newS.sRight = nil
			} else {
				// pair of pairs [[A,B],[X,C]]
				// This can't happen because we're at same depth

			}
			/*
				newS.vLeft += exploder.vLeft
				exploder.vLeft = 0
			*/
			// How do we re-connect?
			//				newS.vRight = -1
		} else if newS.getUp().sRight == exploder {
			// Right neighbor has same parent
			newS.vRight += exploder.vLeft
			newS.getUp().sRight = nil
			newS.getUp().vRight = 0
		} else if newS.sLeft != nil && newS.vLeft == -1 {
			// left neighbor is in different pair and is in a pair
			newS.sRight.vRight += exploder.vLeft
			exploder.vLeft = 0
		} else if newS.sLeft == nil && newS.vLeft != -1 {
			// Left neighbor is in different pair and just a value
			newS.vLeft += exploder.vLeft
			exploder.vLeft = 0
			exploder.getUp().sLeft = nil
			exploder.getUp().vLeft = 0
		} else {
			newS.vLeft += exploder.vLeft
			exploder.vLeft = 0
		}
		exploder.reduceDepth()
	} else {
		fmt.Printf("No snail on the left\n")
		// TODO: Handle orphaned case
		leftEdge = true
		exploder.reduceDepth()
	}
	exploder = rightStack.pop()
	newS = exploder.find("right", &rightStack)
	if newS != nil {
		fmt.Printf("Right snail: ")
		newS.print()
		fmt.Println()
		if newS.sLeft == exploder {
			// Right neighbor is in same pair
			newS.vRight += exploder.vRight
			newS.vLeft = 0
			//				newS.vRight = -1
		} else if newS.getUp().sLeft == exploder {
			// Right neighbor has same parent
			newS.vLeft += exploder.vRight
			newS.getUp().sLeft = nil
			newS.getUp().vLeft = 0
		} else if newS.sLeft != nil && newS.vLeft == -1 {
			// left neighbor is in different pair and is in a pair
			newS.sLeft.vLeft += exploder.vRight
			exploder.vRight = 0
		} else if newS.sLeft == nil && newS.vRight != -1 {
			// left neighbor is in different pair and just a value
			newS.vLeft += exploder.vRight
			exploder.vRight = 0
		} else {
			newS.vRight += exploder.vRight
			exploder.vRight = 0
		}
		exploder.reduceDepth()
	} else {
		fmt.Printf("No snail on the right\n")
		// TODO: Handle orphaned case
		rightEdge = true
		exploder.reduceDepth()
	}
	if leftEdge {
		exploder.getUp().sLeft = nil
		exploder.getUp().vLeft = 0
	} else if rightEdge {
		exploder.getUp().sRight = nil
		exploder.getUp().vRight = 0
	}
}

func (s *Snail) explode_new(ss SnailStack) {
	// Top of the stack has the snail at the lowest depth
	// s is the 'root' snail
	ss.print()
	exploder := ss.pop()
	leftStack := ss
	rightStack := ss
	var isLeft bool

	fmt.Print("Exploding: ")
	exploder.print()
	fmt.Println()

	// Identify the neighbors
	leftSnail := exploder.find("left", &leftStack)
	fmt.Print("Left neighbor found: ")
	leftSnail.print()
	fmt.Println()
	rightSnail := exploder.find("right", &rightStack)
	fmt.Print("Right neighbor found: ")
	rightSnail.print()
	fmt.Println()

	/*
		1. Are we exploding the left snail or right snail?
	*/
	if exploder.getUp().sLeft == exploder {
		isLeft = true
	} else if exploder.getUp().sRight == exploder {
		isLeft = false
	}

	/*
		2. Handle the near neighbor (Same parent snail)
			a. Is the other side a snail, a value, or missing?
				1. If it's a value, just add our side's value to it
				2. If it's a snail, add our side's value to the correct side value
				3. If it's missing, do ??
	*/
	fmt.Print("Handle near neighbor first")
	if isLeft {
		fmt.Println(" to the right")
	} else {
		fmt.Println(" to the left")
	}
	if isLeft {
		fmt.Printf("%d + ", exploder.vRight)
		rightSnail.print()
		// We are the left side, so the near neighbor is the right side
		// and the near neighbor must be a value?
		if rightSnail == exploder.getUp() {
			rightSnail.vLeft += exploder.vRight

		} else if rightSnail != nil {
			rightSnail.vLeft += exploder.vRight
		}
		fmt.Print(" = ")
		rightSnail.print()
		fmt.Println()
	} else {
		leftSnail.print()
		fmt.Printf(" + %d", exploder.vLeft)
		if leftSnail == exploder.getUp() {
			leftSnail.vLeft += exploder.vLeft
		} else if leftSnail != nil {
			leftSnail.vRight += exploder.vLeft
		}
		fmt.Print(" = ")
		leftSnail.print()
		fmt.Println()
	}
	/*
		3. Handle the far neighbor (Different parent snail)
			a. Is the far neighbor a snail or a value?
				1. If it's a value, just add our side's value to it
				2. If it's a snail, add our side's value to the correct side value
				3. If it's missing, do ??
	*/
	fmt.Print("Now handle far neighbor ")
	if isLeft {
		fmt.Println("to the left")
	} else {
		fmt.Println("to the right")
	}

	if isLeft {
		leftSnail.print()
		fmt.Printf(" + %d", exploder.vLeft)
		// We are left so the far neighbor is also on the left
		// And we want the right most value of it
		if leftSnail == nil {
			fmt.Println("Empty neighbor")
		} else if leftSnail.sRight == nil && leftSnail.vRight != -1 {
			// Case 1, right side is just a value
			leftSnail.vRight += exploder.vLeft
		} else if leftSnail.sRight != nil {
			// Case 2, right side is another snail
			// Get right value of right side
			if leftSnail.sRight.sRight == nil {
				if leftSnail.sRight.vRight != -1 {
					leftSnail.sRight.vRight += exploder.vLeft
				}
			} else if leftSnail.sLeft == nil {
				leftSnail.vLeft += exploder.vLeft
			} else {
				fmt.Println("Case 2 error")
			}
		}
		fmt.Print(" = ")
		leftSnail.print()
		fmt.Println()
	} else {
		fmt.Printf("%d + ", exploder.vRight)
		rightSnail.print()
		// We are the right so the far neighbor is also on the right
		// And we want to add to the left most value
		if rightSnail == nil {
			fmt.Println("Empty neighbor")
		} else if rightSnail.sLeft == nil && rightSnail.vLeft != -1 {
			rightSnail.vLeft += exploder.vRight
		} else if rightSnail.sLeft != nil {
			// Case 2, left side is another snail
			// Get left value of left side
			if rightSnail.sLeft.sLeft == nil {
				if rightSnail.sLeft.vLeft != -1 {
					rightSnail.sLeft.vLeft += exploder.vRight
				}
			} else if rightSnail.sRight == nil {
				rightSnail.vRight += exploder.vRight
			} else {
				fmt.Println("Case 2 error")
			}
		}
		fmt.Print(" = ")
		rightSnail.print()
		fmt.Println()
	}

	/*
		4. Fixup the snail we just exploded
			a. From the parent, set the exploded side pointer as nil with value 0
	*/
	if isLeft {
		exploder.sUp.vLeft = 0
		exploder.sUp.sLeft = nil
	} else {
		exploder.sUp.vRight = 0
		exploder.sUp.sRight = nil
	}
}

func (s *Snail) explode() (res bool) {
	var ss SnailStack
	// Find the first snail with depth 5
	if s.findDepth(5, &ss) {
		s.explode_new(ss)
		res = true
	} else {
		res = false
	}
	return
}

func work(snails []Snail) (sum *Snail) {
	snails[0].reduce()
	sum = &snails[0]
	for i := 0; i < len(snails)-1; i++ {
		sum = add(sum, &snails[i+1])
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
