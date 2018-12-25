package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var inputPath = "./input_test.txt"

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

func main() {
	polymer := processInput(inputPath)
	goAgain := true
	fmt.Println(len(polymer))
	for goAgain {
		start := true
		var tempRunes = make([]rune, 1)
		goAgain = false
		var removed []rune
		for k := 0; k < len(polymer); k++ {
			v := rune(polymer[k])
			if k != len(polymer)-1 {
				if unicode.IsLower(v) {
					if v == unicode.ToUpper(rune(polymer[k+1])) {
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
				} else {
					if v == unicode.ToUpper(rune(polymer[k+1])) {
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
		fmt.Println(polymer)
		fmt.Printf("Removed ")
		fmt.Printf("%q\n", removed)
		fmt.Printf("New polymer length: ")
		fmt.Println(len(polymer))
	}

}
