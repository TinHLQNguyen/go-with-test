package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {

	t.Run("test countdown result", func(t *testing.T) {
		// given
		buffer := &bytes.Buffer{}
		spySlepper := &SpyCountdownOperationSleeper{}

		CountDown(buffer, spySlepper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
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

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("Should have slept for %q but slept for %q", sleepTime, spyTime.durationSlept)
	}
}
