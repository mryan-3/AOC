package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
	//	"strings"
	//  "strconv"
)
func check(e error){
    if e != nil{
        panic(e)
    }
}
var numberMap = map[string]string{
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9",
}

func main(){
    fmt.Println("Advent of code day 1")

    resp, err := os.Open("../day1/input.txt")
    check(err)
    defer resp.Close()

    sum := 0

    scanner := bufio.NewScanner(resp)
    fmt.Println(resp)

    for scanner.Scan(){
        line := scanner.Text()

        for word, digit := range numberMap{
            line = strings.Replace(line, word, digit, -1)
        }
       // fmt.Println(line)
        firstDigit, lastDigit := firstAndlastDigit(line)
       // fmt.Println(firstDigit, lastDigit)
        calibratedValue := combinedDigits(firstDigit, lastDigit)
       // fmt.Println(calibratedValue)

        sum = sum + calibratedValue
       // fmt.Println(sum)

        if err := scanner.Err(); err != nil {
            log.Fatal(err)
        }
    }
 //  fmt.Println("The sum is: ", sum)
}

func firstAndlastDigit(line string)(int, int){
    firstDIgit, lastDigit := 0, 0
    for _, char := range line{
        if unicode.IsDigit(char){
            digit, _ := strconv.Atoi(string(char))
            if firstDIgit == 0 {
                firstDIgit= digit
            }
            lastDigit = digit
        }
    }
    return firstDIgit, lastDigit
}

func combinedDigits(firstDIgit, lastDigit int)int{
    return firstDIgit*10 + lastDigit
}
