package main

import (
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
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
}
