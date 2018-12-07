package main

import (
	"fmt"

	readlines "github.com/tomtomjohansson/advent-2018/readlines"
)

var input, _ = readlines.ReadLines("./textfiles/day5.txt")

var line = input[0]

func compareLetters(a, b byte) bool {
	if a^b == 32 {
		return true
	}
	return false
}

func main() {
	for i := 0; i < len(line)-1; i++ {
		if compareLetters(line[i], line[i+1]) {
			line = line[0:i] + line[i+2:len(line)]
			if i == 0 {
				i--
			} else {
				i -= 2
			}
		}
	}
	fmt.Println(len(line))
}
