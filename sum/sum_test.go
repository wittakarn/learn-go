package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		actual := Sum(numbers)
		expected := 6

		if actual != expected {
			t.Errorf("got %d want %d given, %v", actual, expected, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		actual := SumAll([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}

		checkSums(t, actual, expected)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		actual := SumAll()
		expected := []int{}

		checkSums(t, actual, expected)
	})
}

func TestSumAllTails(t *testing.T) {

	t.Run("make the sums of some slices", func(t *testing.T) {
		actual := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		checkSums(t, actual, expected)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		actual := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}

		checkSums(t, actual, expected)
	})

}

func checkSums(t testing.TB, actual, expected []int) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual %v expected %v", actual, expected)
	}
}
