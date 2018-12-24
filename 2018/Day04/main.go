package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var inputPath = "./input.txt"

type guard struct {
	date          string
	guardID       string
	guardActivity [60]string
	sleepTime     int
}

func processInput(inputPath string) []string {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var logEntries []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		logEntries = append(logEntries, line)
	}
	sort.Strings(logEntries)
	return logEntries
}

func main() {
	logEntries := processInput(inputPath)
	var guardLog []guard

	for k, v := range logEntries {
		var date string
		var year, hour, minute int
		entry := strings.Split(v, "]")
		fmt.Sscanf(entry[0], "[%d-%s %d:%d", &year, &date, &hour, &minute)
		event := strings.Fields(strings.TrimSpace(entry[1]))
		// process the activity statement

		if event[0] == "Guard" {
			recordActivity := true
			i := 1
			sleeping := 0
			g := guard{}
			g.date = date
			g.guardID = event[1]
			startTime := 0
			for recordActivity && k+i < len(logEntries) {
				var date2 string
				var year2, hour2, minute2 int
				log2 := strings.Split(logEntries[k+i], "]")
				fmt.Sscanf(log2[0], "[%d-%s %d:%d", &year2, &date2, &hour2, &minute2)
				event2 := strings.Fields(strings.TrimSpace(log2[1]))

				if event2[0] == "Guard" {
					recordActivity = false
				} else { // Activity for the given guard
					if event2[0] == "falls" { // awake until the fall asleep
						for j := startTime; j < minute2 && j < 60; j++ {
							g.guardActivity[j] = " "
						}
						startTime = minute2
					} else if event2[0] == "wakes" { // asleep until they wake up
						for j := startTime; j < minute2 && j < 60; j++ {
							g.guardActivity[j] = "#"
							sleeping++
						}
						startTime = minute2
					}
				}
				i++
				g.date = date2
			}
			g.sleepTime = sleeping
			guardLog = append(guardLog, g)
		}
	}

	fmt.Println("Date   ID     Minute")
	fmt.Println("              000000000011111111112222222222333333333344444444445555555555")
	fmt.Println("              012345678901234567890123456789012345678901234567890123456789")
	for _, v := range guardLog {
		fmt.Printf("%s  %-5v  ", v.date, v.guardID)
		for _, c := range v.guardActivity {
			fmt.Printf("%s", c)
		}
		fmt.Printf("\n")
	}
	var sleepHistogram = make(map[string][60]int)
	for _, v := range guardLog {
		for i := 0; i < 60; i++ {
			if v.guardActivity[i] == "#" {
				temp := sleepHistogram[v.guardID]
				temp[i]++
				sleepHistogram[v.guardID] = temp
			}
		}
	}
	var total = make(map[string]int)
	for _, v := range guardLog {
		total[v.guardID] += v.sleepTime
	}
	var sleepChampion string
	sleepTotal := 0
	for k, v := range total {
		if v > sleepTotal {
			sleepTotal = v
			sleepChampion = k
		}
	}
	fmt.Printf("Guard %s is the sleep champion and ", sleepChampion)

	// Now we have the guard who sleeps the most, sleepChampion
	// make a histogram of all the minutes he sleeps to identify
	// which one is the likely minute
	var histo [60]int
	for _, v := range guardLog {
		if v.guardID == sleepChampion {
			for index, c := range v.guardActivity {
				if c == "#" {
					histo[index]++
				}
			}
		}
	}

	sleepTotal = 0
	sleepMinute := 0
	for k, v := range histo {
		if v > sleepTotal {
			sleepTotal = v
			sleepMinute = k
		}
	}
	fmt.Printf("minute %d is the sleepiest with %d\n", sleepMinute, sleepTotal)
	id, _ := strconv.Atoi(strings.SplitAfter(sleepChampion, "#")[1])
	fmt.Printf("%d ID X %d minute = %d\n", id, sleepMinute, id*sleepMinute)
	var part2 = make(map[string]int)
	//fmt.Println(guardLog)
	for _, v := range guardLog {
		if v.guardActivity[sleepMinute] == "#" {
			part2[v.guardID]++
		}
	}
	var newGuard string
	limit := 0
	for k, v := range sleepHistogram {
		for i := 0; i < 60; i++ {
			if v[i] > limit {
				sleepMinute = i
				limit = v[i]
				newGuard = k
			}
		}
	}

	fmt.Printf("%s : %d on minute %d\n", newGuard, limit, sleepMinute)
	id, _ = strconv.Atoi(strings.SplitAfter(newGuard, "#")[1])
	fmt.Printf("%d ID X %d minute = %d\n", id, sleepMinute, id*sleepMinute)

}
