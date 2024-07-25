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
	t.Run("prompt user to enter number of player and start game", func(t *testing.T) {
		numberOfPlayer := 7

		stdOut := &bytes.Buffer{}
		in := strings.NewReader(fmt.Sprintf("%d\n", numberOfPlayer))
		game := &SpyGame{}

		cli := poker.NewCLI(in, stdOut, game)
		cli.PlayPoker()

		got := stdOut.String()
		poker.AssertEqual(t, got, poker.PlayerPrompt)

		if game.StartedWith != numberOfPlayer {
			t.Errorf("wanted game to start with %d, got %d", numberOfPlayer, game.StartedWith)
		}
	})
	t.Run("record Chris win from user input", func(t *testing.T) {
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
	FinishedWith string
}

func (s *SpyGame) Start(numberOfPlayer int) {
	s.StartedWith = numberOfPlayer
}

func (s *SpyGame) Finish(winner string) {
	s.FinishedWith = winner
}
