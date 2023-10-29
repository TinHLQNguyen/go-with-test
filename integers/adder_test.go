package integers

import (
	"fmt"
	"testing"
)

// use given-when-then pattern
func TestAdder(t *testing.T) {
	when := Add(2, 2)
	then := 4

	if when != then {
		t.Errorf("Expect '%d' but got '%d'", when, then)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
