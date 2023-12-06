package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	timeLine := scanner.Text()
	scanner.Scan()
	distanceLine := scanner.Text()
	numbersRegex := regexp.MustCompile("\\d+")

	var times []int
	var distances []int
	var timeCombinedStr string
	var distanceCombinedStr string
	for _, n := range numbersRegex.FindAllString(timeLine, -1) {
		timeCombinedStr += n
		parsedNumber, _ := strconv.Atoi(n)
		times = append(times, parsedNumber)
	}
	timeCombined, _ := strconv.Atoi(timeCombinedStr)

	for _, n := range numbersRegex.FindAllString(distanceLine, -1) {
		distanceCombinedStr += n
		parsedNumber, _ := strconv.Atoi(n)
		distances = append(distances, parsedNumber)
	}
	distanceCombined, _ := strconv.Atoi(distanceCombinedStr)

	combinedWinningPossibilities := calcWinningPossibilities(times, distances)
	fmt.Printf("(Challenge 1) Winning possibilities: %d\n", combinedWinningPossibilities)

	combinedWinningPossibilities = calcWinningPossibilities([]int{timeCombined}, []int{distanceCombined})
	fmt.Printf("(Challenge 2) Winning possibilities: %d\n", combinedWinningPossibilities)
}

func calcWinningPossibilities(times, distances []int) int {
	var winningPossibilities []int
	var combinedWinningPossibilities int = 1
	for i := 0; i < len(times); i++ {
		for j := 1; j < times[i]; j++ {
			speed := j
			timeLeftToCruise := times[i] - j
			travelDistance := speed * timeLeftToCruise
			if travelDistance > distances[i] {
				upperBound := times[i] - speed
				winningChargingTimes := upperBound + 1 - speed
				winningPossibilities = append(winningPossibilities, winningChargingTimes)
				combinedWinningPossibilities *= winningChargingTimes
				break
			}
		}
	}
	return combinedWinningPossibilities
}
