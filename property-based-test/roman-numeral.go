package romanNumeral

import "strings"

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
	total := 0

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]
		notAtEnd := i+1 < len(roman)
		// look ahead to next symbol if we can and, the current symbol is base 10 (or other valid subtractors)
		if notAtEnd && isSubtractiveSymbole(symbol, roman) && allRomanNumerals.IsExist(symbol, roman[i+1]) {
			// get the value of the two-char string if any
			value := allRomanNumerals.ValueOf(symbol, roman[i+1])
			total += value
			i++ // move past the next char b/c it belong to two-char string
		} else {
			total += allRomanNumerals.ValueOf(symbol)
		}
	}

	return total
}

type RomanNumeral struct {
	Value  int
	Symbol string
}

type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

func (r RomanNumerals) IsExist(symbols ...byte) bool {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return true
		}
	}

	return false
}

var allRomanNumerals = RomanNumerals{
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

// use byte because in Go, index string yields bytes
func isSubtractiveSymbole(currentSymbol byte, roman string) bool {
	// use '' for byte (char)
	return currentSymbol == 'I' || currentSymbol == 'X' || currentSymbol == 'C'
}
