package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var inputPath = "./input.txt"

func processInput(inputPath string) string {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}
	return line
}

func reactPolymer(polymer string) {
	goAgain := true
	fmt.Printf("Starting length: %d ", len(polymer))
	for goAgain {
		start := true
		var tempRunes = make([]rune, 1)
		goAgain = false
		var removed []rune
		for k := 0; k < len(polymer); k++ {
			v := rune(polymer[k])
			if k != len(polymer)-1 {
				if unicode.IsLower(v) {
					if v == unicode.ToLower(rune(polymer[k+1])) && unicode.IsUpper(rune(polymer[k+1])) {
						goAgain = true
						removed = append(removed, v)
						removed = append(removed, rune(polymer[k+1]))
						k++
					} else if start {
						tempRunes[0] = v
						start = false
					} else {
						tempRunes = append(tempRunes, v)
					}
				} else if unicode.IsUpper(v) {
					if v == unicode.ToUpper(rune(polymer[k+1])) && unicode.IsLower(rune(polymer[k+1])) {
						goAgain = true
						removed = append(removed, v)
						removed = append(removed, rune(polymer[k+1]))
						k++
					} else if start {
						tempRunes[0] = v
						start = false
					} else {
						tempRunes = append(tempRunes, v)
					}
				}
			} else {
				tempRunes = append(tempRunes, v)
			}
		}
		polymer = string(tempRunes)
		//	fmt.Println(polymer)
		//	fmt.Printf("Removed ")
		//	fmt.Printf("%q\n", removed)
		//	fmt.Printf("New polymer length: ")
		//	fmt.Println(len(polymer))
	}
	fmt.Printf("Final length %d\n", len(polymer))
}

func stripLetter(polymer string, letter string) string {
	fmt.Println(letter)

	var tempRunes = make([]rune, 1)
	for k, v := range polymer {
		if !strings.EqualFold(string(v), letter) {
			if k == 0 {
				tempRunes[k] = v
			} else {
				tempRunes = append(tempRunes, v)
			}
		}
	}
	return string(tempRunes)
}

func generateAlpha() string {
	p := make([]byte, 26)
	for i := range p {
		p[i] = 'a' + byte(i)
	}
	return string(p)
}

func main() {
	polymer := processInput(inputPath)
	reactPolymer(polymer)
	alphabet := generateAlpha()
	for _, v := range alphabet {
		reactPolymer(stripLetter(polymer, string(v)))
	}
}
