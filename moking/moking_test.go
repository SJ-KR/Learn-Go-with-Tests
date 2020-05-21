package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spy := &SpySleeper{}

	Countdown(buffer, spy)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	if spy.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", spy.Calls)
	}

}
