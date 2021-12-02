package main

import (
	"adventofcode2021/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input, _ := utils.ReadLines("./..input")
	fmt.Printf("Part one: %d\n", solvePartOne(input))
	fmt.Printf("Part two: %d\n", solvePartTwo(input))
}

func solvePartOne(input []string) int {
	y := 0 // depth
	x := 0
	for _, s := range input {
		command := strings.Split(s, " ")
		direction := command[0]
		distance, _ := strconv.Atoi(command[1])

		switch direction {
		case "up":
			y -= distance
		case "down":
			y += distance
		case "forward":
			x += distance
		case "backward":
			x -= distance
		}
	}
	return y * x
}

func solvePartTwo(input []string) int {
	y := 0 // depth
	x := 0
	aim := 0
	for _, s := range input {
		command := strings.Split(s, " ")
		direction := command[0]
		distance, _ := strconv.Atoi(command[1])

		switch direction {
		case "up":
			aim -= distance
		case "down":
			aim += distance
		case "forward":
			x += distance
			y += aim * distance
		case "backward":
			x -= distance
		}
	}
	return x * y
}
