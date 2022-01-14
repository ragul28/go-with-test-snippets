package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMsg := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Hello to people", func(t *testing.T) {
		got := Hello("Tony", "")
		want := "Hello, Tony"
		assertCorrectMsg(t, got, want)
	})

	t.Run("Hello to world, empty input", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMsg(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMsg(t, got, want)
	})
}
