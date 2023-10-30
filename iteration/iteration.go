package iteration

const repeatCount = 5

func Repeat(character string) string {
	var repeated string // only initialize with string's "zero" value
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
