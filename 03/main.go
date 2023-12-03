package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Number struct {
	value  int
	xStart int
	xEnd   int
	y      int
}

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
	var sumGearRatios int = 0

	for lineIndex, line := range lines {
		for charIndex := 0; charIndex < len(line); charIndex++ {
			char := line[charIndex]

			if isDigit(char) {
				number := getNumber(charIndex, line)
				number.y = lineIndex

				if isPartNumber(number, lines) {
					sumPartNumbers += number.value
					charIndex = number.xEnd
				}
			} else if char == '*' {
				xStart := Max(0, charIndex-1)
				xEnd := Min(len(line)-1, charIndex+1)
				yStart := Max(0, lineIndex-1)
				yEnd := Min(len(lines)-1, lineIndex+1)

				var numbers []Number

				for y := yStart; y <= yEnd; y++ {
					for x := xStart; x <= xEnd; x++ {
						if isDigit(lines[y][x]) {
							number := getNumber(x, lines[y])
							number.y = y
							numbers = append(numbers, number)
							x = number.xEnd
						}
					}
				}

				if len(numbers) == 2 {
					n1 := numbers[0]
					n2 := numbers[1]
					if isPartNumber(n1, lines) && isPartNumber(n2, lines) {
						sumGearRatios += n1.value * n2.value
					}
				}
			}
		}
	}

	fmt.Printf("# valid part numbers: %d\n", sumPartNumbers)
	fmt.Printf("sum of gear ratios: %d\n", sumGearRatios)
}

func isPartNumber(number Number, lines []string) bool {
	isPartNumber := false

	for x := number.xStart; x <= number.xEnd; x++ {
		xStart := Max(0, x-1)
		xEnd := Min(len(lines[number.y])-1, x+1)
		yStart := Max(0, number.y-1)
		yEnd := Min(len(lines)-1, number.y+1)

		// minesweeper stuff
		for x := xStart; x <= xEnd; x++ {
			for y := yStart; y <= yEnd; y++ {
				if !isDigit(lines[y][x]) && lines[y][x] != '.' {
					isPartNumber = true
				}
			}
		}
	}
	return isPartNumber
}

func getNumber(randomIndexWithinNumber int, line string) Number {
	startIndex := randomIndexWithinNumber
	endIndex := randomIndexWithinNumber
	var startIndexFound bool = false
	var endIndexFound bool = false

	// look behind
	for !startIndexFound {
		if startIndex-1 >= 0 && isDigit(line[startIndex-1]) {
			startIndex--
		} else {
			startIndexFound = true
		}
	}

	// look ahead
	for !endIndexFound {
		if endIndex+1 < len(line) && isDigit(line[endIndex+1]) {
			endIndex++
		} else {
			endIndexFound = true
		}
	}

	number, _ := strconv.Atoi(line[startIndex : endIndex+1])
	return Number{value: number, xStart: startIndex, xEnd: endIndex}
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
