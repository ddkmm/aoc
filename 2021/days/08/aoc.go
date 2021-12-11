package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/dkim/aoc/2021/utils"
)

func part1(text []string) {
	defer utils.TimeTrack(time.Now(), "Part 1")
	total := 0
	for _, each := range text {
		line := strings.Split(each, "|")
		display := strings.Fields(line[1])
		for _, d := range display {
			if len(d) == 3 ||
				len(d) == 2 ||
				len(d) == 4 ||
				len(d) == 7 {
				total++
			}
		}
	}
	fmt.Println(total)
}

//   6:      2:      5:      5:      4:
//  AAAA    ....    AAAA    AAAA    ....
// B    C  .    C  .    C  .    C  B    C
// B    C  .    C  .    C  .    C  B    C
//  ....    ....    DDDD    DDDD    DDDD
// E    F  .    F  E    .  .    F  .    F
// E    F  .    F  E    .  .    F  .    F
//  GGGG    ....    GGGG    GGGG    ....
//
//   5:      6:      3:      7:      6:
//  AAAA    AAAA    AAAA    AAAA    AAAA
// B    .  B    .  .    C  B    C  B    C
// B    .  B    .  .    C  B    C  B    C
//  DDDD    DDDD    ....    DDDD    DDDD
// .    F  E    F  .    F  E    F  .    F
// .    F  E    F  .    F  E    F  .    F
//  GGGG    GGGG    ....    GGGG    GGGG

/*
# lines	A B C D E F G
1, 2   	. . ? . . ? .
2, 5	? . ? ? ? . ?
3, 5	? . ? ? . ? ?
4, 4	. ? ? ? . ? .
5, 5	? ? . ? . ? ?
6, 6	? ? . ? ? ? ?
7, 3	? . ? . . ? .
8, 7	? ? ? ? ? ? ?
9, 6	? ? ? ? . ? ?
0, 6	? ? ? . ? ? ?

*/
/*
-0, 2
-1, 3
-2, 4
3, 5
4, 5
5, 5
6, 6
7, 6
8, 6
-9, 7
*/

func countCommon(a string, b string) (count int) {
	for _, c := range a {
		if strings.Contains(b, string(c)) {
			count++

		}

	}
	return
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func decode(text string) int {
	sum := 0
	code := make(map[string]int)
	input := strings.Fields(strings.Replace(text, "|", "", -1))
	output := input[len(input)-4:]
	input = input[0 : len(input)-4]
	sort.Slice(input, func(i, j int) bool {
		return len(input[i]) < len(input[j])
	})
	code[SortString(input[0])] = 1
	code[SortString(input[1])] = 7
	code[SortString(input[2])] = 4
	code[SortString(input[9])] = 8

	// 5 segment
	for i := 0; i < 3; i++ {
		if countCommon(SortString(input[3+i]), SortString(input[0])) == 2 {
			code[SortString(input[3+i])] = 3
		} else if countCommon(SortString(input[3+i]), SortString(input[2])) == 2 {
			code[SortString(input[3+i])] = 2
		} else if countCommon(SortString(input[3+1]), SortString(input[0])) == 1 &&
			countCommon(SortString(input[3+i]), SortString(input[2])) == 3 {
			code[SortString(input[3+i])] = 5
		} else {
			code[SortString(input[3+i])] = 5
		}
	}

	for i := 0; i < 3; i++ {
		if countCommon(SortString(input[6+i]), SortString(input[0])) == 1 {
			code[SortString(input[6+i])] = 6
		} else if countCommon(SortString(input[6+i]), SortString(input[2])) == 4 {
			code[SortString(input[6+i])] = 9
		} else if countCommon(SortString(input[6+i]), SortString(input[0])) == 2 &&
			countCommon(SortString(input[6+i]), SortString(input[2])) == 3 {
			code[SortString(input[6+i])] = 0
		}
	}
	sum += code[SortString(output[0])] * 1000
	sum += code[SortString(output[1])] * 100
	sum += code[SortString(output[2])] * 10
	sum += code[SortString(output[3])]
	return sum

}
func part2(text []string) {
	defer utils.TimeTrack(time.Now(), "Part 2")
	total := 0
	for _, line := range text {
		total += decode(line)
	}
	fmt.Println(total)
}

func main() {

	text := utils.ReadInput(1)
	part1(text)
	part2(text)

}
