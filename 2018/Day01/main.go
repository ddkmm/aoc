package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	sum := 0
	doScan := true
	m := make(map[int]int)
	var freqs []int
	m[sum]++
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		freq, _ := strconv.Atoi(line[1:])
		if line[0] == '-' {
			freq = -1 * freq
		}
		freqs = append(freqs, freq)
	}
	file.Close()
	fmt.Println("File read")

	for doScan == true {
		fmt.Println("Run loop")
		for _, v := range freqs {
			sum += v
			m[sum]++
			if m[sum] > 1 {
				doScan = false
				fmt.Println(sum)
				break
			}
		}
	}
	fmt.Println("Done")
}
