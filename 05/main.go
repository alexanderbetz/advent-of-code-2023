package main

import (
	"bufio"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Conversion struct {
	destinationStart int
	sourceStart      int
	rangeLength      int
}

type SeedRange struct {
	start       int
	rangeLength int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var seeds []int
	var seedRanges []SeedRange

	scanner.Scan()
	seedLine := scanner.Text()
	seedStrings := regexp.MustCompile("\\d+").FindAllString(seedLine, -1)

	for i, seedString := range seedStrings {
		n, _ := strconv.Atoi(seedString)
		seeds = append(seeds, n)

		if i%2 == 0 {
			seedRanges = append(seedRanges, SeedRange{start: n})
		} else {
			seedRanges[i/2].rangeLength = n
		}
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

	lowestLocation := math.MaxInt
	for _, seed := range seeds {
		for i := 1; i < len(categoryOrder); i++ {
			seed = convertToCategory(seed, categoryOrder[i], categoryOrder[i-1], conversionMap)
			if categoryOrder[i] == "location" {
				lowestLocation = Min(lowestLocation, seed)
			}
		}
	}
	fmt.Printf("(Challenge 1) Lowest location: %d\n", lowestLocation)

	var seedsCount int

	for _, seedRange := range seedRanges {
		seedsCount += seedRange.rangeLength
	}

	var count int
	lowestLocationChallenge2 := math.MaxInt
	bar := progressbar.Default(int64(seedsCount))
	for _, seedRange := range seedRanges {
		for i := seedRange.start; i < seedRange.start+seedRange.rangeLength; i++ {
			seedValue := i
			for j := 1; j < len(categoryOrder); j++ {
				seedValue = convertToCategory(seedValue, categoryOrder[j], categoryOrder[j-1], conversionMap)
				if categoryOrder[j] == "location" {
					lowestLocationChallenge2 = Min(lowestLocationChallenge2, seedValue)
				}
			}
			count++
			if count%10000000 == 0 {
				bar.Add(10000000)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("(Challenge 2) Lowest location: %d\n", lowestLocationChallenge2)
}

func convertToCategory(seed int, destinationCategory string, sourceCategory string, conversionMap map[string][]Conversion) int {
	destinationValue := seed

	for _, conversion := range conversionMap[destinationCategory] {
		if seed >= conversion.sourceStart && seed < conversion.sourceStart+conversion.rangeLength {
			diff := seed - conversion.sourceStart
			destinationValue = conversion.destinationStart + diff
			return destinationValue
		}
	}

	return destinationValue
}

func Min(x, y int) int {
	if y < x {
		return y
	}
	return x
}
