package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Steve", "")
		want := "Hello, Steve"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World', when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in french", func(t *testing.T) {
		got := Hello("Marc", "French")
		want := "Bonjour, Marc"

		assertCorrectMessage(t, got, want)
	})
}
