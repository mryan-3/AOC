package main

import (
	"bufio"
	"fmt"
	"net/http"
    "regexp"
    "strconv"
)

func main(){
    fmt.Println("Advent of code day one!!")
    response, _ := http.Get("https://adventofcode.com/2023/day/1/input")

    fileScanner := bufio.NewScanner(response.Body)
    sum := 0

    for fileScanner.Scan(){
        line := fileScanner.Text()
        // Extract first and last digits
        match := regexp.MustCompile(`(\d).*(\d$)`).FindStringSubmatch(line)
        if len(match) > 1 {
            firstDigit, lastDigit := match[1], match[2]
            digits := firstDigit + lastDigit
            value, _ := strconv.Atoi(digits)
            sum += value
        }

    }
    fmt.Println(sum)

}
