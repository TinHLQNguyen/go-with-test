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
				age  int
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
}
