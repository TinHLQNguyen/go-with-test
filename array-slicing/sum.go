package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	// This will allocate sums with [0,0,0,0...] that has enough capacity
	sums := make([]int, lengthOfNumbers)

	for index, numbers := range numbersToSum {
		sums[index] = Sum(numbers)
	}
	return sums
}
