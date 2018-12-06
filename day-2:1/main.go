package main

import (
	"fmt"
	"strings"

	readlines "github.com/tomtomjohansson/advent-2018/readlines"
)

var input, _ = readlines.ReadLines("./textfiles/day2.txt")

func main() {
	var twos, threes int
	for i := 0; i < len(input); i++ {
		line := input[i]
		used := map[string]bool{}
		twoCount, threeCount := false, false
		for j := 0; j < len(line); j++ {
			char := string(line[j])
			if used[char] {
				continue
			}
			used[char] = true
			count := strings.Count(line, char)
			if count == 2 && !twoCount {
				twoCount = true
				twos++
			} else if count == 3 && !threeCount {
				threeCount = true
				threes++
			}
		}
	}
	fmt.Println(twos * threes)
}
