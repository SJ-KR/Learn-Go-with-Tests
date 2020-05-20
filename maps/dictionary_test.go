package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		_, got := dictionary.Search("test")
		want := error(nil)

		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		if got == nil {
			t.Fatal("expected to get an error")
		}
		assertStrings(t, got, ErrNotFound)
	})

}
func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	word := "test"
	definition := "this is just a test"

	dictionary.Add(word, definition)

	assertDefinition(t, dictionary, word, definition)

}
func assertStrings(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if got != definition {
		t.Errorf("got %q, want %q", got, definition)
	}
}
