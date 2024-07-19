package main

import (
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("get league from reader", func(t *testing.T) {
		database := strings.NewReader(`[
      {"Name": "Abe", "Wins": 2},
      {"Name": "Ben", "Wins": 10}
      ]`)

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
		database := strings.NewReader(`[
      {"Name": "Abe", "Wins": 2},
      {"Name": "Ben", "Wins": 10}
      ]`)

		store := FileSystemPlayerStore{database}

		got := store.GetPlayerScore("Abe")

		want := 2

		AssertEqual(t, got, want)
	})
}
