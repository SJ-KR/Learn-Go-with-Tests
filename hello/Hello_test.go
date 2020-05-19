package main

import "testing"

func TestHello(t *testing.T) {

	assertMassage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertMassage(t, got, want)

	})
	t.Run("in English", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"

		assertMassage(t, got, want)

	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Codd", "French")
		want := "Bonjour, Codd"

		assertMassage(t, got, want)

	})

}
