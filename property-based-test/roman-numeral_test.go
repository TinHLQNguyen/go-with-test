package romanNumeral

import "testing"

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 get converted to I", 1, "I"},
		{"2 get converted to II", 2, "II"},
		{"3 get converted to III", 3, "III"},
		{"4 get converted to IV (can't repeat more than 3)", 4, "IV"},
		{"5 get converted to V ", 5, "V"},
	}
	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {

			got := ConvertToRoman(test.Arabic)
			want := test.Want

			if got != want {
				t.Errorf("got %q want %q", got, test.Want)
			}
		})
	}
}
