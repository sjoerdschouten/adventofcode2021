package main

import (
	"adventofcode2021/utils"
	"fmt"
)

func main() {
	input, _ := utils.ReadIntegers("./input")
	fmt.Println(solvePartOne(input))
	fmt.Println(solvePartTwo(input))
}

func solvePartOne(input []int) int {
	count := 0
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			count++
		}
		continue
	}
	return count
}

func solvePartTwo(input []int) int {
	count := 0
	for i := 3; i < len(input); i++ {
		window1 := input[i-1] + input[i-2] + input[i-3]
		window2 := input[i] + input[i-1] + input[i-2]
		if window2 > window1 {
			count++
		}
	}
	return count
}
