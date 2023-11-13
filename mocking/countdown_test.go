package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {

	t.Run("test countdown result", func(t *testing.T) {
		// given
		buffer := &bytes.Buffer{}
		spySlepper := &SpySleeper{}

		CountDown(buffer, spySlepper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if spySlepper.Calls != 3 {
			t.Errorf("not enough calls to sleep")
		}
	})
	t.Run("sleep before next print", func(t *testing.T) {
		// given
		spySleeperPrinter := &SpyCountdownOperationSleeper{}
		// this SpyCountdownOperationSleeper class has both methods for Sleeper interface and io.Writer interface
		CountDown(spySleeperPrinter, spySleeperPrinter)

		got := spySleeperPrinter.Calls
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
