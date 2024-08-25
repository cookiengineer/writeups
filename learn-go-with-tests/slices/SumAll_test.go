package slices

import "reflect"
import "testing"

func assertIntegers(t testing.TB, got []int, want []int) {

	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected '%v' but got '%v'", want, got)
	}

}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1,2,3}, []int{0,9})
	want := []int{6,9}
	assertIntegers(t, got, want)

}

func BenchmarkSumAll(b *testing.B) {

	for i := 0; i < b.N; i++ {
		SumAll([]int{1,2,3}, []int{0,9})
	}

}
