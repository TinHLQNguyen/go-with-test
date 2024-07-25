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
		in := userSends("7", "Abe wins")

		stdOut := &bytes.Buffer{}
		game := &SpyGame{}

		cli := poker.NewCLI(in, stdOut, game)
		cli.PlayPoker()

		assertMessageSentToUser(t, stdOut, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 7)
		assertGameFinishedWith(t, game, "Abe")
	})
	t.Run("start 7 players game and Chris win", func(t *testing.T) {
		in := userSends("3", "Chris wins")

		game := &SpyGame{}

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		assertGameStartedWith(t, game, 3)
		assertGameFinishedWith(t, game, "Chris")
	})
	t.Run("print error when non-numeric player is input", func(t *testing.T) {
		in := userSends("NotInt")

		stdOut := &bytes.Buffer{}
		game := &SpyGame{}

		cli := poker.NewCLI(in, stdOut, game)
		cli.PlayPoker()

		assertGameNotStart(t, game)
		assertMessageSentToUser(t, stdOut, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})
	t.Run("print error when receive bad format winner command", func(t *testing.T) {
		in := userSends("3", "Abe lose")

		stdOut := &bytes.Buffer{}
		game := &SpyGame{}

		cli := poker.NewCLI(in, stdOut, game)
		cli.PlayPoker()

		assertGameNotFinish(t, game)
		assertGameStartedWith(t, game, 3)
		assertMessageSentToUser(t, stdOut, poker.PlayerPrompt, poker.BadWinnerInputErrMsg)
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
	FinishCalled bool
}

func (s *SpyGame) Start(numberOfPlayer int) {
	s.StartedWith = numberOfPlayer
	s.StartCalled = true
}

func (s *SpyGame) Finish(winner string) {
	s.FinishedWith = winner
}

func userSends(inputs ...string) *strings.Reader {
	return strings.NewReader(strings.Join(inputs, "\n"))
}

func assertGameStartedWith(t testing.TB, game *SpyGame, numberOfPlayer int) {
	if game.StartedWith != numberOfPlayer {
		t.Errorf("wanted game to start with %d, got %d", numberOfPlayer, game.StartedWith)
	}
}

func assertGameFinishedWith(t testing.TB, game *SpyGame, winner string) {
	if game.FinishedWith != winner {
		t.Errorf("wanted game to finish with %s as winner, got %s", winner, game.FinishedWith)
	}
}

func assertGameNotStart(t testing.TB, game *SpyGame) {
	if game.StartCalled {
		t.Error("game should not have started")
	}
}

func assertGameNotFinish(t testing.TB, game *SpyGame) {
	if game.FinishCalled {
		t.Error("game should not have finished")
	}
}

func assertMessageSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("wanted %s to be sent to stdOut but got %+v", want, messages)
	}
}
