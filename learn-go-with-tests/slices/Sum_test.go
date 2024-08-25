package slices

import "testing"

func assertInteger(t testing.TB, got int, want int) {

	t.Helper()

	if got != want {
		t.Errorf("Expected '%d' but got '%d'", want, got)
	}

}

func TestSum(t *testing.T) {

	t.Run("5 numbers", func(t *testing.T) {

		numbers := []int{1,2,3,4,5}

		got := Sum(numbers)
		want := 15
		assertInteger(t, got, want)

	})

	t.Run("6 numbers", func(t *testing.T) {

		numbers := []int{1,2,3,4,5,6}

		got := Sum(numbers)
		want := 21
		assertInteger(t, got, want)

	})

}

func BenchmarkSum(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Sum([]int{1,2,3,4,5})
	}

}
