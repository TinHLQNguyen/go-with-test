package main

import "testing"

// use given-when-then framework
func TestSum(t *testing.T) {
	given := [5]int{1, 2, 3, 4, 5}

	when := Sum(given)
	then := 15

	if when != then {
		t.Errorf("when %d then %d, given %v", when, then, given)
	}
}
