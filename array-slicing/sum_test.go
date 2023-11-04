package main

import (
	"testing"
)

// use given-when-then framework
func TestSum(t *testing.T) {
	t.Run("collection of some slice numbers", func(t *testing.T) {
		given := []int{1, 2, 3, 4, 5}

		when := Sum(given)
		then := 15

		if when != then {
			t.Errorf("when %d then %d, given %v", when, then, given)
		}
	})
	t.Run("collection of any slice numbers", func(t *testing.T) {
		given := []int{1, 2, 3}

		when := Sum(given)
		then := 6

		if when != then {
			t.Errorf("when %d then %d, given %v", when, then, given)
		}
	})
}

func TestSumAll(t *testing.T) {
	when := SumAll([]int{1, 2, 4}, []int{0, 9})
	then := []int{7, 9}

	if when != then {
		t.Errorf("when %v , then %d", when, then)
	}
}
