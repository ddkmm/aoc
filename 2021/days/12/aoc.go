package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/dkim/aoc/2021/utils"
)

type Graph struct {
	Edges   [][]int
	Labels  map[string]int
	IDs     map[int]string
	Visited map[int]bool
}

func newGraph(n int) *Graph {
	return &Graph{
		Edges:   make([][]int, n),
		Labels:  make(map[string]int),
		IDs:     make(map[int]string),
		Visited: make(map[int]bool),
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

func (g *Graph) getNext(n int) []int {
	return g.Edges[n]
}

func (g *Graph) visited(n int, visited map[int]bool) bool {
	if strings.ToUpper(g.IDs[n]) != g.IDs[n] {
		return visited[n]
	} else {
		return false
	}
}

func (g *Graph) findEnd(visited map[int]bool, history []int) [][]int {
	n := history[len(history)-1]
	visited[n] = true
	if g.IDs[n] == "end" {
		fmt.Printf("Found end: ")
		history = append(history, n)
		return [][]int{history}
	}
	fmt.Printf("%d edges from node %d: ", len(g.Edges[n]), n)
	fmt.Println(g.Edges[n])
	paths := [][]int{}
	for _, node := range g.getNext(n) {
		if g.visited(node, visited) {
			continue
		}
		fmt.Printf("Calling findEnd on node %d\n", node)
		history = append(history, node)
		newPath := append([]int{}, history...)
		paths = append(paths, g.findEnd(visited, newPath)...)
	}
	return paths
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
	fmt.Println(cave.Labels)
	for _, line := range text {
		nodes := strings.Split(line, "-")
		fmt.Printf("Adding %s to cave system\n", nodes)
		cave.addEdge(cave.Labels[nodes[0]], cave.Labels[nodes[1]])
	}

	fmt.Println(cave.Edges)
	cave.printNodes()
	cave.printEdges()
	return cave
}

func part1(cave *Graph) {
	defer utils.TimeTrack(time.Now(), "Part 1")
	var history []int
	paths := cave.findPaths(cave.Labels["start"], history)
	fmt.Println(paths)
	fmt.Printf("%d", len(paths))
}

func main() {
	text := utils.ReadInput(0)
	cave := makeCave(text)
	part1(cave)
}
