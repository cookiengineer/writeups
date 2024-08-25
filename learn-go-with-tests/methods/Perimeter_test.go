package methods

import "testing"

func TestArea(t *testing.T) {

	got := Area(12.0, 6.0)
	want := 72.0

	if got != want {
		t.Errorf("Expected %.2f and got %.2f", want, got)
	}

}
func TestPerimeter(t *testing.T) {

	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("Expected %.2f and got %.2f", want, got)
	}

}


