package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputPath = "./input.txt"

func readInput(input string) []string {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	codelist := make([]string, 0)
	for scanner.Scan() {
		code := scanner.Text()
		// the list of all codes
		codelist = append(codelist, code)
	}
	return codelist
}

func part1() {
	two, three := 0, 0
	codelist := readInput(inputPath)
	for _, code := range codelist {
		m := make(map[rune]int)

		for _, v := range code {
			m[v]++
		}
		hasTwo, hasThree := false, false
		for _, v := range m {
			if v == 2 {
				hasTwo = true
			}
			if v == 3 {
				hasThree = true
			}
		}
		if hasTwo {
			two++
		}
		if hasThree {
			three++
		}

	}
	fmt.Printf("%d doubles, %d triples.\nTotal: ", two, three)
	fmt.Println(two * three)
}

func part2() {
	codelist := readInput(inputPath)
	for index, code := range codelist { // for each code
		for i := index + 1; i < len(codelist); i++ { // look at every other code
			diffCount := 0
			for letter := 0; letter < len(codelist[0]) && diffCount < 2; letter++ {
				if code[letter] != codelist[i][letter] {
					diffCount++
				}
			}
			if diffCount == 1 {
				fmt.Printf("%s\n%s\n", code, codelist[i])
			}
		}
	}
}

func main() {
	fmt.Println("Part 1")
	part1()
	fmt.Println("Part 2")
	part2()
}
