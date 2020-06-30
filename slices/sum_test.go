package slices

import "testing"

func TestSum(t *testing.T) {

	numbers := []int{2, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 16

	if got != want {
		t.Errorf("got %d expect %d given %v", got, want, numbers)
	}

}
