package main

import (
	"adventofcode2021/utils"
	"fmt"
)

func main() {
	input, _ := utils.ReadIntegers("./input")
	fmt.Println(solvePart1(input))
	fmt.Println(solvePart2(input))
}

func solvePart1(input []int) int {
	count := 0
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			count++
		}
		continue
	}
	return count
}

func solvePart2(input []int) int {
	count := 0
	for i := 3; i < len(input); i++ {
		window1 := input[i] + input[i-1] + input[i-2]
		window2 := input[i-1] + input[i-2] + input[i-3]
		if window2 > window1 {
			count++
		}
	}
	return count
}
