package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var fileName string
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . \"input.txt\"")
		os.Exit(1)
	}
	fileName = os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	directions := scanner.Text()

	scanner.Scan()

	stepMap := make(map[string][]string)
	for scanner.Scan() {
		a := strings.Split(scanner.Text(), " = (")
		b := strings.Split(a[1], ", ")
		stepMap[a[0]] = []string{b[0], b[1][:3]}
	}

	var currentSteps []string
	for step, _ := range stepMap {
		if step[2] == 'A' {
			currentSteps = append(currentSteps, step)
		}
	}

	var stepCounts []int = make([]int, len(currentSteps))
	steps := 0
	for true {
		leftOrRight := 0
		if directions[steps%len(directions)] == 'R' {
			leftOrRight = 1
		}

		for i := 0; i < len(currentSteps); i++ {
			if currentSteps[i][2] != 'Z' {
				currentSteps[i] = stepMap[currentSteps[i]][leftOrRight]
				stepCounts[i]++
			}
		}

		steps++

		allZs := true
		for i := 0; i < len(currentSteps); i++ {
			if currentSteps[i][2] != 'Z' {
				allZs = false
				break
			}
		}

		if allZs {
			break
		}
	}

	fmt.Printf("Steps required for a ghost: %d\n", LCM(stepCounts[0], stepCounts[1], stepCounts[2:]...))
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
