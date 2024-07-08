package generics

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})
	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "Tin")
	})
	// AssertEqual(t, 1, "1") // This should not be allowed
}

func AssertEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("GOT %d, WANT %d", got, want)
	}
}

func AssertNotEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if got == want {
		t.Errorf("GOT %d, NOT WANT %d", got, want)
	}
}
