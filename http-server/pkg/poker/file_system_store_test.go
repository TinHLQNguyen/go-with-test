package poker

import (
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("get sorted league from reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
      {"Name": "Abe", "Wins": 2},
      {"Name": "Ben", "Wins": 10}
      ]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		AssertNoError(t, err)

		got := store.GetLeague()
		want := []Player{
			{"Ben", 10},
			{"Abe", 2},
		}
		assertLeague(t, got, want)

		// 2nd read to make sure can GetLeague multiple times
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
      {"Name": "Abe", "Wins": 2},
      {"Name": "Ben", "Wins": 10}
      ]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		AssertNoError(t, err)

		got, _ := store.GetPlayerScore("Abe")
		want := 2
		AssertEqual(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
      {"Name": "Abe", "Wins": 2},
      {"Name": "Ben", "Wins": 10}
      ]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		AssertNoError(t, err)

		store.RecordWin("Ben")

		got, _ := store.GetPlayerScore("Ben")
		want := 11
		AssertEqual(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
      {"Name": "Abe", "Wins": 2},
      {"Name": "Ben", "Wins": 10}
      ]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		AssertNoError(t, err)

		store.RecordWin("Pepper")

		got, _ := store.GetPlayerScore("Pepper")
		want := 1
		AssertEqual(t, got, want)
	})

	t.Run("work with empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayerStore(database)

		AssertNoError(t, err)
	})
}

// func() is destructor
func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
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
