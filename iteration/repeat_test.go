package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("y")
	expect := "yyyyy"

	if got != expect {
		t.Errorf("got %q expect %q", got, expect)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("m")
	}
}
