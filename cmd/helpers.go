package cmd

import (
	"strconv"
)

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func StringListToIntList(listStrings []string) []int {
	n := len(listStrings)
	numbers := make([]int, n)

	for i, s := range listStrings {
		numbers[i] = StringToInt(s)
	}

	return numbers
}
