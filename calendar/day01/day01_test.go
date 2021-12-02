package main

import (
	"adventofcode2021/utils"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	input, _ := utils.ReadIntegers("./test")
	got := solvePartOne(input)
	want := 7

	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}

func TestSolvePartTwo(t *testing.T) {
	input, _ := utils.ReadIntegers("./test")
	got := solvePartTwo(input)
	want := 5

	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}
