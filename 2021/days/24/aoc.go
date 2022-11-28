package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dkim/aoc/2021/utils"
)

type ALU struct {
	vars  map[string]int
	stack []string
}

func (a *ALU) inp(varName string, input int) {
	a.vars[varName] = input
}

func (a *ALU) mul(dest string, multiplier int, val int) {
	a.vars[dest] = multiplier * val
}

func (a *ALU) print() {
	fmt.Println(a.vars)
}

func part1(text []string, inputs map[string][9]int) {
	defer utils.TimeTrack(time.Now(), "Part 1")
	vars := make(map[string]int)
	var stack []string
	alu := ALU{vars, stack}

	for i := 0; i < len(inputs); i++ {
		for n := 1; n < 10; n++ {
			for ln, line := range text {
				prog := strings.Fields(line)
				switch prog[0] {
				case "inp":
					fmt.Printf("Line %d: inp %s with %d\n", ln, prog[1], n)
					alu.inp(prog[1], n)
				case "mul":
					fmt.Printf("Line %d: mul %s %s with %d\n", ln, prog[1], prog[2], n)
					multiplier, _ := strconv.Atoi(prog[2])
					alu.mul(prog[1], multiplier, n)
				case "eql":
					//					fmt.Printf("Line %d: eql %s %s with %d\n", ln, prog[1],
				}

			}
			alu.print()
		}
	}
}

func processInput(text []string) map[string][9]int {
	inputs := make(map[string][9]int)
	digits := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, line := range text {
		if strings.Contains(line, "inp") {
			inputs[strings.Fields(line)[1]] = digits
		}
	}

	fmt.Printf("%d digit input required\n", len(inputs))
	return inputs
}

func main() {
	text := utils.ReadInput(0)
	n := processInput(text)
	part1(text, n)
}
