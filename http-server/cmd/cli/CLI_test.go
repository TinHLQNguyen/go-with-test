package main

import (
	"go-with-test/http-server/pkg/poker"
	"strings"
	"testing"
	"time"
)

func TestCLI(t *testing.T) {
	t.Run("record Chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")

		playerStore := &poker.StubPlayerStore{}
		dummyAlerter := &SpyBlindAlerter{}

		cli := NewCLI(playerStore, in, dummyAlerter)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})
	t.Run("record Abe win from user input", func(t *testing.T) {
		in := strings.NewReader("Abe wins\n")

		playerStore := &poker.StubPlayerStore{}
		dummyAlerter := &SpyBlindAlerter{}

		cli := NewCLI(playerStore, in, dummyAlerter)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Abe")
	})
	t.Run("schedule printing of blind value", func(t *testing.T) {
		in := strings.NewReader("Abe wins\n")

		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := NewCLI(playerStore, in, blindAlerter)
		cli.PlayPoker()

		if len(blindAlerter.alerts) != 1 {
			t.Fatal("expected a blind alert to be scheduled")
		}
	})
}

type SpyBlindAlerter struct {
	alerts []struct {
		scheduledAt time.Duration
		amount      int
	}
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, struct {
		scheduledAt time.Duration
		amount      int
	}{duration, amount})
}
