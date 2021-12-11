package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text1 []string
	var text2 []string
	for scanner.Scan() {
		text1 = append(text1, scanner.Text())
		text2 = append(text2, scanner.Text())
	}
	file.Close()
	bitsize := 12

	emptyCount := len(text1)
	for i := 0; i < bitsize; i++ {
		var one = 0
		var zero = 0
		// Count up the bit positions
		for _, line := range text1 {
			if line != "EMPTY" {
				bit, _ := strconv.Atoi(string(line[i]))
				if bit == 1 {
					one++
				} else {
					zero++
				}
			}
		}

		if one >= zero {
			for x := 0; x < len(text1); x++ {
				line := text1[x]
				if line != "EMPTY" {
					bit, _ := strconv.Atoi(string(line[i]))
					if bit == 1 {
						// Oxygen generator rating candidate
					} else {
						text1[x] = "EMPTY"
						emptyCount--
					}
				}
			}
		} else {
			for x := 0; x < len(text1); x++ {
				line := text1[x]
				if line != "EMPTY" {
					bit, _ := strconv.Atoi(string(line[i]))
					if bit == 0 {
						// Oxygen generator rating candidate
					} else {
						text1[x] = "EMPTY"
						emptyCount--
					}
				}
			}
		}
		if emptyCount == 1 {
			for _, line := range text1 {
				if line != "EMPTY" {
					fmt.Printf("Oxygen generator rating is %s\n", line)
				}
			}
		}
	}

	emptyCount = len(text2)
	for i := 0; i < bitsize; i++ {
		var one = 0
		var zero = 0
		// Count up the bit positions
		for _, line := range text2 {
			if line != "EMPTY" {
				bit, _ := strconv.Atoi(string(line[i]))
				if bit == 1 {
					one++
				} else {
					zero++
				}
			}
		}

		if one < zero {
			for x := 0; x < len(text2); x++ {
				line := text2[x]
				if line != "EMPTY" {
					bit, _ := strconv.Atoi(string(line[i]))
					if bit == 1 {
						// CO2 scrubber
					} else {
						text2[x] = "EMPTY"
						emptyCount--
					}
				}
			}
		} else {
			for x := 0; x < len(text2); x++ {
				line := text2[x]
				if line != "EMPTY" {
					bit, _ := strconv.Atoi(string(line[i]))
					if bit == 0 {
						// CO2 scrubber
					} else {
						text2[x] = "EMPTY"
						emptyCount--
					}
				}
			}
		}
		if emptyCount == 1 {
			for _, line := range text2 {
				if line != "EMPTY" {
					fmt.Printf("CO2 scrubber rating is %s", line)
				}
			}
		}
	}
}
