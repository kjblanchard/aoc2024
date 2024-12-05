package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var (
	rules                    [][]int
	pageUpdates              [][]int
	incorrectUpdatesUnsorted [][]int
	total                    int
	part1                    = false
)

func checkError(e error) {
	if e != nil {
		log.Panicf("Failed for some reason, %s", e)
	}
}

func getRule(line string) {
	var first, second int
	_, err := fmt.Sscanf(line, "%d|%d", &first, &second)
	checkError(err)
	rules = append(rules, []int{first, second})
}

func getPageUpdate(line string) {
	split := strings.Split(line, ",")
	var pageUpdate []int
	for _, num := range split {
		numInt, err := strconv.Atoi(num)
		checkError(err)
		pageUpdate = append(pageUpdate, numInt)
	}
	pageUpdates = append(pageUpdates, pageUpdate)
}

func checkUpdates() {
	for _, update := range pageUpdates {
		goodUpdate := true
		for _, rule := range rules {
			first := slices.Index(update, rule[0])
			second := slices.Index(update, rule[1])
			if first == -1 || second == -1 || second > first {
				continue
			}
			incorrectUpdatesUnsorted = append(incorrectUpdatesUnsorted, update)
			goodUpdate = false
			break
		}
		if part1 && goodUpdate {
			updateLength := len(update)
			if updateLength%2 == 0 {
				log.Panicf("Somehow it is an even number?: %d", updateLength)
			}
			total += update[updateLength/2]
		}
	}
}

func sortIncorrectUpdate(update []int) {
	var matchedRules [][]int
	for _, rule := range rules {
		first := slices.Index(update, rule[0])
		second := slices.Index(update, rule[1])
		// pt1,
		// if first == -1 || second == -1 || second > first {
		// pt2
		if first == -1 || second == -1 {
			continue
		}
		matchedRules = append(matchedRules, rule)
	}
	// topological sort, proper way to do this to save for later
	g := &Graph{
		edges:    make(map[int][]int),
		vertices: update,
	}
	for _, rule := range matchedRules {
		g.addEdge(rule[0], rule[1])
	}
	update = g.topologicalSort()
	updateLength := len(update)
	if updateLength%2 == 0 {
		log.Panicf("Somehow it is an even number?: %d", updateLength)
	}
	num := update[updateLength/2]
	total += num
}

func main() {
	// fileName := "testinput.txt"
	fileName := "input.txt"
	file, err := os.Open(fileName)
	checkError((err))
	defer file.Close()
	gettingRules := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		if lineText == "" {
			gettingRules = false
			continue
		}
		if gettingRules {
			getRule(lineText)
		} else {
			getPageUpdate(lineText)
		}
	}
	checkUpdates()
	if !part1 {
		for _, incorrectUpdate := range incorrectUpdatesUnsorted {
			sortIncorrectUpdate(incorrectUpdate)
		}
	}
	log.Printf("Total number is: %d", total)
}

// Topological sort implementation for pt 2:
// https://reintech.io/blog/topological-sorting-in-go
type Graph struct {
	edges    map[int][]int
	vertices []int
}

func (g *Graph) addEdge(u, v int) {
	g.edges[u] = append(g.edges[u], v)
}

func (g *Graph) topologicalSortUtil(v int, visited map[int]bool, stack *[]int) {
	visited[v] = true
	for _, u := range g.edges[v] {
		if !visited[u] {
			g.topologicalSortUtil(u, visited, stack)
		}
	}
	*stack = append([]int{v}, *stack...)
}

func (g *Graph) topologicalSort() []int {
	stack := []int{}
	visited := make(map[int]bool)
	for _, v := range g.vertices {
		if !visited[v] {
			g.topologicalSortUtil(v, visited, &stack)
		}
	}

	return stack
}
