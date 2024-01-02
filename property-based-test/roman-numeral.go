package romanNumeral

import "strings"

func ConvertToRoman(arabic int) string {
	// https://pkg.go.dev/strings#Builder
	var result strings.Builder

	for i := 0; i < arabic; i++ {
		result.WriteString("I")
	}

	return result.String()
}
