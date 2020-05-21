package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {

	t.Run("spy sleep test", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &DefaultSleeper{})

		got := buffer.String()
		want := "3\n2\n1\nGo!"
		if got != want {
			t.Errorf("got %q want %q", got, want)
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
func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &Spytime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}

	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}

}
