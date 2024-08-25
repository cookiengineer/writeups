package slices

func SumAll(slices ...[]int) []int {

	length := len(slices)
	sums := make([]int, length)

	for n, numbers := range slices {
		sums[n] = Sum(numbers)
	}

	return sums

}
