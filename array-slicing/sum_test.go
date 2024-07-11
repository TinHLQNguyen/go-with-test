package arraySlicing

import (
	"reflect"
	"testing"
)

// use given-when-then framework
func TestSum(t *testing.T) {
	assertCorrectSum := func(t testing.TB, when, then int, given []int) {
		t.Helper()
		if !reflect.DeepEqual(when, then) {
			t.Errorf("when %d then %d, given %v", when, then, given)
		}
	}
	t.Run("collection of some slice numbers", func(t *testing.T) {
		given := []int{1, 2, 3, 4, 5}

		when := Sum(given)
		then := 15

		assertCorrectSum(t, when, then, given)
	})
	t.Run("collection of any slice numbers", func(t *testing.T) {
		given := []int{1, 2, 3}

		when := Sum(given)
		then := 6

		assertCorrectSum(t, when, then, given)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("make sums of some slice", func(t *testing.T) {
		when := SumAll([]int{1, 2, 4}, []int{0, 9})
		then := []int{7, 9}

		assertCorrectSlice(t, when, then)
	})
	t.Run("make sums of some empty slice", func(t *testing.T) {
		when := SumAll([]int{}, []int{0, 9})
		then := []int{0, 9}

		assertCorrectSlice(t, when, then)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("make tail sums of some slice", func(t *testing.T) {
		when := SumAllTails([]int{1, 2, 4}, []int{0, 9})
		then := []int{6, 9}

		assertCorrectSlice(t, when, then)
	})
	t.Run("make tail sums of empty slice", func(t *testing.T) {
		when := SumAllTails([]int{}, []int{0, 9})
		then := []int{0, 9}

		assertCorrectSlice(t, when, then)
	})
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}
		AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concat := func(a, b string) string {
			return a + b
		}
		AssertEqual(t, Reduce([]string{"a", "b", "c"}, concat, ""), "abc")
	})
}

func assertCorrectSlice(t testing.TB, when, then []int) {
	t.Helper()
	if !reflect.DeepEqual(when, then) {
		t.Errorf("when %v , then %v", when, then)
	}
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("GOT %v, WANT %v", got, want)
	}
}
