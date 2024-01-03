package romanNumeral

import (
	"fmt"
	"testing"
)

var cases = []struct {
	Arabic int
	Roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{14, "XIV"},
	{18, "XVIII"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{100, "C"},
	{90, "XC"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
	{1984, "MCMLXXXIV"},
	{3999, "MMMCMXCIX"},
	{2014, "MMXIV"},
	{1006, "MVI"},
	{798, "DCCXCVIII"},
	{1984, "MCMLXXXIV"},
}

func TestConvertingToRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {

			got := ConvertToRoman(test.Arabic)
			want := test.Roman

			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases[:4] {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {

			got := ConvertToArabic(test.Roman)
			want := test.Arabic

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
