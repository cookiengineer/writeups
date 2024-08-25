package main

import "testing"

func assertMessage(t testing.TB, got string, want string) {

	t.Helper()

	if got != want {
		t.Errorf("Expected '%q' but got '%q'", got, want)
	}

}

func TestHello(t *testing.T) {

	t.Run("Say hello to world", func(t *testing.T) {

		got := Hello("world")
		want := "Hello, world!"
		assertMessage(t, got, want)

	})

	t.Run("Say hello to world by default", func(t *testing.T) {

		got := Hello("")
		want := "Hello, world!"
		assertMessage(t, got, want)

	})

	t.Run("Say hello to people", func (t *testing.T) {

		got := Hello("John")
		want := "Hello, John!"
		assertMessage(t, got, want)

	})

}
