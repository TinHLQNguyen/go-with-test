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
	t.Run("prompt user to enter number of player and start game that ends with Abe win", func(t *testing.T) {
		numberOfPlayer := 7
		winner := "Abe"

		stdOut := &bytes.Buffer{}
		in := strings.NewReader(fmt.Sprintf("%d\n%s wins\n", numberOfPlayer, winner))
		game := &SpyGame{}

		cli := poker.NewCLI(in, stdOut, game)
		cli.PlayPoker()

		got := stdOut.String()
		poker.AssertEqual(t, got, poker.PlayerPrompt)

		if game.StartedWith != numberOfPlayer {
			t.Errorf("wanted game to start with %d, got %d", numberOfPlayer, game.StartedWith)
		}
	})
	t.Run("start 7 players game and Chris win", func(t *testing.T) {
		numberOfPlayer := 7
		winner := "Chris"
		in := strings.NewReader(fmt.Sprintf("%d\n%s wins\n", numberOfPlayer, winner))

		game := &SpyGame{}

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		if game.StartedWith != numberOfPlayer {
			t.Errorf("wanted game to start with %d, got %d", numberOfPlayer, game.StartedWith)
		}

		if game.FinishedWith != winner {
			t.Errorf("wanted game to finish with %s as winner, got %s", winner, game.FinishedWith)
		}
	})
	t.Run("print error when non-numeric player is input", func(t *testing.T) {
		in := strings.NewReader("NotNum\n")

		stdOut := &bytes.Buffer{}
		game := &SpyGame{}

		cli := poker.NewCLI(in, stdOut, game)
		cli.PlayPoker()

		if game.StartCalled {
			t.Error("game should not have started")
		}

		gotPrompt := stdOut.String()
		wantPrompt := poker.PlayerPrompt + poker.BadPlayerInputErrMsg
		poker.AssertEqual(t, gotPrompt, wantPrompt)
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

type SpyGame struct {
	StartedWith  int
	StartCalled  bool
	FinishedWith string
}

func (s *SpyGame) Start(numberOfPlayer int) {
	s.StartedWith = numberOfPlayer
	s.StartCalled = true
}

func (s *SpyGame) Finish(winner string) {
	s.FinishedWith = winner
}
