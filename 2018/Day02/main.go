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
	max := len(codelist)
	for index, code := range codelist { // for each code
		for i := index + 1; i < len(codelist); i++ { // look at every other code
			for letter := 0; letter < len(codelist[0]); letter++ {

			}
		}
	}
	for i := 0; i < len(codelist[0]); i++ {

		freq := make(map[rune]int)
		for _, code := range codelist {
			freq[rune(code[i])]++
			fmt.Println(freq[rune(code[i])])
			if freq[rune(code[i])] == max {
				fmt.Println(rune(code[i]))
			}
		}
	}
}

func main() {
	//part1()
	part2()
}
