package main

import (
	"adventofcode2021/utils"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	input, _ := utils.ReadLines("./test")
	got := solvePartOne(input)
	want := 150

	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}

func TestSolvePartTwo(t *testing.T) {
	input, _ := utils.ReadLines("./test")
	got := solvePartTwo(input)
	want := 900

	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}
