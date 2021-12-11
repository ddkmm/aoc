package utils

import (
	"bufio"
	"log"
	"os"
	"time"
)

func ReadInput(flag int) []string {
	var file *os.File
	if flag == 0 {
		file, _ = os.Open("./tinput.txt")
	} else {
		file, _ = os.Open("./input.txt")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()
	return text
}

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
