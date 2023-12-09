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
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: go run . \"input.txt\"")
		os.Exit(1)
	}
	fileName := args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	sumOfExtrapolations := 0
	sumOfPreExtrapolations := 0
	for scanner.Scan() {
		var sequences [][]int = [][]int{{}}

		for _, num := range strings.Split(scanner.Text(), " ") {
			n, _ := strconv.Atoi(num)
			sequences[0] = append(sequences[0], n)

		}

		areDeltasEqual := false
		currentSequenceIdx := 0
		for !areDeltasEqual {
			var tmpDeltas []int
			_tmpDeltasEqual := true
			for i := 0; i < len(sequences[currentSequenceIdx])-1; i++ {
				tmpDeltas = append(tmpDeltas, sequences[currentSequenceIdx][i+1]-sequences[currentSequenceIdx][i])
				if i > 1 && tmpDeltas[i] != tmpDeltas[i-1] {
					_tmpDeltasEqual = false
				}
			}
			areDeltasEqual = _tmpDeltasEqual
			if !areDeltasEqual {
				sequences = append(sequences, tmpDeltas)
				currentSequenceIdx++
			}
		}

		for i := len(sequences) - 1; i >= 0; i-- {
			diff := 0
			prevcastDiff := 0
			if i+1 >= len(sequences) {
				diff = sequences[i][len(sequences[i])-1] - sequences[i][len(sequences[i])-2]
				prevcastDiff = diff
			} else {
				diff = sequences[i+1][len(sequences[i+1])-1]
				prevcastDiff = sequences[i+1][0]
			}
			forecast := sequences[i][len(sequences[i])-1] + diff
			prevcast := sequences[i][0] - prevcastDiff
			sequences[i] = append([]int{prevcast}, sequences[i]...)
			sequences[i] = append(sequences[i], forecast)
		}

		sumOfExtrapolations += sequences[0][len(sequences[0])-1]
		sumOfPreExtrapolations += sequences[0][0]
	}

	fmt.Printf("(Challenge 1) Sum of extrapolations: %d\n", sumOfExtrapolations)
	fmt.Printf("(Challenge 2) Sum of extrapolations: %d\n", sumOfPreExtrapolations)
}
