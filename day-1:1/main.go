package main

import "fmt"
import "strconv"
import readlines "github.com/tomtomjohansson/advent-2018/readlines"

var input, _ = readlines.ReadLines("./textfiles/day1.txt")

func main() {
	var sum int64
	for i := 0; i < len(input); i++ {
		number, _ := strconv.ParseInt(input[i], 10, 64)
		sum += number
	}
	fmt.Println(sum)
}
