package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputPath = "./input.txt"

// Claim contains an elf's claim
type Claim struct {
	coord Coordinates
	dimen Dimensions
}

// Coordinates describe the upper left hand position of the claim
type Coordinates struct {
	X int
	Y int
}

// Dimensions are the dimensions of the claim
type Dimensions struct {
	X int
	Y int
}

func main() {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var claims []Claim
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s1 := strings.Split(line, "@")
		s2 := strings.Split(s1[1], ":")
		coordinates := s2[0]
		dimensions := s2[1]
		s3 := strings.Split(coordinates, ",")
		x, _ := strconv.Atoi(strings.TrimSpace(s3[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(s3[1]))
		coor := Coordinates{x, y}
		s4 := strings.Split(dimensions, "x")
		x, _ = strconv.Atoi(strings.TrimSpace(s4[0]))
		y, _ = strconv.Atoi(strings.TrimSpace(s4[1]))
		dim := Dimensions{x, y}
		claim := Claim{coor, dim}
		claims = append(claims, claim)
	}
	// Now plot the fabric values
	total := 0
	var fabric [1001][1001]int
	// Stake the claims
	for _, v := range claims {
		claimX := v.coord.X
		claimY := v.coord.Y
		for i := 0; i < v.dimen.Y; i++ {
			for j := 0; j < v.dimen.X; j++ {
				fabric[claimX+j][claimY+i]++
			}
		}
	}

	for _, v := range fabric {
		for _, row := range v {
			if row > 1 {
				total++
			}
		}
	}
	fmt.Printf("%d overlapping square inches.\n", total)

	// read the claims
	for k, v := range claims {
		overlap := false
		claimX := v.coord.X
		claimY := v.coord.Y
		for i := 0; i < v.dimen.Y; i++ {
			for j := 0; j < v.dimen.X; j++ {
				if fabric[claimX+j][claimY+i] > 1 {
					overlap = true
				}
			}
		}
		if !overlap {
			fmt.Printf("Claim %d has no overlap\n", k+1)
		}
	}

}
