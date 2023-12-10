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
	part2MazeMask := make([]string, len(maze))
	part2OnesCount := 0
	for i := 0; i < len(part2MazeMask); i++ {
		part2MazeMask[i] = strings.Repeat(".", len(maze[i]))
	}

	// find out what S is
	l := maze[y][x-1]
	t := maze[y-1][x]
	r := maze[y][x+1]
	b := maze[y+1][x]
	hasTop := t == '|' || t == 'F' || t == '7'
	hasLeft := l == '-' || l == 'F' || l == 'L'
	hasBottom := b == 'L' || b == 'J' || b == '|'
	hasRight := r == '7' || r == 'J' || r == '-'
	var s byte = 'S'

	if hasTop && hasBottom {
		s = '|'
	} else if hasTop && hasLeft {
		s = 'J'
	} else if hasTop && hasRight {
		s = 'L'
	} else if hasLeft && hasRight {
		s = '-'
	} else if hasLeft && hasBottom {
		s = '7'
	} else if hasRight && hasBottom {
		s = 'F'
	}

	maze[y] = ReplaceIndex(maze[y], s, x)

	for !finishedLoop {
		tile := maze[y][x]
		part2MazeMask[y] = ReplaceIndex(part2MazeMask[y], tile, x)

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
		}

		if x == startingX && y == startingY {
			finishedLoop = true
		}
		pathLength++
	}

	// even-odd rule from vectors
	// if a tile has an even amount of edges on each side -> it's outside
	// else it's inside
	for i := 0; i < len(part2MazeMask); i++ {
		within := false
		up := '!'

		for _, rowChar := range part2MazeMask[i] {
			if rowChar == '|' {
				within = !within
			} else if rowChar == 'L' || rowChar == 'F' {
				up = rowChar
			} else if rowChar == 'J' || rowChar == '7' {
				if (up == 'L' && rowChar == '7') || (up == 'F' && rowChar == 'J') {
					within = !within
				}
				up = '!'
			}
			if within && rowChar == '.' {
				part2OnesCount++
			}
		}
	}

	fmt.Printf("(Challenge 1): Distance to furthest away tile: %d\n", pathLength/2)
	fmt.Printf("(Challenge 2): Enclosed tiles count: %d\n", part2OnesCount)
}

func ReplaceIndex(s string, char byte, index int) string {
	out := []rune(s)
	out[index] = []rune(string(char))[0]
	return string(out)
}

func isEnclosed(mazeMask []string, y, x int) bool {
	if x == 0 || x == len(mazeMask[0])-1 || y == 0 || y == len(mazeMask)-1 {
		return false
	}

	if mazeMask[y-1][x] == '0' || mazeMask[y+1][x] == '0' || mazeMask[y][x-1] == '0' || mazeMask[y][x+1] == '0' {
		return false
	}

	return true
}
