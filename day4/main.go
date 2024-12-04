package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const (
	fullWord    = "XMAS"
	UP          = iota
	RIGHT       = iota
	DOWN        = iota
	LEFT        = iota
	TOPRIGHT    = iota
	BOTTOMRIGHT = iota
	BOTTOMLEFT  = iota
	TOPLEFT     = iota
)

func checkError(e error) {
	if e != nil {
		log.Panicf("Failed for some reason, %s", e)
	}
}

func getLetterInFullFile(document string, column int, row int, lineLength int) rune {
	stripped := strings.ReplaceAll(document, "\n", "")
	if column < 0 || row < 0 || column*lineLength+row >= len(stripped) || row >= lineLength {
		return 0
	}
	totalCount := column*lineLength + row
	return rune(stripped[totalCount])
}

func checkDirection(column int, row int, currentLetter rune, lineLength int, document string, direction int) bool {
	if currentLetter == 'S' {
		return true
	}
	if column < 0 || row < 0 || row >= lineLength {
		return false
	}
	nextLetter := rune(fullWord[strings.IndexRune(fullWord, currentLetter)+1])
	nextLetterString := string(nextLetter)
	log.Println(nextLetterString)
	nextColumn := column
	nextRow := row
	if direction == TOPRIGHT || direction == UP || direction == TOPLEFT {
		nextColumn--
	} else if direction == BOTTOMLEFT || direction == DOWN || direction == BOTTOMRIGHT {
		nextColumn++
	}
	if direction == RIGHT || direction == TOPRIGHT || direction == BOTTOMRIGHT {
		nextRow++
	} else if direction == LEFT || direction == TOPLEFT || direction == BOTTOMLEFT {
		nextRow--
	}
	foundLetter := getLetterInFullFile(document, nextColumn, nextRow, lineLength)
	if foundLetter == nextLetter {
		return checkDirection(nextColumn, nextRow, foundLetter, lineLength, document, direction)
	}
	return false
}

func main() {
	// fileName := "/Users/kevin/git/misc/aoc2024/day4/testinput.txt"
	fileName := "/Users/kevin/git/misc/aoc2024/day4/input.txt"
	fullFile, err := os.ReadFile(fileName)
	checkError((err))
	file, err := os.Open(fileName)
	checkError((err))
	defer file.Close()
	scanner := bufio.NewScanner(file)
	currentColumn := 0
	xmasCount := 0
	for scanner.Scan() {
		lineText := scanner.Text()
		lineLength := len(lineText)
		for row, currentRune := range lineText {
			if currentRune == 'X' {
				for _, direction := range []int{UP, DOWN, LEFT, RIGHT, TOPRIGHT, BOTTOMRIGHT, BOTTOMLEFT, TOPLEFT} {
					if checkDirection(currentColumn, row, currentRune, lineLength, string(fullFile), direction) {
						xmasCount++
					}
				}
			}
		}
		currentColumn++
	}
	log.Printf("Xmas count is %d", xmasCount)
}
