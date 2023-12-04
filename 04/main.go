package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sumOfWinningMatches int = 0
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(strings.Split(line, ": ")[1], " | ")
		winningNumbers := getNumbersFromString(numbers[0])
		_myNumbers := getNumbersFromString(numbers[1])
		myNumbers := map[int]bool{}

		for _, n := range _myNumbers {
			myNumbers[n] = true
		}

		var score int = 0
		for _, wn := range winningNumbers {
			if myNumbers[wn] {
				if score == 0 {
					score = 1
				} else {
					score = score * 2
				}
			}
		}

		sumOfWinningMatches += score
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sum of winning matches: %d\n", sumOfWinningMatches)
}

func getNumbersFromString(line string) (numbers []int) {
	for _, str := range strings.Split(line, " ") {
		number, err := strconv.Atoi(str)
		if err != nil {
			continue
		}
		numbers = append(numbers, number)
	}
	return numbers
}
