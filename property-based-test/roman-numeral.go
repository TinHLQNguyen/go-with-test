package romanNumeral

import "strings"

func ConvertToRoman(arabic uint16) string {
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

func ConvertToArabic(roman string) (total uint16) {
	for _, symbol := range windowedRoman(roman).Symbols() {
		total += allRomanNumerals.ValueOf(symbol...)
	}
	return
}

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbols ...byte) uint16 {
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

// used to make slice of grouped symbols
type windowedRoman string

func (w windowedRoman) Symbols() (symbols [][]byte) {

	for i := 0; i < len(w); i++ {
		symbol := w[i]
		notAtEnd := i+1 < len(w)
		// look ahead to next symbol if we can and, the current symbol is base 10 (or other valid subtractors)
		if notAtEnd && isSubtractive(symbol) && allRomanNumerals.IsExist(symbol, w[i+1]) {
			// get the value of the two-char string if any
			symbols = append(symbols, []byte{symbol, w[i+1]})
			i++ // move past the next char b/c it belong to two-char string
		} else {
			symbols = append(symbols, []byte{symbol})
		}
	}
	return
}

// use byte because in Go, index string yields bytes
func isSubtractive(currentSymbol uint8) bool {
	// use '' for byte (char)
	return currentSymbol == 'I' || currentSymbol == 'X' || currentSymbol == 'C'
}
