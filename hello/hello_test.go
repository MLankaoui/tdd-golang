package hello

import "testing"

func TestXxx(t *testing.T) {
	// testing the case where something is passed to our function
	t.Run("when something is passed", func(t *testing.T) {
		got := Hello("Marouane", "")

		want := "Hello, Marouane!"

		assertCorrectMessage(t, got, want)
	})

	// testing the case when nothing is passed
	t.Run("when nothing is passed", func(t *testing.T) {

		got := Hello("", "")

		want := "Hello, world!"

		assertCorrectMessage(t, got, want)

	})

	t.Run("saying hello in spanish", func(t *testing.T) {

		got := Hello("Marouane", "Spanish")

		want := "Hola, Marouane!"

		assertCorrectMessage(t, got, want)

	})
	t.Run("saying hello in french", func(t *testing.T) {

		got := Hello("Marouane", "French")

		want := "Bonjour, Marouane!"

		assertCorrectMessage(t, got, want)

	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q expected %q", got, want)
	}
}
