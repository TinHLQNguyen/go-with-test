package poker_test

import (
	"go-with-test/http-server/pkg/poker"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("get sorted league from reader", func(t *testing.T) {
		database, cleanDatabase := poker.CreateTempFile(t, `[
      {"Name": "Abe", "Wins": 2},
      {"Name": "Ben", "Wins": 10}
      ]`)
		defer cleanDatabase()

		store, err := poker.NewFileSystemPlayerStore(database)
		poker.AssertNoError(t, err)

		got := store.GetLeague()
		want := []poker.Player{
			{"Ben", 10},
			{"Abe", 2},
		}
		poker.AssertLeague(t, got, want)

		// 2nd read to make sure can GetLeague multiple times
		got = store.GetLeague()
		poker.AssertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := poker.CreateTempFile(t, `[
      {"Name": "Abe", "Wins": 2},
      {"Name": "Ben", "Wins": 10}
      ]`)
		defer cleanDatabase()

		store, err := poker.NewFileSystemPlayerStore(database)
		poker.AssertNoError(t, err)

		got, _ := store.GetPlayerScore("Abe")
		want := 2
		poker.AssertEqual(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := poker.CreateTempFile(t, `[
      {"Name": "Abe", "Wins": 2},
      {"Name": "Ben", "Wins": 10}
      ]`)
		defer cleanDatabase()

		store, err := poker.NewFileSystemPlayerStore(database)
		poker.AssertNoError(t, err)

		store.RecordWin("Ben")

		got, _ := store.GetPlayerScore("Ben")
		want := 11
		poker.AssertEqual(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := poker.CreateTempFile(t, `[
      {"Name": "Abe", "Wins": 2},
      {"Name": "Ben", "Wins": 10}
      ]`)
		defer cleanDatabase()

		store, err := poker.NewFileSystemPlayerStore(database)
		poker.AssertNoError(t, err)

		store.RecordWin("Pepper")

		got, _ := store.GetPlayerScore("Pepper")
		want := 1
		poker.AssertEqual(t, got, want)
	})

	t.Run("work with empty file", func(t *testing.T) {
		database, cleanDatabase := poker.CreateTempFile(t, "")
		defer cleanDatabase()

		_, err := poker.NewFileSystemPlayerStore(database)

		poker.AssertNoError(t, err)
	})
}
