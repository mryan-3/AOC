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
    sessionCookie := "53616c7465645f5ffd74c5a04945e53f43ad3b2d016e44a59b5f9c062264d186f003da69cabbc5b8b898b81e964c00db5485b8ce26ba9772b86a11a6728d7fd8"

	// Make an HTTP GET request to the URL with the session cookie
	request, _ := http.NewRequest("GET", "https://adventofcode.com/2023/day/1/input", nil)
	request.AddCookie(&http.Cookie{Name: "session", Value: sessionCookie})

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("Error fetching calibration document:", err)
		return
	}
    fmt.Println(response)

    sum := 0
    body, _ := bufio.NewReader(response.Body).ReadString('\n')
    fmt.Println("response:", body)
    fileScanner := bufio.NewScanner(strings.NewReader(body))
    for fileScanner.Scan(){
        line := fileScanner.Text()

        firstDigit, lastDigit := firstAndlastDigit(line)
        calibrationValue := combineDigits(firstDigit, lastDigit)

        sum += calibrationValue
    }

    fmt.Println("The sum is :", sum)
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
func combineDigits(firstDigit, lastDigit int)int{
    return firstDigit*10 + lastDigit
}
