package main

import (
    "fmt"
    "log"
    "os"
    "bufio"
    "strings"
    "strconv"
)

type ElfBag struct {
    Red, Green, Blue int
}

func (bag ElfBag) IsGameRoundPossible(red, green, blue int) bool {
    return bag.Red > red && bag.Green > green && bag.Blue > blue
}

func Max(x, y int) int {
    if(y > x) {
        return y
    }
    return x
}

func main() {
    file, err := os.Open("input.txt")
    if(err != nil) {
        log.Fatal(err)
    }
    defer file.Close()
    
    fileScanner := bufio.NewScanner(file)
    var sumOfValidGameIds int = 0
    var sumPowerOfAllGames int = 0
    for fileScanner.Scan() {
        // do work
        line := fileScanner.Text()

        isCurrentGamePossible := true
        gameAndRounds := strings.Split(line, ": ")
        game := gameAndRounds[0][5:]

        rounds := strings.Split(gameAndRounds[1], ";")
        var maxRed, maxGreen, maxBlue int = 0, 0, 0

        for _, round := range rounds {
            colors := strings.Split(round, ",")
            var red, green, blue int = 0, 0, 0

            for _, color := range colors {
                colorParts := strings.Split(strings.Trim(color, " "), " ")
                colorValue, _ := strconv.Atoi(colorParts[0])
                switch(colorParts[1]) {
                    case "red": red = colorValue; break;
                    case "green": green = colorValue; break;
                    case "blue": blue = colorValue; break;
                }
            }

            if(red > 12 || green > 13 || blue > 14) {
                isCurrentGamePossible = false
            }

            maxRed = Max(red, maxRed)
            maxGreen = Max(green, maxGreen)
            maxBlue = Max(blue, maxBlue)
        }

        if(isCurrentGamePossible) {
            gameId, _ := strconv.Atoi(game)
            sumOfValidGameIds += gameId
        }

        sumPowerOfAllGames += maxRed * maxGreen * maxBlue;
    }

    if err := fileScanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("(Challenge 1) Sum of valid game ids: %d\n", sumOfValidGameIds)
    fmt.Printf("(Challenge 2) Sum of power of valid games: %d\n", sumPowerOfAllGames)
}
