package main

import "testing"

func TestHello(t *testing.T) {
	actual := SayHello()
	expected := "Hello world"

	if actual != expected {
		t.Errorf("actual %q expected %q", actual, expected)
	}
}
