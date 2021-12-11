package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"

	"github.com/dkim/aoc/2021/utils"
)

// Make a map which contains matching pairs
// Make a stack
// push each left onto the stack
// when we get a right symbol pop and see if it's a match
// ignore the pop if it's another left symbol

func part1(text []string) []string {
	defer utils.TimeTrack(time.Now(), "Part 1")
	var stack []string
	var incomplete []string
	sum := 0
	match := map[string]string{
		"[": "]", "]": "[",
		"{": "}", "}": "{",
		"<": ">", ">": "<",
		"(": ")", ")": "(",
	}
	left := "([{<"
	for _, line := range text {
		corrupt := false
		for _, a := range line {
			next := string(a)
			if strings.Contains(left, next) {
				stack = append(stack, next)
			} else if stack[len(stack)-1] == match[next] {
				stack = stack[:len(stack)-1]
			} else {
				switch next {
				case ")":
					fmt.Printf("Found %s, adding %d\n", next, 3)
					sum += 3
				case "]":
					fmt.Printf("Found %s, adding %d\n", next, 57)
					sum += 57
				case "}":
					fmt.Printf("Found %s, adding %d\n", next, 1197)
					sum += 1197
				case ">":
					fmt.Printf("Found %s, adding %d\n", next, 25137)
					sum += 25137
				}
				corrupt = true
				break
			}
		}
		if !corrupt {
			incomplete = append(incomplete, line)
		}
	}
	fmt.Println(sum)
	return incomplete
}

func part2(incomplete []string) {
	defer utils.TimeTrack(time.Now(), "Part 2")
	var stack []string
	var scores []int
	score := 0
	match := map[string]string{
		"[": "]", "]": "[",
		"{": "}", "}": "{",
		"<": ">", ">": "<",
		"(": ")", ")": "(",
	}
	left := "([{<"
	for _, line := range incomplete {
		for _, a := range line {
			next := string(a)
			if strings.Contains(left, next) {
				stack = append(stack, next)
			} else if stack[len(stack)-1] == match[next] {
				stack = stack[:len(stack)-1]
			}
		}
		// Now the stack only contains leftovers
		for i := len(stack) - 1; i >= 0; i-- {
			switch match[stack[i]] {
			case ")":
				score *= 5
				score += 1
			case "]":
				score *= 5
				score += 2
			case "}":
				score *= 5
				score += 3
			case ">":
				score *= 5
				score += 4
			}
		}
		scores = append(scores, score)
		score = 0
		stack = nil
	}
	sort.Ints(scores)
	fmt.Println(scores)
	fmt.Println(scores[int(math.Floor(float64(len(scores))/2))])
}

func main() {
	text := utils.ReadInput(1)
	fmt.Println(len(text))
	part2(part1(text))

}
