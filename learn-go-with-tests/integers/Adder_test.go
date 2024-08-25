package integers

import "fmt"
import "testing"

func assertInteger(t testing.TB, got int, want int) {

	t.Helper()

	if got != want {
		t.Errorf("Expected '%d' but got '%d'", want, got)
	}

}

func TestAdder(t *testing.T) {

	sum := Add(2, 2)
	expected := 4
	assertInteger(t, sum, expected)

}

func ExampleAdd() {

	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6

}
