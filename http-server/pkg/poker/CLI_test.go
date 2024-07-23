package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	in := strings.NewReader("Chris wins\n")

	playerStore := &StubPlayerStore{}
	cli := &CLI{playerStore, in}
	cli.PlayPoker()

	if len(playerStore.winCalls) != 1 {
		t.Fatalf("expect a wincall but got none")
	}

	got := playerStore.winCalls[0]
	want := "Chris"

	AssertEqual(t, got, want)
}
