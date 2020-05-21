package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {

	t.Run("spy sleep test", func(t *testing.T) {
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
	})

	t.Run("sleep before every print", func(t *testing.T) {
		SpySleepPrinter := &CountdownOperationsSpy{}
		Countdown(SpySleepPrinter, SpySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(SpySleepPrinter.Calls, want) {
			t.Errorf("wanted calls %v got %v", want, SpySleepPrinter.Calls)
		}
	})

}
