package main

import (
	"fmt"

	readlines "github.com/tomtomjohansson/advent-2018/readlines"
)

var input, _ = readlines.ReadLines("./textfiles/day2.txt")

func getDiff(line, match string) int {
	diff := 0
	for i := 0; i < len(line); i++ {
		if line[i] != match[i] {
			diff++
		}
	}
	return diff
}

func excludeDifferingLetter(line, match string) string {
	answer := ""
	for i := 0; i < len(line); i++ {
		if line[i] == match[i] {
			answer = answer + string(line[i])
		}
	}
	return answer
}

func main() {
	diff := 0
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			diff = getDiff(input[i], input[j])
			if diff == 1 {
				fmt.Println(excludeDifferingLetter(input[i], input[j]))
				return
			}
		}
	}
}
