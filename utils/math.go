package utils

import (
"fmt"
"strconv"
"strings"
)

// GetInt returns the given string as an int, or panics if it is invalid
func GetInt(in string) int {
	res, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return res
}

func GetMultipleInt(input string) []int {
	integers := strings.Split(input, " ")

	// integer array
	var out []int
	for _, integer := range integers {
		res, err := strconv.Atoi(integer)
		if err != nil {
			fmt.Errorf(err.Error())
			panic(err)
		}
		out = append(out, res)

	}
	return out
}

