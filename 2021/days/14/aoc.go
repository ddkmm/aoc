package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/dkim/aoc/2021/utils"
)

type Node struct {
	Data string
	Next *Node
}

type LinkedList struct {
	Length int
	Head   *Node
	Tail   *Node
}

type InsertPair struct {
	Data        string
	InsertPoint *Node
}

func (list *LinkedList) initList(a *Node) {
	list.Head = a
	list.Tail = a
	list.Length = 1
}

func (list *LinkedList) insertAfter(a *Node, b *Node) {
	tempNode := a.Next
	a.Next = b
	b.Next = tempNode
	list.Length++
}

func (list *LinkedList) insertAtEnd(a *Node) {
	list.Tail.Next = a
	list.Tail = a
	list.Length++
}

func (list *LinkedList) printList() {
	fmt.Printf("Size %d\n", list.Length)
	node := list.Head
	for node != list.Tail {
		fmt.Printf("%s", node.Data)
		node = node.Next
	}
	fmt.Printf("%s\n", node.Data)

}

func processInput(text []string) (map[string]string, string) {
	var start string
	rules := make(map[string]string)
	for _, line := range text {
		if !strings.Contains(line, "->") && len(line) != 0 {
			start = line
		} else if len(line) != 0 {
			line := strings.Fields(line)
			rules[line[0]] = line[2]
		}
	}
	return rules, start
}

func part1(text []string, steps int) {
	defer utils.TimeTrack(time.Now(), "Part 1")
	rules, start := processInput(text)
	chain := []rune(start)
	var polymer LinkedList

	// Make the template
	for i := 0; i < len(chain); i++ {
		node := Node{string(chain[i]), nil}
		if i == 0 {
			polymer.initList(&node)
		} else if i != len(chain) {
			polymer.insertAtEnd(&node)
		}
	}
	polymer.printList()

	// grow for n steps
	for n := 0; n < steps; n++ {
		fmt.Printf("After step %d\n", n+1)
		// Identify what the value of the Node will be and where we should insert it after
		var newChain []InsertPair
		node := polymer.Head
		for node != polymer.Tail {
			text := string(node.Data) + string(node.Next.Data)
			val := rules[string(text)]
			newChain = append(newChain, InsertPair{val, node})
			node = node.Next
		}
		// Do the insertions
		for _, ip := range newChain {
			polymer.insertAfter(ip.InsertPoint, &Node{ip.Data, nil})
		}
		newChain = nil
		polymer.printList()
	}

	// Calculate the most common and least common
	score := make(map[string]int)
	walker := polymer.Head
	for walker != polymer.Tail {
		score[string(walker.Data)]++
		walker = walker.Next
	}
	// and the tail
	score[string(walker.Data)]++

	// Now find the min and max values
	min := score[string(walker.Data)]
	max := 0
	for _, value := range score {
		if value < min {
			min = value
		} else if value > max {
			max = value
		}
	}
	fmt.Printf("%d - %d = %d", max, min, max-min)
}

type Pair struct {
	L string
	R string
}

func part2(text []string, steps int) {
	// order doesn't matter, only pairs
	// the template is the seed and we break it up into
	// 3 pairs. each pair uses the rule and makes n - 1 more pairs
	// where n is the original length before applying rules
	defer utils.TimeTrack(time.Now(), "Part 2")
	rules, start := processInput(text)
	poly := make(map[Pair]int)
	chain := []rune(start)
	for i := 0; i < len(chain)-1; i++ {
		pair := Pair{string(chain[i]), string(chain[i+1])}
		poly[pair]++
	}

	for n := 0; n < steps; n++ {
		newPoly := make(map[Pair]int)
		for pair, num := range poly {
			text := string(pair.L) + string(pair.R)
			val := rules[string(text)]
			newPair := Pair{pair.L, val}
			newPoly[newPair] += num
			newPair = Pair{val, pair.R}
			newPoly[newPair] += num
		}
		// Add the new pairs back in
		poly = newPoly
		/*
			for key, value := range newPoly {
				poly[key] += value
			}
		*/
		newPoly = nil
	}
	// Count up letters
	score := make(map[string]int)
	for key, value := range poly {
		score[key.L] += value
	}
	// And finally add one for the end of the template
	score[string(chain[len(chain)-1])]++

	// Now find the min and max values
	min := score["N"]
	max := 0
	for _, value := range score {
		if value < min {
			min = value
		} else if value > max {
			max = value
		}
	}
	fmt.Printf("%d - %d = %d", max, min, max-min)
}

func main() {
	text := utils.ReadInput(1)
	part1(text, 10)
	part2(text, 40)
}
