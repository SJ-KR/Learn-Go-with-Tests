package main

import "testing"

func TestHello(t *testing.T) {

	assertMassage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		assertMassage(t, got, want)
	})

	t.Run("empty string defaults to 'World''", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertMassage(t, got, want)

	})
}
