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
	t.Run("known word", func(t *testing.T) {
		when, _ := dictionary.Search("test")
		then := "this is a test"

		assertString(t, when, then)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		then := "could not find the word you are looking for"

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertString(t, err.Error(), then)
	})

}
