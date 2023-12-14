package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	var rocks []string

	for scanner.Scan() {
		rocks = append(rocks, scanner.Text())
	}

	part1(rocks)

	currentNorthWeight := 0
	cycleCount := 0
	seen := make(map[string]int)
	var seenIndex []string
	firstIndexOfCycle := 0

	for true {
		// tilt round rocks north
		for x := 0; x < len(rocks[0]); x++ {
			mostNorthFreeSpace := 0
			for y := 0; y < len(rocks); y++ {
				if rocks[y][x] == '#' {
					mostNorthFreeSpace = y + 1
				} else if rocks[y][x] == 'O' {
					// move north
					rocks[y] = ReplaceIndex(rocks[y], '.', x)
					rocks[mostNorthFreeSpace] = ReplaceIndex(rocks[mostNorthFreeSpace], 'O', x)
					mostNorthFreeSpace++
				}
			}
		}
		// tilt round rocks west
		for y := 0; y < len(rocks); y++ {
			mostWestFreeSpace := 0
			for x := 0; x < len(rocks[y]); x++ {
				if rocks[y][x] == '#' {
					mostWestFreeSpace = x + 1
				} else if rocks[y][x] == 'O' {
					// move west
					rocks[y] = ReplaceIndex(rocks[y], '.', x)
					rocks[y] = ReplaceIndex(rocks[y], 'O', mostWestFreeSpace)
					mostWestFreeSpace++
				}
			}
		}
		// tilt round rocks south
		for x := 0; x < len(rocks[0]); x++ {
			mostSouthFreeSpace := len(rocks) - 1
			for y := len(rocks) - 1; y >= 0; y-- {
				if rocks[y][x] == '#' {
					mostSouthFreeSpace = y - 1
				} else if rocks[y][x] == 'O' {
					// move south
					rocks[y] = ReplaceIndex(rocks[y], '.', x)
					rocks[mostSouthFreeSpace] = ReplaceIndex(rocks[mostSouthFreeSpace], 'O', x)
					mostSouthFreeSpace--
				}
			}
		}
		// tilt round rocks east
		for y := 0; y < len(rocks); y++ {
			mostEastFreeSpace := len(rocks[0]) - 1
			for x := len(rocks[y]) - 1; x >= 0; x-- {
				if rocks[y][x] == '#' {
					mostEastFreeSpace = x - 1
				} else if rocks[y][x] == 'O' {
					// move west
					rocks[y] = ReplaceIndex(rocks[y], '.', x)
					rocks[y] = ReplaceIndex(rocks[y], 'O', mostEastFreeSpace)
					mostEastFreeSpace--
				}
			}
		}

		cycleCount++
		rocksAsOneLine := strings.Join(rocks, ",")
		if v, ok := seen[rocksAsOneLine]; ok {
			firstIndexOfCycle = v
			break
		}
		seen[rocksAsOneLine] = cycleCount
		seenIndex = append(seenIndex, rocksAsOneLine)
	}

	rocksIndex := (1000000000-(firstIndexOfCycle))%(cycleCount-firstIndexOfCycle) + firstIndexOfCycle

	currentNorthWeight = 0
	for i, v := range strings.Split(seenIndex[rocksIndex-1], ",") {
		weight := len(rocks) - i
		currentNorthWeight += weight * strings.Count(v, "O")
	}

	fmt.Printf("(Challenge 2): Sum of weights after 1.000.000.000 cycles: %d\n", currentNorthWeight)
}

func part1(r []string) {
	rocks := make([]string, len(r))
	copy(rocks, r)

	for x := 0; x < len(rocks[0]); x++ {
		mostNorthFreeSpace := 0
		for y := 0; y < len(rocks); y++ {
			if rocks[y][x] == '#' {
				mostNorthFreeSpace = y + 1
			} else if rocks[y][x] == 'O' {
				// move north
				rocks[y] = ReplaceIndex(rocks[y], '.', x)
				rocks[mostNorthFreeSpace] = ReplaceIndex(rocks[mostNorthFreeSpace], 'O', x)
				mostNorthFreeSpace++
			}
		}
	}
	currentNorthWeight := 0
	for i, v := range rocks {
		weight := len(rocks) - i
		currentNorthWeight += weight * strings.Count(v, "O")
	}

	fmt.Printf("(Challenge 1): Sum of weights: %d\n", currentNorthWeight)
}

func ReplaceIndex(s string, char byte, index int) string {
	out := []rune(s)
	out[index] = []rune(string(char))[0]
	return string(out)
}
