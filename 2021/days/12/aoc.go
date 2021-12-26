package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/dkim/aoc/2021/utils"
)

type Graph struct {
	Edges  [][]int
	Labels map[string]int
	IDs    map[int]string
}

func newGraph(n int) *Graph {
	return &Graph{
		Edges:  make([][]int, n),
		Labels: make(map[string]int),
		IDs:    make(map[int]string),
	}
}

func (g *Graph) printNodes() {
	for i := 1; i <= len(g.IDs); i++ {
		fmt.Printf("ID %d: %s\n", i, g.IDs[i])
	}
}

func (g *Graph) printEdges() {
	for i := 1; i <= len(g.Labels); i++ {
		fmt.Printf("Node %s%d to ", g.IDs[i], i)
		for _, v := range g.Edges[i] {
			fmt.Printf("%s%d ", g.IDs[v], v)
		}
		fmt.Println()
	}
}

func (g *Graph) addEdge(u, v int) {
	g.Edges[u] = append(g.Edges[u], v)
	g.Edges[v] = append(g.Edges[v], u)
}

func (g *Graph) findPaths(n int, path []int) [][]int {
	if g.IDs[n] == "end" {
		path = append(path, n)
		return [][]int{path}
	}
	// Check if we've been to lower case nodes before
	if strings.ToUpper(g.IDs[n]) != g.IDs[n] {
		for _, c := range path {
			if n == c {
				return [][]int{}
			}
		}
	}
	paths := [][]int{}
	path = append(path, n)
	for _, c := range g.Edges[n] {
		newPath := append([]int{}, path...)
		paths = append(paths, g.findPaths(c, newPath)...)
	}
	return paths
}

func (g *Graph) findPaths2(n int, path []int, revisit bool) [][]int {
	if g.IDs[n] == "end" {
		path = append(path, n)
		return [][]int{path}
	}
	// Check if we've been to lower case nodes before
	if strings.ToUpper(g.IDs[n]) != g.IDs[n] {
		for _, c := range path {
			if n == c {
				if g.Labels["start"] == n {
					return [][]int{}
				} else if revisit {
					return [][]int{}
				} else {
					revisit = true
				}
			}
		}
	}
	paths := [][]int{}
	path = append(path, n)
	for _, c := range g.Edges[n] {
		newPath := append([]int{}, path...)
		paths = append(paths, g.findPaths2(c, newPath, revisit)...)
	}
	return paths
}

func makeCave(text []string) *Graph {
	n := len(text)
	cave := newGraph(n)
	i := 1
	for _, line := range text {
		nodes := strings.Split(line, "-")
		for _, n := range nodes {
			if cave.Labels[n] == 0 {
				cave.Labels[n] = i
				cave.IDs[i] = n
				i++
			}
		}
	}
	//	fmt.Println(cave.Labels)
	for _, line := range text {
		nodes := strings.Split(line, "-")
		//		fmt.Printf("Adding %s to cave system\n", nodes)
		cave.addEdge(cave.Labels[nodes[0]], cave.Labels[nodes[1]])
	}

	/*
		fmt.Println(cave.Edges)
		cave.printNodes()
		cave.printEdges()
	*/
	return cave
}

func part1(cave *Graph) {
	defer utils.TimeTrack(time.Now(), "Part 1")
	var history []int
	paths := cave.findPaths(cave.Labels["start"], history)
	fmt.Printf("%d\n", len(paths))
}

func part2(cave *Graph) {
	defer utils.TimeTrack(time.Now(), "Part 2")
	var history []int
	paths := cave.findPaths2(cave.Labels["start"], history, false)
	fmt.Printf("%d\n", len(paths))
}

func main() {
	text := utils.ReadInput(1)
	cave := makeCave(text)
	part1(cave)
	cave2 := makeCave(text)
	part2(cave2)
}
