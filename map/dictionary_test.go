package main

import "testing"

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

func TestAdd(t *testing.T) {
	t.Run("add new word", func(t *testing.T) {
		// given
		dictionary := Dictionary{}
		word := "test"
		definition := "this is a test"

		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("add existing word", func(t *testing.T) {
		// given
		word := "test"
		definition := "this is a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, definition)

		assertError(t, err, ErrWordExist)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing word", func(t *testing.T) {
		// given
		word := "test"
		definition := "this is a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		dictionary.Update(word, newDefinition)

		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("update new word", func(t *testing.T) {
		// given
		word := "test"
		definition := "this is a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

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

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	when, err := dictionary.Search(word)
	then := definition
	if err != nil {
		t.Fatal("Should find added word", err)
	}

	assertString(t, when, then)
}
