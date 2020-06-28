package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectGreeting := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Brian", "")
		want := "Hello, Brian"
		assertCorrectGreeting(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Juan", "Spanish")
		want := "Hola, Juan"
		assertCorrectGreeting(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Quasimodo", "French")
		want := "Bonjour, Quasimodo"
		assertCorrectGreeting(t, got, want)
	})

	t.Run("say 'Hello, World' if no name", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectGreeting(t, got, want)
	})

}
