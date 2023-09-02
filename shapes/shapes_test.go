package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	actual := Perimeter(Rectangle{10.0, 10.0})
	expected := 40.0

	if actual != expected {
		t.Errorf("actual %.2f expected %.2f", actual, expected)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()
		actual := shape.Area()
		if actual != expected {
			t.Errorf("actual %g expected %g", actual, expected)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		expected := 72.0
		checkArea(t, Rectangle{12.0, 6.0}, expected)
	})

	t.Run("circles", func(t *testing.T) {
		expected := 314.1592653589793
		checkArea(t, Circle{10}, expected)
	})

	t.Run("triangle", func(t *testing.T) {
		expected := 36.0
		checkArea(t, Triangle{12.0, 6.0}, expected)
	})
}
