package main

import "testing"

func assertString(t testing.TB, when, then string) {
	t.Helper()

	if when != then {
		t.Errorf("when %q then %q", when, then)
	}
}

func assertError(t testing.TB, when, then error) {
	t.Helper()
	if when == nil {
		t.Fatal("expected to get an error.")
	}
	if when != then {
		t.Errorf("when got error %q then %q", when, then)
	}
}

func TestSearch(t *testing.T) {
	// given
	dictionary := Dictionary{"test": "this is a test"}
	t.Run("known word", func(t *testing.T) {
		when, _ := dictionary.Search("test")
		then := "this is a test"

		assertString(t, when, then)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, when := dictionary.Search("unknown")
		then := ErrNotFound

		assertError(t, when, then)
	})

}
