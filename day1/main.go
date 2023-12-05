package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)
func check(e error){
    if e != nil{
        panic(e)
    }
}

func main(){
    fmt.Println("Advent of code day 1")

    resp, err := os.Open("../day1/input.txt")
    check(err)
    defer resp.Close()

    sum := 0

    scanner := bufio.NewScanner(resp)

    for scanner.Scan(){
        line := scanner.Text()
        firstDigit, lastDigit := firstAndLast(line)
        fmt.Println("Line:", line, "First Digit: ", firstDigit, "Last DIgit: ",lastDigit)

        calibrationValue := combinedValue(firstDigit, lastDigit)
        sum += calibrationValue
        fmt.Println(sum)


        if err := scanner.Err(); err != nil {
            log.Fatal(err)
        }
    }
    fmt.Println("The sum is: ", sum)
}

func firstAndLast(line string)(int, int){
    firstDigit, lastDigit := 0, 0


        for _, char := range line{
            if unicode.IsDigit(char){
                digit, _ := strconv.Atoi(string(char))
                if firstDigit == 0 {
                    firstDigit = digit
                }
                lastDigit = digit
            }
        }
        return firstDigit, lastDigit
}

func combinedValue(firstDigit, lastDigit int)int{
    return firstDigit*10 + lastDigit
}
