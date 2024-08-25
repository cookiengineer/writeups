package slices

func Sum(numbers []int) int {

	var result int

	for n := 0; n < len(numbers); n++ {
		result += numbers[n]
	}

	return result

}
