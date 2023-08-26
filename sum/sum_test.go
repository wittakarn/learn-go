package sum

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		actual := Sum(numbers)
		expected := 15

		if actual != expected {
			t.Errorf("got %d want %d given, %v", actual, expected, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		actual := Sum(numbers)
		expected := 6

		if actual != expected {
			t.Errorf("got %d want %d given, %v", actual, expected, numbers)
		}
	})
}
