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

	currentStep := "AAA"
	stepCount := 0
	for true {
		leftOrRight := 0
		if directions[stepCount%len(directions)] == 'R' {
			leftOrRight = 1
		}
		currentStep = stepMap[currentStep][leftOrRight]
		stepCount++

		if currentStep == "ZZZ" {
			break
		}
	}

	fmt.Printf("Steps to ZZZ: %d\n", stepCount)
}
