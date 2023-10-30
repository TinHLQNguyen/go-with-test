package iteration

import (
	"testing"
)

// use given-when-then pattern
func TestRepeatSingle(t *testing.T) {
	given := Repeat("a", 3)
	then := "aaa"

	if given != then {
		t.Errorf("given %q but got %q", given, then)
	}
}
func TestRepeatMultiple(t *testing.T) {
	given := Repeat("ab", 5)
	then := "ababababab"

	if given != then {
		t.Errorf("given %q but got %q", given, then)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}
