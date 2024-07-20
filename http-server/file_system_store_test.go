package main

import (
	"io"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("get league from reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
      {"Name": "Abe", "Wins": 2},
      {"Name": "Ben", "Wins": 10}
      ]`)
		defer cleanDatabase()

		store := FileSystemPlayerStore{database}

		got := store.GetLeague()
		want := []Player{
			{"Abe", 2},
			{"Ben", 10},
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

		store := FileSystemPlayerStore{database}

		got := store.GetPlayerScore("Abe")
		want := 2
		AssertEqual(t, got, want)
	})
}

// func() is destructor
func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
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
