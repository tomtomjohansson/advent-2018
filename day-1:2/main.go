package main

import "fmt"
import "strconv"

import readlines "github.com/tomtomjohansson/advent-2018/readlines"

var sum, finalAnswer, answer int64
var seen = make(map[int64]bool)
var input, _ = readlines.ReadLines("./textfiles/day1.txt")

func loopThroughInput() int64 {
	for i := 0; i < len(input); i++ {
		number, _ := strconv.ParseInt(input[i], 10, 64)
		sum += number
		if seen[sum] {
			finalAnswer = sum
			break
		}
		seen[sum] = true
	}
	return finalAnswer
}

func main() {
	for finalAnswer == 0 {
		finalAnswer = loopThroughInput()
	}
	fmt.Println(finalAnswer)
}
