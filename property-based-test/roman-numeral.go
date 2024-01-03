package romanNumeral

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	// https://pkg.go.dev/strings#Builder
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			if arabic >= numeral.Value {
				result.WriteString(numeral.Symbol)
				arabic -= numeral.Value
			}
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) int {
	return 1
}
