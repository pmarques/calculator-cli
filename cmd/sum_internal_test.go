package cmd

import "testing"

func TestSum(t *testing.T) {
	t.Parallel()

	input := []int{1, 2, 3}

	result := sum(input, false)
	if result != 6 {
		t.Fail()
	}
}
