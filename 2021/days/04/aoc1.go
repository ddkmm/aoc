package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func read() []string {
	//file, _ := os.Open("./tinput.txt")
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()
	return text
}

func generate(text []string) ([]string, [][]string) {
	var numbers []string
	var boards [][]string
	var board [][]string
	for i := 0; i < len(text); i++ {
		if i == 0 {
			numbers = strings.Split(text[i], ",")
		} else if text[i] != "" {
			board = append(board, strings.Fields(text[i]))
		} else if i != 1 {
			boards = append(boards, board...)
			board = nil
		}
	}
	boards = append(boards, board...)
	return numbers, boards
}

func play(numbers []string, boards [][]string) {
	// call a number and go through the boards, setting the number to 0.
	var win = false
	for i := 0; i < len(numbers); i++ {
		number := string(numbers[i])
		for boardOffset := 0; boardOffset < len(boards); boardOffset = boardOffset + 5 {
			// fmt.Printf("Checking board %d\n", boardOffset/5+1)
			for b := boardOffset; b < boardOffset+5; b++ {
				//				fmt.Printf("Checking row %d\n", b)
				for spot := 0; spot < 5; spot++ {
					// fmt.Printf("Checking %s\n", string(boards[b][spot]))
					if string(boards[b][spot]) == number {
						//	fmt.Printf("Found %s on board %d, row %d, column %d\n", number, boardOffset/5+1, b%5+1, spot+1)
						boards[b][spot] = string('X')
						// Check this board
						win = check(number, boardOffset, boards)
						if win {
							score(number, boardOffset, boards)
						}
					}
				}
			}
		}

	}
}

func check(number string, boardOffset int, boards [][]string) bool {
	// Does this board have a full row of Xs
	for row := 0; row < 5; row++ {
		count := 0
		for spot := 0; spot < 5; spot++ {
			if boards[row+boardOffset][spot] == string('X') {
				count++
			}
			if count == 5 {
				//fmt.Printf("Board %d is a winner\n", boardOffset)
				return true
			}
		}
	}
	// Does this board have a full column of Xs
	for spot := 0; spot < 5; spot++ {
		count := 0
		for row := 0; row < 5; row++ {
			if boards[row+boardOffset][spot] == string('X') {
				count++
			}
			if count == 5 {
				//fmt.Printf("Board %d is a winner\n", boardOffset)
				return true
			}
		}
	}
	return false
}

func score(number string, boardOffset int, boards [][]string) {
	total := 0
	for row := 0; row < 5; row++ {
		for spot := 0; spot < 5; spot++ {
			if boards[row+boardOffset][spot] != string('X') {
				val, _ := strconv.Atoi(boards[row+boardOffset][spot])
				total += val
			}
		}
	}

	num, _ := strconv.Atoi(number)
	fmt.Printf("%d\n", total*num)
	// And erase the board
	for row := 0; row < 5; row++ {
		for spot := 0; spot < 5; spot++ {
			boards[row+boardOffset][spot] = string('X')
		}
	}
}

func main() {

	text := read()
	numbers, boards := generate(text)
	fmt.Printf("%d moves, %d boards\n", len(numbers), len(boards)/5)
	play(numbers, boards)

}
