package main

import (
	"fmt"
	"time"

	"github.com/dkim/aoc/2021/utils"
)

type Velocity struct {
	xVel int
	yVel int
}

type Probe struct {
	xPos int
	yPos int
	xVel int
	yVel int
}

type Target struct {
	minX int
	maxX int
	minY int
	maxY int
}

func (p Probe) inTarget(target Target) bool {
	if (p.xPos >= target.minX && p.xPos <= target.maxX) &&
		(p.yPos >= target.minY && p.yPos <= target.maxY) {
		return true
	}
	return false
}

func (p Probe) step(target Target) (int, bool) {
	maxY := p.yPos
	result := false
	for i := 1; i < 1000; i++ {
		// Change position
		p.xPos += p.xVel
		p.yPos += p.yVel
		// Change velocity
		if p.xVel > 0 {
			p.xVel--
		} else if p.xVel < 0 {
			p.xVel++
		}
		p.yVel--
		if p.yPos > maxY {
			maxY = p.yPos
		}

		// Check if in target
		result = p.inTarget(target)
		if result {
			//fmt.Printf("Step %d: (%d, %d) ", i, p.xPos, p.yPos)
			//fmt.Println(p.inTarget(target))
			break

		}
	}
	return maxY, result
}

func part1(target Target) {
	defer utils.TimeTrack(time.Now(), "Part 1")
	max := -100
	velocities := make(map[Velocity]int)
	for x := 0; x < 1000; x++ {
		for y := -1000; y < 1000; y++ {
			probe := Probe{0, 0, x, y}
			maxY, result := probe.step(target)
			if result {
				velocities[Velocity{x, y}]++
				//	fmt.Printf("Initial velocity %d, %d: gives %d max Y\n", x, y, maxY)
				if maxY > max {
					max = maxY
				}
			}
		}
	}
	fmt.Printf("Part 1: %d\n", max)
	fmt.Printf("Part 2: %d\n", len(velocities))
}

func main() {
	target := Target{253, 280, -73, -46}
	part1(target)
}
