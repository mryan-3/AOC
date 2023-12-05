package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode"
)

func main(){
    fmt.Println("Advent of code day one!!")
    response, _ := http.Get("https://adventofcode.com/2023/day/1/input")

    fileScanner := bufio.NewScanner(strings.NewReader(response))





}

func firstAndlastDigit(line string)(int, int){
    firstDigit, lastDigit := 0, 0

    for _, char := range line {
        if unicode.IsDigit(char){
            digit, _ := strconv.Atoi(string(char))
            if firstDigit == 0{
                firstDigit = digit
            }
            lastDigit = digit
        }
    }
    if firstDigit == 0 {
        firstDigit = lastDigit
    }

    return firstDigit, lastDigit

}
