package cmd

import "testing"

func TestSub(t *testing.T) {
	t.Parallel()

	input := []int{1, 2, 3}

	result := sub(input, false)
	if result != -4 {
		t.Fail()
	}
}
