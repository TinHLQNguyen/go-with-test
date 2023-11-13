package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
	sleep          = "sleep"
	write          = "write"
)

// make a Sleeper interface using Sleep() to add dependency to the behavior
type Sleeper interface {
	Sleep()
}

// This is the mock Sleeper we use our tests to hasten it
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

// This is the mock Sleeper to test operation of write and sleep
type SpyCountdownOperationSleeper struct {
	Calls []string
}

func (s *SpyCountdownOperationSleeper) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// this function will match Writer interface, which can be called by io.Writer
func (s *SpyCountdownOperationSleeper) Writer(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

// This is the real Sleeper used in the application
type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func CountDown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	CountDown(os.Stdout, sleeper)
}
