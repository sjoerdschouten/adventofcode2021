package main

import (
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	got := solvePuzzleOne()
	want := int64(198)

	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}
