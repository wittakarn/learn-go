package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	actual := Perimeter(10.0, 10.0)
	expected := 40.0

	if actual != expected {
		t.Errorf("actual %.2f expected %.2f", actual, expected)
	}
}
