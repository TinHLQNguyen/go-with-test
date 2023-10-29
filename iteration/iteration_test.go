package iteration

import (
	"testing"
)

// use given-when-then pattern
func TestRepeat(t *testing.T) {
	given := Repeat("a")
	then := "aaaaa"

	if given != then {
		t.Errorf("given %q but got %q", given, then)
	}
}
