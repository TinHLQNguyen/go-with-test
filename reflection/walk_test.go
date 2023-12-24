package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}
type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Christ"},
			[]string{"Christ"},
		},
		{
			"struct with two string field",
			struct {
				Name string
				City string
			}{"Christ", "Tokyo"},
			[]string{"Christ", "Tokyo"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Christ", 30},
			[]string{"Christ"},
		},
		{
			"struct nested fields",
			Person{
				"Christ",
				Profile{30, "Tokyo"},
			},
			[]string{"Christ", "Tokyo"},
		},
		{
			"pointers to things",
			&Person{
				"Christ",
				Profile{30, "Tokyo"},
			},
			[]string{"Christ", "Tokyo"},
		},
		{
			"slices",
			[]Profile{
				{30, "Tokyo"},
				{20, "Vietnam"},
			},
			[]string{"Tokyo", "Vietnam"},
		},
		{
			"array",
			[2]Profile{
				{30, "Tokyo"},
				{20, "Vietnam"},
			},
			[]string{"Tokyo", "Vietnam"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
