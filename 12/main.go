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
	sumPossibilitiesPart2 := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		fragments := line[0]
		fragmentsPart2 := strings.Repeat(line[0]+"?", 5)
		fragmentsPart2 = fragmentsPart2[:len(fragmentsPart2)-1]
		brokenSpringsStr := strings.Split(line[1], ",")
		bsPart2 := strings.Repeat(line[1]+",", 5)
		brokenSpringsStrPart2 := strings.Split(bsPart2[:len(bsPart2)-1], ",")
		var brokenSprings []int
		var brokenSpringsPart2 []int
		for _, s := range brokenSpringsStr {
			n, _ := strconv.Atoi(s)
			brokenSprings = append(brokenSprings, n)
		}
		for _, s := range brokenSpringsStrPart2 {
			n, _ := strconv.Atoi(s)
			brokenSpringsPart2 = append(brokenSpringsPart2, n)
		}

		cache = make(map[string]int)
		sumPossibilities += Count2(fragments, brokenSprings, 0, 0, 0)
		cache = make(map[string]int)
		sumPossibilitiesPart2 += Count2(fragmentsPart2, brokenSpringsPart2, 0, 0, 0)
	}

	fmt.Printf("(Challenge 1): Sum of possible arrangements: %d\n", sumPossibilities)
	fmt.Printf("(Challenge 2): Sum of possible arrangements: %d\n", sumPossibilitiesPart2)
}

var cache map[string]int = make(map[string]int)

// # i == current position within dots
// # bi == current position within blocks
// # current == length of current block of '#'
// # state space is len(dots) * len(blocks) * len(dots)
func Count2(dots string, blocks []int, i, bi, current int) int {
	var key string = strconv.Itoa(i) + " " + strconv.Itoa(bi) + " " + strconv.Itoa(current)
	if val, ok := cache[key]; ok {
		return val
	}

	if i == len(dots) {
		if bi == len(blocks) && current == 0 {
			return 1
		} else if bi == len(blocks)-1 && blocks[bi] == current {
			return 1
		} else {
			return 0
		}
	}

	result := 0

	for _, ch := range []byte{'.', '#'} {
		if dots[i] == ch || dots[i] == '?' {
			if ch == '.' && current == 0 {
				result += Count2(dots, blocks, i+1, bi, 0)
			} else if ch == '.' && current > 0 && bi < len(blocks) && blocks[bi] == current {
				result += Count2(dots, blocks, i+1, bi+1, 0)
			} else if ch == '#' {
				result += Count2(dots, blocks, i+1, bi, current+1)
			}
		}
	}
	cache[key] = result
	return result
}

func Count(fragment string, nums []int) int {
	if len(fragment) == 0 {
		if len(nums) == 0 {
			return 1
		} else {
			return 0
		}
	}

	if len(nums) == 0 {
		if strings.Contains(fragment, "#") {
			return 0
		} else {
			return 1
		}
	}

	result := 0

	if fragment[0] == '.' || fragment[0] == '?' {
		result += Count(fragment[1:], nums)
	}

	if fragment[0] == '#' || fragment[0] == '?' {
		if nums[0] <= len(fragment) && !strings.Contains(fragment[:nums[0]], ".") && (nums[0] == len(fragment) || fragment[nums[0]] != '#') {
			result += Count(fragment[nums[0]+1:], nums[1:])
		}
	}

	return result
}
