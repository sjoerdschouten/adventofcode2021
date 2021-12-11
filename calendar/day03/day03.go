package main

import (
	"adventofcode2021/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("The answer to puzzle one is: %d\n", solvePuzzleOne())
}

func solvePuzzleOne() int64 {
	lines, _ := utils.ReadLines("./input")
	bitLength := len(lines[0])
	verticalBits := make([]string, bitLength)

	for _, line := range lines {
		for i, bits := range line {
			bitCharacters := fmt.Sprintf("%c", bits)
			verticalBits[i] = verticalBits[i] + bitCharacters
		}
	}
	gammaRate := findGammaRate(verticalBits)
	epsilonRate := findEpsilonRate(verticalBits)

	gammaRateDecimal, _ := strconv.ParseInt(gammaRate, 2, 64)
	epsilonRateDecimal, _ := strconv.ParseInt(epsilonRate, 2, 64)

	return gammaRateDecimal * epsilonRateDecimal
}

func findGammaRate(binaryDigits []string) string {
	var result string
	for _, value := range binaryDigits {
		ones := strings.Count(value, "1")
		zeros := strings.Count(value, "0")
		if ones < zeros {
			result += "0"
		} else {
			result += "1"
		}
	}
	return result
}

func findEpsilonRate(binaryDigits []string) string {
	var result string
	for _, value := range binaryDigits {
		ones := strings.Count(value, "1")
		zeros := strings.Count(value, "0")
		if ones > zeros {
			result += "0"
		} else {
			result += "1"
		}
	}
	return result
}
