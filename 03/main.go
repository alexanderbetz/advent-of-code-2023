package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var sumPartNumbers int = 0

	for lineIndex, line := range lines {
		// minesweeper stuff
		var numberStartIndex int = -1
		var numberEndIndex int = -1
		for charIndex := 0; charIndex < len(line); charIndex++ {
			char := line[charIndex]
			if isDigit(char) {
				// set bounds for symbol lookup
				numberStartIndex = charIndex
				numberEndIndex = charIndex
				xStart := Max(0, charIndex-1)
				xEnd := Min(len(line)-1, charIndex+1)
				yStart := Max(0, lineIndex-1)
				yEnd := Min(len(lines)-1, lineIndex+1)

				isPartNumber := false

				for x := xStart; x <= xEnd; x++ {
					for y := yStart; y <= yEnd; y++ {
						if !isDigit(lines[y][x]) && lines[y][x] != '.' {
							isPartNumber = true
						}
					}
				}

				if isPartNumber {
					var startIndexFound bool = false
					var endIndexFound bool = false
					// look behind
					for !startIndexFound {
						if numberStartIndex-1 >= 0 && isDigit(line[numberStartIndex-1]) {
							numberStartIndex--
						} else {
							startIndexFound = true
						}
					}

					// look ahead
					for !endIndexFound {
						if numberEndIndex+1 < len(line) && isDigit(line[numberEndIndex+1]) {
							numberEndIndex++
						} else {
							endIndexFound = true
						}
					}

					number, _ := strconv.Atoi(line[numberStartIndex : numberEndIndex+1])
					sumPartNumbers += number

					charIndex = numberEndIndex + 1
				}
			}
		}
	}

	fmt.Printf("# valid part numbers: %d\n", sumPartNumbers)
}

func isDigit(char byte) bool {
	return char >= 48 && char <= 57
}

func Min(x, y int) int {
	if y > x {
		return x
	}
	return y
}

func Max(x, y int) int {
	if y > x {
		return y
	}
	return x
}
