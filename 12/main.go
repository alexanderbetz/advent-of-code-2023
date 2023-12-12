package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . input.txt")
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	sumPossibilities := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		fragments := line[0]
		brokenSpringsStr := strings.Split(line[1], ",")
		var brokenSprings []int
		for _, s := range brokenSpringsStr {
			n, _ := strconv.Atoi(s)
			brokenSprings = append(brokenSprings, n)
		}
		// brute force all possibilities by building every possible arrangement and then testing it

		// every unknown fragment has 2 possible values
		unknownCount := strings.Count(fragments, "?")
		var unknownIndexes []int
		unknownIndexes = append(unknownIndexes, strings.Index(fragments, "?"))
		for len(unknownIndexes) < unknownCount {
			unknownIndexes = append(unknownIndexes, unknownIndexes[len(unknownIndexes)-1]+1+strings.Index(fragments[unknownIndexes[len(unknownIndexes)-1]+1:], "?"))
		}
		bruteForceCount := 1 << unknownCount

		fmt.Println(bruteForceCount, line)

		currentTry := 0
		for currentTry < bruteForceCount {
			/// ### => 111
			/// 0## => 011
			reconstructed := fragments

			for i := 0; i < unknownCount; i++ {
				var ch byte = '.'
				if (currentTry>>i)%2 == 1 {
					ch = '#'
				}
				reconstructed = ReplaceIndex(reconstructed, ch, unknownIndexes[i])
			}

			if testArrangement(reconstructed, brokenSprings) {
				sumPossibilities++
			}

			currentTry++
		}
	}

	fmt.Printf("(Challenge 1): Sum of possible arrangements: %d\n", sumPossibilities)
}

func testArrangement(arrangement string, brokenSprings []int) bool {
	// get regex matches
	// test if length == len(brokenSprings)
	// test if every len(match) == brokenSprings[i]

	matches := regexp.MustCompile("#+").FindAllString(arrangement, -1)

	if len(matches) != len(brokenSprings) {
		return false
	}

	for i := 0; i < len(matches); i++ {
		if len(matches[i]) != brokenSprings[i] {
			return false
		}
	}
	return true
}

func ReplaceIndex(s string, char byte, index int) string {
	out := []rune(s)
	out[index] = []rune(string(char))[0]
	return string(out)
}
