package romanNumeral

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
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
