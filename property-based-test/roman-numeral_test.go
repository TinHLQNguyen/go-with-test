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
		{"6 get converted to VI ", 6, "VI"},
		{"7 get converted to VII ", 7, "VII"},
		{"8 get converted to VIII ", 8, "VIII"},
		{"9 get converted to IX (can't repeat more than 3)", 9, "IX"},
		{"10 get converted to X ", 10, "X"},
		{"14 gets converted to XIV", 14, "XIV"},
		{"18 gets converted to XVIII", 18, "XVIII"},
		{"20 gets converted to XX", 20, "XX"},
		{"39 gets converted to XXXIX", 39, "XXXIX"},
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
