package arraySlicing

func Reduce[A, B any](collection []A, accumulator func(B, A) B, initValue B) B {
	result := initValue
	for _, x := range collection {
		result = accumulator(result, x)
	}
	return result
}
