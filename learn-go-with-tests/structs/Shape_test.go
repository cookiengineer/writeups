package structs

import "testing"

func assertArea(t testing.TB, shape Shape, want float64) {

	t.Helper()

	got := shape.Area()

	if got != want {
		t.Errorf("Expected %.2f and got %.2f", want, got)
	}

}

func TestArea(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {

		rectangle := Rectangle{12, 6}
		assertArea(t, rectangle, 72.0)

	})

	t.Run("circles", func(t *testing.T) {

		circle := Circle{10}
		assertArea(t, circle, 314.1592653589793)

	})

	t.Run("triangles", func(t *testing.T) {

		triangle := Triangle{12, 6}
		assertArea(t, triangle, 36.0)

	})

}

