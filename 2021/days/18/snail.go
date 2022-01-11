package main

import (
	"fmt"
	"math"
)

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

func (ss *SnailStack) print() {
	for i := ss.len(); i > 0; i-- {
		ss.Stack[i-1].print()
		fmt.Print(" ")
		fmt.Println(ss.Stack[i-1])
	}
}

type Snail struct {
	sUp    *Snail
	sLeft  *Snail
	sRight *Snail
	vLeft  int
	vRight int
	//	depth  int
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
	if s.vLeft > 9 { //|| s.vRight > 9 {
		ret = true
	} else {
		if s.sLeft != nil {
			ret = s.sLeft.findSplit(ss)
			if !ret {
				ss.pop()
			}
		}
	}
	if s.vRight > 9 {
		ret = true
	} else {
		if !ret && s.sRight != nil {
			ret = s.sRight.findSplit(ss)
			if !ret {
				ss.pop()
			}
		}
	}
	return ret
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
		fmt.Print("Splitting ")
		sp.print()
		fmt.Println()
		if sp.sLeft == nil && sp.vLeft > 9 {
			temp := float64(sp.vLeft) / 2
			newSnail := Snail{sp, nil, nil, int(math.Floor(temp)), int(math.Ceil(temp))}
			sp.sLeft = &newSnail
			sp.vLeft = -1
			res = true
		} else if sp.sRight == nil && sp.vRight > 9 {
			temp := float64(sp.vRight) / 2
			newSnail := Snail{sp, nil, nil, int(math.Floor(temp)), int(math.Ceil(temp))}
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

func (s *Snail) getMostLeft() *Snail {
	if s.sLeft == nil {
		return s
	} else {
		return s.sLeft.getMostLeft()
	}
}

func (s *Snail) getMostRight() *Snail {
	if s.sRight == nil {
		return s
	} else {
		return s.sRight.getMostRight()
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
	if ss.Stack[ss.len()-1].sLeft == exploder {
		isLeft = true
	} else if ss.Stack[ss.len()-1].sRight == exploder {
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
		if rightSnail == ss.Stack[ss.len()-1] {
			rightSnail.vRight += exploder.vRight

		} else if rightSnail != nil {
			rightSnail.vLeft += exploder.vRight
		}
		fmt.Print(" = ")
		rightSnail.print()
		fmt.Println()
	} else {
		leftSnail.print()
		fmt.Printf(" + %d", exploder.vLeft)
		if leftSnail == ss.Stack[ss.len()-1] {
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
		} else {
			if leftSnail.sRight != nil && leftSnail.sRight.getMostLeft() == exploder {
				leftSnail.vLeft += exploder.vLeft
			} else {
				leftSnail.getMostRight().vRight += exploder.vLeft
			}
			/*
				if leftSnail.sRight == nil && leftSnail.vRight != -1 {
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
			*/
			fmt.Print(" = ")
			leftSnail.print()
			fmt.Println()
		}
	} else {
		fmt.Printf("%d + ", exploder.vRight)
		rightSnail.print()
		// We are the right so the far neighbor is also on the right
		// And we want to add to the left most value
		if rightSnail == nil {
			fmt.Println("Empty neighbor")
		} else {
			if rightSnail.sLeft != nil && rightSnail.sLeft.getMostRight() == exploder {
				rightSnail.vRight += exploder.vRight
			} else {
				rightSnail.getMostLeft().vLeft += exploder.vRight
			}
			/*
				if rightSnail.sLeft == nil && rightSnail.vLeft != -1 {
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
			*/
			fmt.Print(" = ")
			rightSnail.print()
			fmt.Println()
		}
	}

	/*
		4. Fixup the snail we just exploded
			a. From the parent, set the exploded side pointer as nil with value 0
	*/
	fmt.Println("Snail fixup")
	fmt.Println("snail stack: ")
	ss.print()
	fmt.Print("Exploder parent: ")
	fmt.Print(&exploder.sUp)
	fmt.Print(" is ")
	fmt.Println(exploder.getUp())
	if isLeft {
		s.print()
		fmt.Println()
		ss.Stack[ss.len()-1].vLeft = 0
		ss.Stack[ss.len()-1].sLeft = nil
		exploder.sUp.vLeft = 0
		exploder.sUp.sLeft = nil
		exploder.sUp.print()
		fmt.Print("\nExploder sUp: ")
		fmt.Println(exploder.sUp)
		s.print()
		fmt.Print("\nAlternate up: ")
		ss.Stack[ss.len()-1].print()
	} else {
		s.print()
		ss.Stack[ss.len()-1].vRight = 0
		ss.Stack[ss.len()-1].sRight = nil
		exploder.sUp.vRight = 0
		exploder.sUp.sRight = nil
		exploder.sUp.print()
		fmt.Print("\nExploder sUp: ")
		fmt.Println(exploder.sUp)
		s.print()
		fmt.Print("\nAlternate up: ")
		ss.Stack[ss.len()-1].print()
	}
	exploder = nil
	fmt.Println()
	exploder.print()
}

func (s *Snail) explode() (res bool) {
	var ss SnailStack
	// Find the first snail with depth 5
	if s.findDepth(5, &ss) {
		s.explode_new(ss)
		ss.print()
		res = true
	} else {
		res = false
	}
	return
}
