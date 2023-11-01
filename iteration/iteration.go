package iteration

import "strings"

const repeatCount = 5

func Repeat(character string, count int) string {
	var repeated string // should only initialize with string's "zero" value
	repeated = strings.Repeat(character, count)
	return repeated
}
