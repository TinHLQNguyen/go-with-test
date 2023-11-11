package main

import "testing"

func assertString(t testing.TB, when, then string) {
	t.Helper()

	if when != then {
		t.Errorf("when %q then %q", when, then)
	}
}

func TestSearch(t *testing.T) {
	// given
	dictionary := Dictionary{"test": "this is a test"}

	when := dictionary.Search("test")
	then := "this is a test"

	assertString(t, when, then)
}
