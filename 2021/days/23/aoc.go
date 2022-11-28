package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/dkim/aoc/2021/utils"
)

type Piece struct {
	pos     Coord
	stepVal int
	steps   int
}

type Game struct {
	board []string
	A1    Piece
	B1    Piece
	C1    Piece
	D1    Piece
	A2    Piece
	B2    Piece
	C2    Piece
	D2    Piece
}

type Coord struct {
	y int
	x int
}

func (g *Game) print() {
	for i := 0; i < len(g.board); i++ {
		fmt.Println(g.board[i])
	}
}

func scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func (g *Game) play() {
	g.print()
	fmt.Printf("Enter Piece to move: ")
	reader := bufio.NewReader(os.Stdin)
	char, _ := reader.ReadString('\n')
	switch char {
	case "A1":
		fmt.Println(g.A1)

	}

}

func part1(text []string) {
	defer utils.TimeTrack(time.Now(), "Part 1")
	var A1, B1, C1, D1 Piece
	for i, letter := range text[2] {
		if letter == 'A' {
			A1 = Piece{Coord{2, i}, 1, 0}
		} else if letter == 'B' {
			B1 = Piece{Coord{2, i}, 10, 0}
		} else if letter == 'C' {
			C1 = Piece{Coord{2, i}, 100, 0}
		} else if letter == 'D' {
			D1 = Piece{Coord{2, i}, 1000, 0}
		}
	}
	var A2, B2, C2, D2 Piece
	for i, letter := range text[3] {
		if letter == 'A' {
			A2 = Piece{Coord{3, i}, 1, 0}
		} else if letter == 'B' {
			B2 = Piece{Coord{3, i}, 10, 0}
		} else if letter == 'C' {
			C2 = Piece{Coord{3, i}, 100, 0}
		} else if letter == 'D' {
			D2 = Piece{Coord{3, i}, 1000, 0}
		}
	}
	g := Game{text, A1, B1, C1, D1, A2, B2, C2, D2}
	g.play()
}

func main() {
	text := utils.ReadInput(0)
	part1(text)
}
