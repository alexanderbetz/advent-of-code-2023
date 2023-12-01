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
    var sum int = 0

    for _, line := range lines {
        numbers, err := getNumbers(line)
        if(err != nil) {
            log.Fatal(err)
        }
        if(len(numbers) > 0) {
            sum += (numbers[0] * 10) + numbers[len(numbers)-1]
        } 
    }
    fmt.Printf("Sum of all calibration values: %d\n", sum)
}

func getNumbers(line string) ([]int, error) {
    numbers := regexp.MustCompile(`\d`).FindAllString(line, -1)
    parsedNumbers := make([]int, len(numbers))
    for i, char := range numbers {
        n , err := strconv.Atoi(char)
        if(err != nil) {
            fmt.Println(err)
        }
        parsedNumbers[i] = int(n)
    }
    return parsedNumbers, nil
}
