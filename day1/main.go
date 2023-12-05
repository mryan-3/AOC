package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
    "strconv"
)
func check(e error){
    if e != nil{
        panic(e)
    }
}
var numberMap = map[string]int{
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
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

        var nums []int
        for word, num := range numberMap {
            if strings.Contains(line, word) {
                line = strings.Replace(line, word, "", -1)
                nums = append(nums, num)
            }
        }
        for _, char := range line {
            if char >= '0' && char <= '9'{
                num := int(char - '0')
                nums = append(nums, num)
            }
        }
        if len(nums) > 0 {
			firstLast, _ := strconv.Atoi(fmt.Sprintf("%d%d", nums[0], nums[len(nums)-1]))
			sum += firstLast
		}


        if err := scanner.Err(); err != nil {
            log.Fatal(err)
        }
    }
   fmt.Println("The sum is: ", sum)
}
