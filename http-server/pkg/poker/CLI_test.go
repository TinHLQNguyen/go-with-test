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
	dummyStdOut := &bytes.Buffer{}
	t.Run("prompt user to enter number of player", func(t *testing.T) {
		stdOut := &bytes.Buffer{}
		in := strings.NewReader("")
		blindAlerter := &SpyBlindAlerter{}
		game := poker.NewGame(dummyPlayerStore, blindAlerter)

		cli := poker.NewCLI(in, stdOut, game)
		cli.PlayPoker()

		got := stdOut.String()
		poker.AssertEqual(t, got, poker.PlayerPrompt)
	})
	t.Run("record Chris win from user input", func(t *testing.T) {
		// we don't care how many players, just use 5
		in := strings.NewReader("5\nChris wins\n")

		playerStore := &poker.StubPlayerStore{}
		dummyAlerter := &SpyBlindAlerter{}
		game := poker.NewGame(playerStore, dummyAlerter)

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})
	t.Run("record Abe win from user input", func(t *testing.T) {
		// we don't care how many players, just use 5
		in := strings.NewReader("5\nAbe wins\n")

		playerStore := &poker.StubPlayerStore{}
		dummyAlerter := &SpyBlindAlerter{}
		game := poker.NewGame(playerStore, dummyAlerter)

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Abe")
	})
}

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
