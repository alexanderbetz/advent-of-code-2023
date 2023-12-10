package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: go run . input.txt")
		os.Exit(1)
	}
	file, err := os.Open(args[1])
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var maze []string
	startingX := 0
	startingY := 0
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
		indexOfS := strings.Index(line, "S")
		if indexOfS >= 0 {
			startingX = indexOfS
			startingY = y
		}
		y++
	}

	finishedLoop := false
	pathLength := 0
	x := startingX
	y = startingY
	// contains 't', 'l', 'r', 'b' depending on from which direction i am comming
	commingFrom := 'S'

	// move 1 tile away from S to start path tracing
	l := maze[y][x-1]
	t := maze[y-1][x]
	r := maze[y][x+1]
	if l == '-' || l == 'F' || l == 'L' {
		x--
		commingFrom = 'r'
	} else if t == '|' || t == 'F' || t == '7' {
		y--
		commingFrom = 'b'
	} else if r == 'J' || r == '-' || r == '7' {
		x++
		commingFrom = 'l'
	} else {
		y++
		commingFrom = 't'
	}

	for !finishedLoop {
		tile := maze[y][x]

		switch tile {
		case '|':
			{
				if commingFrom == 't' {
					y++
				} else {
					y--
				}
				break
			}
		case '-':
			{
				if commingFrom == 'l' {
					x++
				} else {
					x--
				}
				break
			}
		case 'J':
			{
				if commingFrom == 'l' {
					y--
					commingFrom = 'b'
				} else {
					x--
					commingFrom = 'r'
				}
				break
			}
		case 'L':
			{
				if commingFrom == 't' {
					x++
					commingFrom = 'l'
				} else {
					y--
					commingFrom = 'b'
				}
				break
			}
		case '7':
			{
				if commingFrom == 'b' {
					x--
					commingFrom = 'r'
				} else {
					y++
					commingFrom = 't'
				}
				break
			}
		case 'F':
			{
				if commingFrom == 'b' {
					x++
					commingFrom = 'l'
				} else {
					y++
					commingFrom = 't'
				}
				break
			}
		case 'S':
			finishedLoop = true
		}

		pathLength++
	}

	fmt.Printf("(Challenge 1): Distance to furthest away tile: %d\n", pathLength/2)
}
