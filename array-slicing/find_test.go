package arraySlicing

import (
	"strings"
	"testing"
)

func TestFind(t *testing.T) {
	t.Run("find first even num", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})

		AssertTrue(t, found)
		AssertEqual(t, firstEvenNumber, 2)
	})

	type Person struct {
		Name string
	}

	t.Run("Find a person", func(t *testing.T) {
		people := []Person{
			{Name: "Ape Ben"},
			{Name: "Cant Too"},
			{Name: "Suzuki"},
		}

		king, found := Find(people, func(p Person) bool {
			return strings.Contains(p.Name, "Suzu")
		})

		AssertTrue(t, found)
		AssertEqual(t, king, people[2])
	})
}
