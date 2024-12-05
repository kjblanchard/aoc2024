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
	rules       [][]int
	pageUpdates [][]int
	total       int
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
			goodUpdate = false
			break
		}
		if goodUpdate {
			updateLength := len(update)
			if updateLength%2 == 0 {
				log.Panicf("Somehow it is an even number?: %d", updateLength)
			}
			total += update[updateLength/2]
		}
	}
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
	log.Printf("Total number is: %d", total)
}
