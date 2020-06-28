package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Brian")
	want := "Hello, Brian"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
