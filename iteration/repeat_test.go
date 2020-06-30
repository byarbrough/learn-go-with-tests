package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("y")
	expect := "yyyyy"

	if got != expect {
		t.Errorf("got %q expect %q", got, expect)
	}
}
