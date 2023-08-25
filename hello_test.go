package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		actual := SayHello("Wittakarn")
		expected := "Hello, Wittakarn"

		if actual != expected {
			t.Errorf("actual %q expected %q", actual, expected)
		}
	})

	t.Run("saying hello world when name is not present", func(t *testing.T) {
		actual := SayHello("")
		expected := "Hello, World"

		if actual != expected {
			t.Errorf("actual %q expected %q", actual, expected)
		}
	})
}
