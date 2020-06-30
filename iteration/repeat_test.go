package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("y", 5)
	expect := "yyyyy"

	if got != expect {
		t.Errorf("got %q expect %q", got, expect)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("m", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("s", 4)
	fmt.Println(repeated)
	// Output: ssss
}
