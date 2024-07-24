package poker

import "testing"

// stub for test, following PlayerStore Interface
type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	leaque   []Player
}

func newStubPlayerStore(score map[string]int, winCalls []string, leaqueTable []Player) *StubPlayerStore {
	return &StubPlayerStore{score, winCalls, leaqueTable}
}

func (s *StubPlayerStore) GetPlayerScore(name string) (int, bool) {
	score, ok := s.scores[name]
	return score, ok
}

// spy on POST calls
func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.leaque
}

// comparable type parameter indicates that we only accept things that are comparable
func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("GOT %v, WANT %v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("GOT %v, NOT WANT %v", got, want)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("GOT %v, WANT true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("GOT %v, WANT false", got)
	}
}

func AssertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didnt expect error but got one, %v", err)
	}
}

func AssertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()
	if len(store.winCalls) != 1 {
		t.Errorf("got %d calls to RecordWin, want %d", len(store.winCalls), 1)
	}

	// assert if correct name is recorded
	if store.winCalls[0] != winner {
		t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], winner)
	}
}
