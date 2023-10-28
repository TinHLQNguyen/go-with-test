package integers

import "testing"

// use given-when-then pattern
func TestAdder(t *testing.T) {
	when := Add(2, 2)
	then := 4

	if when != then {
		t.Errorf("Expect '%d' but got '%d'", when, then)
	}
}
