package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/dkim/aoc/2021/utils"
)

func createCrabs(text []string) map[int]int {
	crabs := make(map[int]int)
	for _, each := range text {
		line := strings.Split(each, ",")
		for _, num := range line {
			val, _ := strconv.Atoi(num)
			crabs[val]++
		}
	}
	return crabs

}

func summer(delta int) int {
	total := 0
	for i := 1; i <= int(math.Abs(float64(delta))); i++ {
		total += i
	}
	return total
}

func crabShuffle(crabs map[int]int) {
	defer utils.TimeTrack(time.Now(), "crabShuffle")
	meanVal := 0
	meanNumber := 0
	for key, element := range crabs {
		if element > meanVal {
			meanVal = element
			meanNumber = key
			fmt.Printf("Meanval: %d with %d instances\n", meanNumber, meanVal)
		}
		//fmt.Printf("%d:%d\n", key, element)
	}

	// Align crabs to mean value
	sum := 0
	for key, value := range crabs {
		delta := meanNumber - key
		delta = summer(delta)
		//fmt.Printf("Moving crab %d to %d. Uses %d fuel\n", key, meanNumber, int(delta))
		sum += delta * value
		//	fmt.Printf("%d:%d\n", key, meanNumber)
	}
	fmt.Printf("Mean value solution: %d\n", int(sum))

	// brute force
	bestSum := sum
	for pass := 0; pass < len(crabs); pass++ {
		sum = 0
		for key, value := range crabs {
			delta := pass - key
			delta = summer(delta)
			sum += delta * value
			//	fmt.Printf("%d:%d\n", key, meanNumber)
		}
		if sum < bestSum {
			bestSum = sum
		}
	}
	fmt.Printf("Brute force solution: %d\n", int(bestSum))

}

func main() {

	text := utils.ReadInput(1)
	crabs := createCrabs(text)
	crabShuffle(crabs)

}
