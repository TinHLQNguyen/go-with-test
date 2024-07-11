package arraySlicing

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	add := func(acc, x int) int { return acc + x }
	return Reduce(numbers, add, 0)
}

func SumAll(numbersToSum ...[]int) []int {
	sumAll := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			return append(acc, Sum(x))
		}
	}
	return Reduce(numbersToSum, sumAll, []int{})
}

func SumAllTails(numbersToSumTails ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}
	return Reduce(numbersToSumTails, sumTail, []int{})
}

func Reduce[A any](collection []A, accumulator func(A, A) A, initValue A) A {
	result := initValue
	for _, x := range collection {
		result = accumulator(result, x)
	}
	return result
}
