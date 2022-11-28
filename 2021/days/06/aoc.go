package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dkim/aoc/2021/utils"
)

// Process input into lish list
func createFish(text []string) [9]int {
	fish := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, each := range text {
		line := strings.Split(each, ",")
		for _, num := range line {
			val, _ := strconv.Atoi(num)
			fish[val]++
		}
	}
	return fish
}

// Spawn
func spawn(fish [9]int, days int) {
	defer utils.TimeTrack(time.Now(), "Fish")
	for d := 0; d < days; d++ {
		temp := fish[0]
		fish[0] = fish[1]
		fish[1] = fish[2]
		fish[2] = fish[3]
		fish[3] = fish[4]
		fish[4] = fish[5]
		fish[5] = fish[6]
		fish[6] = fish[7] + temp
		fish[7] = fish[8]
		fish[8] = temp
	}
	total := 0
	for i := 0; i < len(fish); i++ {
		total += fish[i]
	}
	fmt.Printf("%d\n", total)
}

func main() {

	text := utils.ReadInput(1)
	fish := createFish(text)
	spawn(fish, 256)

}
