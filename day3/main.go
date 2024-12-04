package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	regexpression = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	numRegex      = regexp.MustCompile(`\d+`)
)

func checkError(e error) {
	if e != nil {
		log.Panicf("Failed for some reason, %s", e)
	}
}
func getNumbers(byteBoi []byte) (int, int) {
	strBytes := string(byteBoi)
	numMatches := numRegex.FindAllString(strBytes, -1)
	var numberSlice []int
	for _, numbers := range numMatches {
		num, err := strconv.Atoi(numbers)
		checkError(err)
		numberSlice = append(numberSlice, num)
	}
	if len(numberSlice) != 2 {
		log.Panic("More than 2 numbers")
	}
	return numberSlice[0], numberSlice[1]
}

func main() {
	data, err := os.ReadFile("input.txt")
	checkError(err)
	matches := regexpression.FindAll(data, -1)
	total := 0
	for _, v := range matches {
		lhs, rhs := getNumbers(v)
		total += lhs * rhs
	}
	log.Printf("Total is %d", total)
}
