package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	//time.Sleep(time.Second)
	time.Sleep(time.Millisecond * 250) // for saving my time
}

func Countdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(w, i)
		//fmt.Println(i)
	}
	s.Sleep()
	fmt.Fprint(w, finalWord)
}

const sleep = "sleep"
const write = "write"

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}
func (s *CountdownOperationsSpy) Write(b []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type Spytime struct {
	durationSlept time.Duration
}

func (s *Spytime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
