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

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
