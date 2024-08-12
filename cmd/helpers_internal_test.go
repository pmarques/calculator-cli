package cmd

import (
	"slices"
	"testing"
)

func TestStringToInt(t *testing.T) {
	t.Parallel()

	input := "1"
	result := StringToInt(input)

	if result != 1 {
		t.Fail()
	}
}

func TestStringListToIntList(t *testing.T) {
	t.Parallel()

	input := []string{"1", "2"}
	result := StringListToIntList(input)

	if !slices.Equal(result, []int{1, 2}) {
		t.Fail()
	}
}
