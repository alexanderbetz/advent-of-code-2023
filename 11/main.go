package main

import (
	"bufio"
	"fmt"
	"github.com/schollz/progressbar/v3"
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

	var input []string

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	totalStepsPart1 := calcSteps(input, 1)
	totalStepsPart2 := calcSteps(input, 999999)

	fmt.Printf("(Challenge 1) Sum of shortest distances: %d\n", totalStepsPart1)
	fmt.Printf("(Challenge 2) Sum of shortest distances: %d\n", totalStepsPart2)
}

func calcSteps(galaxy []string, expansionFactor int) int {
	var expandedXs []int
	var expandedYs []int

	// expand lines
	for y, line := range galaxy {
		if strings.Count(line, ".") == len(line) {
			expandedYs = append(expandedYs, y)
		}
	}

	// expand cols
	for col := len(galaxy[0]) - 1; col >= 0; col-- {
		shouldExpand := true
		for row := 0; row < len(galaxy); row++ {
			if galaxy[row][col] != '.' {
				shouldExpand = false
				break
			}
		}
		if shouldExpand {
			expandedXs = append(expandedXs, col)
		}
	}

	var galaxies [][]int

	for y, v := range galaxy {
		for x, ch := range v {
			if ch == '#' {
				galaxies = append(galaxies, []int{x, y})
			}
		}
	}

	bar := progressbar.Default(int64(len(galaxies)))

	totalDistances := 0
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			x0 := galaxies[i][0]
			y0 := galaxies[i][1]
			x1 := galaxies[j][0]
			y1 := galaxies[j][1]
			betweenX := 0
			for _, x := range expandedXs {
				if Min(x0, x1) < x && x < Max(x0, x1) {
					betweenX += expansionFactor
				}
			}
			betweenY := 0
			for _, y := range expandedYs {
				if Min(y0, y1) < y && y < Max(y0, y1) {
					betweenY += expansionFactor
				}
			}

			dx := abs(x1-x0) + betweenX
			dy := abs(y1-y0) + betweenY
			x1 += betweenX
			y1 += betweenY
			distance := dx + dy
			totalDistances += distance
		}

		bar.Add(1)
	}

	return totalDistances
}

func Min(x, y int) int {
	if y < x {
		return y
	}
	return x
}

func Max(x, y int) int {
	if y > x {
		return y
	}
	return x
}

func AddIndex(s string, char byte, index int) string {
	out := s
	out = out[:index] + string(char) + out[index:]
	return out
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

// https://github.com/StephaneBunel/bresenham/blob/ec76d7b8e923/drawline.go
func Bresenham(x1, y1, x2, y2 int) int {
	var dx, dy, e, slope int

	// Because drawing p1 -> p2 is equivalent to draw p2 -> p1,
	// I sort points in x-axis order to handle only half of possible cases.
	if x1 > x2 {
		x1, y1, x2, y2 = x2, y2, x1, y1
	}

	dx, dy = x2-x1, y2-y1
	// Because point is x-axis ordered, dx cannot be negative
	if dy < 0 {
		dy = -dy
	}

	pointsCrossedCount := 0

	switch {

	// Is line a point ?
	case x1 == x2 && y1 == y2:
		// p.Set(x1, y1, col)
		pointsCrossedCount = 0

	// Is line an horizontal ?
	case y1 == y2:
		for ; dx != 0; dx-- {
			// p.Set(x1, y1, col)
			// fmt.Println(x1, y1)
			pointsCrossedCount++
			x1++
		}
		// p.Set(x1, y1, col)

	// Is line a vertical ?
	case x1 == x2:
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for ; dy != 0; dy-- {
			// p.Set(x1, y1, col)
			// fmt.Println(x1, y1)
			pointsCrossedCount++
			y1++
		}
		// p.Set(x1, y1, col)

	// Is line a diagonal ?
	case dx == dy:
		if y1 < y2 {
			for ; dx != 0; dx-- {
				// p.Set(x1, y1, col)
				pointsCrossedCount += 2
				x1++
				y1++
			}
		} else {
			for ; dx != 0; dx-- {
				// p.Set(x1, y1, col)
				pointsCrossedCount += 2
				x1++
				y1--
			}
		}
		// p.Set(x1, y1, col)

	// wider than high ?
	case dx > dy:
		if y1 < y2 {
			// BresenhamDxXRYD(img, x1, y1, x2, y2, col)
			dy, e, slope = 2*dy, dx, 2*dx
			for ; dx != 0; dx-- {
				// p.Set(x1, y1, col)
				pointsCrossedCount++
				x1++
				e -= dy
				if e < 0 {
					pointsCrossedCount++
					y1++
					e += slope
				}
			}
		} else {
			// BresenhamDxXRYU(img, x1, y1, x2, y2, col)
			dy, e, slope = 2*dy, dx, 2*dx
			for ; dx != 0; dx-- {
				// p.Set(x1, y1, col)
				pointsCrossedCount++
				x1++
				e -= dy
				if e < 0 {
					pointsCrossedCount++
					y1--
					e += slope
				}
			}
		}
		// p.Set(x2, y2, col)

	// higher than wide.
	default:
		if y1 < y2 {
			// BresenhamDyXRYD(img, x1, y1, x2, y2, col)
			dx, e, slope = 2*dx, dy, 2*dy
			for ; dy != 0; dy-- {
				// p.Set(x1, y1, col)
				pointsCrossedCount++
				y1++
				e -= dx
				if e < 0 {
					pointsCrossedCount++
					x1++
					e += slope
				}
			}
		} else {
			// BresenhamDyXRYU(img, x1, y1, x2, y2, col)
			dx, e, slope = 2*dx, dy, 2*dy
			for ; dy != 0; dy-- {
				// p.Set(x1, y1, col)
				pointsCrossedCount++
				y1--
				e -= dx
				if e < 0 {
					pointsCrossedCount++
					x1++
					e += slope
				}
			}
		}
		// p.Set(x2, y2, col)
	}

	return pointsCrossedCount
}
