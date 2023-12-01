package main

import (
    "fmt"
    "os"
    "log"
    "strings"
    "regexp"
    "strconv"
)

func main() {
    file, err := os.Open("input.txt")
    if(err != nil) {
        log.Fatal(err)
    }

    var content string
    var contentLeftToRead bool = true

    for contentLeftToRead {
        data := make([]byte, 100)
        count, err := file.Read(data)
        if(err != nil) {
            log.Fatal(err)
        }
        contentLeftToRead = count == 100
        content += string(data)
    }

    lines := strings.Split(content, "\n")
    var sum1 int = 0
    var sum2 int = 0

    for _, line := range lines {
        numbers1 := getNumbers1(line)
        numbers2 := getNumbers2(line)
        if(len(numbers1) > 0) {
            sum1 += (numbers1[0] * 10) + numbers1[len(numbers1)-1]
        }
        if(len(numbers2) > 0) {
            sum2 += (numbers2[0] * 10) + numbers2[len(numbers2)-1]
        } 
    }
    fmt.Printf("Challenge 1: %d\n", sum1)
    fmt.Printf("Challenge 2: %d\n", sum2)
}

func getNumbers1(line string) []int {
    numbers := regexp.MustCompile(`\d`).FindAllString(line, -1)
    parsedNumbers := make([]int, len(numbers))
    for i, char := range numbers {
        n , err := strconv.Atoi(char)
        if(err != nil) {
            fmt.Println(err)
        }
        parsedNumbers[i] = int(n)
    }
    return parsedNumbers
}

func getNumbers2(line string) []int {
    var parsedNumbers []int
    tmpLine := line
    for {
        if(strings.HasPrefix(tmpLine, "1") || strings.HasPrefix(tmpLine, "one")) {
            parsedNumbers = append(parsedNumbers, 1)
        } else if(strings.HasPrefix(tmpLine, "2") || strings.HasPrefix(tmpLine, "two")) {
            parsedNumbers = append(parsedNumbers, 2)
        } else if(strings.HasPrefix(tmpLine, "3") || strings.HasPrefix(tmpLine, "three")) {
            parsedNumbers = append(parsedNumbers, 3)
        } else if(strings.HasPrefix(tmpLine, "4") || strings.HasPrefix(tmpLine, "four")) {
            parsedNumbers = append(parsedNumbers, 4)
        } else if(strings.HasPrefix(tmpLine, "5") || strings.HasPrefix(tmpLine, "five")) {
            parsedNumbers = append(parsedNumbers, 5)
        } else if(strings.HasPrefix(tmpLine, "6") || strings.HasPrefix(tmpLine, "six")) {
            parsedNumbers = append(parsedNumbers, 6)
        } else if(strings.HasPrefix(tmpLine, "7") || strings.HasPrefix(tmpLine, "seven")) {
            parsedNumbers = append(parsedNumbers, 7)
        } else if(strings.HasPrefix(tmpLine, "8") || strings.HasPrefix(tmpLine, "eight")) {
            parsedNumbers = append(parsedNumbers, 8)
        } else if(strings.HasPrefix(tmpLine, "9") || strings.HasPrefix(tmpLine, "nine")) {
            parsedNumbers = append(parsedNumbers, 9)
        }

        tmpLine = tmpLine[1:];

        if(len(tmpLine) == 0) {
            break;
        }
    }

    return parsedNumbers
}
