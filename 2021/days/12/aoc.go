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

type Path struct {
	History []int
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

func (g *Graph) findEnd(path Path) {
	n := path.History[len(path.History)-1]
	path.Visited[n] = true
	if g.IDs[n] == "end" {
		fmt.Printf("Found end: ")
		fmt.Println(path.History)
	}
	fmt.Printf("%d edges from node %d: ", len(g.Edges[n]), n)
	fmt.Println(g.Edges[n])
	for _, node := range g.getNext(n) {
		if g.visited(node, path.Visited) {
			continue
		}
		fmt.Printf("Calling findEnd on node %d\n", node)
		path.History = append(path.History, node)
		g.findEnd(path)
	}
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
	history = append(history, 1)
	visited := make(map[int]bool)
	path := Path{history, visited}
	cave.findEnd(path)
	//	var answer [][]string

}

func main() {
	text := utils.ReadInput(0)
	cave := makeCave(text)
	part1(cave)
}
