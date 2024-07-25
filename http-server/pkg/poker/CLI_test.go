package poker_test

import (
	"bytes"
	"fmt"
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

		cli := poker.NewCLI(playerStore, in, dummyStdOut, dummyAlerter)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})
	t.Run("record Abe win from user input", func(t *testing.T) {
		in := strings.NewReader("Abe wins\n")

		playerStore := &poker.StubPlayerStore{}
		dummyAlerter := &SpyBlindAlerter{}

		cli := poker.NewCLI(playerStore, in, dummyStdOut, dummyAlerter)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Abe")
	})
	t.Run("schedule printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Abe wins\n")

		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := poker.NewCLI(playerStore, in, dummyStdOut, blindAlerter)
		cli.PlayPoker()

		cases := []scheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
				}

				got := blindAlerter.alerts[i]
				assertScheduledAlert(t, got, want)
			})
		}
	})

	t.Run("prompt user to enter number of player", func(t *testing.T) {
		stdOut := &bytes.Buffer{}
		cli := poker.NewCLI(dummyPlayerStore, dummyStdIn, stdOut, dummyBlindAlerter)
		cli.PlayPoker()

		got := stdOut.String()

		poker.AssertEqual(t, got, poker.PlayerPrompt)
	})
}

var (
	dummyBlindAlerter = &SpyBlindAlerter{}
	dummyPlayerStore  = &poker.StubPlayerStore{}
	dummyStdIn        = &bytes.Buffer{}
	dummyStdOut       = &bytes.Buffer{}
)

type scheduledAlert struct {
	at     time.Duration
	amount int
}

// This is Stringer interface
func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{duration, amount})
}

func assertScheduledAlert(t testing.TB, got, want scheduledAlert) {
	t.Helper()
	if got != want {
		t.Errorf("got alert %v, want alert %v", got, want)
	}
}
