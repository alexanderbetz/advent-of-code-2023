package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

type CategoryConverter struct {
	fromCategory string
	toCategory   string
	conversion   []Conversion
}

type Conversion struct {
	destinationStart int
	sourceStart      int
	rangeLength      int
}

type Seed struct {
	categoryValues map[string]int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var seeds []Seed

	scanner.Scan()
	seedLine := scanner.Text()
	seedStrings := regexp.MustCompile("\\d+").FindAllString(seedLine, -1)

	for _, seedString := range seedStrings {
		n, _ := strconv.Atoi(seedString)
		seed := new(Seed)
		seed.categoryValues = make(map[string]int)
		seed.categoryValues["seed"] = n
		seeds = append(seeds, *seed)
	}

	scanner.Scan()
	mapRegex := regexp.MustCompile("(\\w+)-to-(\\w+)")
	conversionRegex := regexp.MustCompile("(\\d+) (\\d+) (\\d+)")
	var categoryOrder []string
	categoryOrder = append(categoryOrder, "seed")
	conversionMap := make(map[string][]Conversion)
	var currentCategory string
	for scanner.Scan() {
		line := scanner.Text()

		if categories := mapRegex.FindStringSubmatch(line); len(categories) == 3 {
			currentCategory = categories[2]
			categoryOrder = append(categoryOrder, currentCategory)
		} else if conversion := conversionRegex.FindStringSubmatch(line); len(conversion) == 4 {
			destinationStart, _ := strconv.Atoi(conversion[1])
			sourceStart, _ := strconv.Atoi(conversion[2])
			rangeLength, _ := strconv.Atoi(conversion[3])
			conversion := Conversion{destinationStart: destinationStart, sourceStart: sourceStart, rangeLength: rangeLength}
			conversionMap[currentCategory] = append(conversionMap[currentCategory], conversion)
		}
	}

	for _, seed := range seeds {
		for i := 1; i < len(categoryOrder); i++ {
			convertToCategory(seed, categoryOrder[i], categoryOrder[i-1], conversionMap)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	lowestLocation := math.MaxInt
	for _, seed := range seeds {
		lowestLocation = Min(seed.categoryValues["location"], lowestLocation)
	}

	fmt.Printf("Lowest location: %d\n", lowestLocation)
}

func convertToCategory(seed Seed, destinationCategory string, sourceCategory string, conversionMap map[string][]Conversion) {
	seed.categoryValues[destinationCategory] = seed.categoryValues[sourceCategory]

	sourceValue := seed.categoryValues[sourceCategory]
	for _, conversion := range conversionMap[destinationCategory] {
		if sourceValue >= conversion.sourceStart && sourceValue < conversion.sourceStart+conversion.rangeLength {
			diff := seed.categoryValues[sourceCategory] - conversion.sourceStart
			seed.categoryValues[destinationCategory] = conversion.destinationStart + diff
		}
	}
}

func Min(x, y int) int {
	if y < x {
		return y
	}
	return x
}
