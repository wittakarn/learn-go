package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		actual := SayHello("Wittakarn", "English")
		expected := "Hello, Wittakarn"

		assertCorrectMessage(t, actual, expected)
	})

	t.Run("saying hello world when name is not present", func(t *testing.T) {
		actual := SayHello("", "English")
		expected := "Hello, World"

		assertCorrectMessage(t, actual, expected)
	})

	t.Run("saying hello in Spanish", func(t *testing.T) {
		actual := SayHello("Elodie", "Spanish")
		expected := "Hola, Elodie"

		assertCorrectMessage(t, actual, expected)
	})

	t.Run("saying hello in French", func(t *testing.T) {
		actual := SayHello("Darcy", "French")
		expected := "Bonjour, Darcy"

		assertCorrectMessage(t, actual, expected)
	})
}

func assertCorrectMessage(t testing.TB, actual string, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("actual %q expected %q", actual, expected)
	}
}
