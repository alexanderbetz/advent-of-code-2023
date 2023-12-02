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

func main() {
    file, err := os.Open("input.txt")
    if(err != nil) {
        log.Fatal(err)
    }
    defer file.Close()
    
    fileScanner := bufio.NewScanner(file)
    var sumOfValidGameIds int = 0
    for fileScanner.Scan() {
        // do work
        line := fileScanner.Text()

        isCurrentGamePossible := true
        gameAndRounds := strings.Split(line, ": ")
        game := gameAndRounds[0][5:]

        rounds := strings.Split(gameAndRounds[1], ";")

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
                break
            }
        }

        if(isCurrentGamePossible) {
            gameId, _ := strconv.Atoi(game)
            sumOfValidGameIds += gameId
        }
    }

    if err := fileScanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Sum of valid game ids: %d\n", sumOfValidGameIds)
}
