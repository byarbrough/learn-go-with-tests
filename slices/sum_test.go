package slices

import (
	"reflect"
	"testing"
)

// assert that slices are equal
func compareSlices(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSum(t *testing.T) {

	numbers := []int{2, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 16

	if got != want {
		t.Errorf("got %d expect %d given %v", got, want, numbers)
	}

}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	compareSlices(t, got, want)
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2}, []int{0, 9})
	}
}

func TestSumAllTails(t *testing.T) {

	t.Run("test sum of tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{5, 9, 2})
		want := []int{2, 11}

		compareSlices(t, got, want)
	})

	t.Run("test empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 11, 99})
		want := []int{0, 110}

		compareSlices(t, got, want)
	})

}
