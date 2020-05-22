package reflections

import (
	"fmt"
	"testing"
)

func TestWalk(t *testing.T) {
	expected := "Chris"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	fmt.Println(x)
	fmt.Println(got)
	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	}
	if got[0] != expected {
		t.Errorf("got %q want %q", got[0], expected)
	}
}
