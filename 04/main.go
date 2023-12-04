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
	var gameWinsCopies []int
	var gameIndex int = 0
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
		gameWinsCopies = append(gameWinsCopies, 0)
		for _, wn := range winningNumbers {
			if myNumbers[wn] {
				if score == 0 {
					score = 1
				} else {
					score = score * 2
				}
				gameWinsCopies[gameIndex]++
			}
		}

		sumOfWinningMatches += score
		gameIndex++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("(Challenge 1) Sum of winning matches: %d\n", sumOfWinningMatches)
	fmt.Printf("(Challenge 2) Sum of total scratch cards: %d\n", countTotalScratchCards(gameWinsCopies))
}

func countTotalScratchCards(scratchCardWinsCopies []int) int {
	var totalScratchCards int = 0
	var scratchCardExistsXTimes []int = make([]int, len(scratchCardWinsCopies))

	for scratchCardIndex, scratchCardCopies := range scratchCardWinsCopies {
		scratchCardExistsXTimes[scratchCardIndex]++

		for i := 0; i < scratchCardCopies; i++ {
			scratchCardExistsXTimes[scratchCardIndex+i+1] += scratchCardExistsXTimes[scratchCardIndex]
		}
	}

	for _, existCount := range scratchCardExistsXTimes {
		totalScratchCards += existCount
	}

	return totalScratchCards
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
