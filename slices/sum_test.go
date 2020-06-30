package slices

import "testing"

func TestSum(t *testing.T) {

	numbers := [5]int{2, 2, 3, 4, 5}

	got := Sum(numbers)
	expect := 16

	if got != expect {
		t.Errorf("got %d expect %d", got, expect)
	}
}
