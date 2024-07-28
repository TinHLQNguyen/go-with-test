package poker

import (
	"os"
	"reflect"
	"testing"
)

// //////////////// mocks and stub
// stub for test, following PlayerStore Interface
type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	leaque   []Player
}

func NewStubPlayerStore(score map[string]int, winCalls []string, leaqueTable []Player) *StubPlayerStore {
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

// func() is destructor for fs db
func CreateTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpFile.Write([]byte(initialData))

	removeFile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}

	return tmpFile, removeFile
}

// /////////////// Asserttion
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

func AssertLeague(t testing.TB, got, want []Player) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GOT %+v, WANT %+v", got, want)
	}
}
