package slices

import (
	"reflect"
	"testing"
)

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

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2}, []int{0, 9})
	}
}
