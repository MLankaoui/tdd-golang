package hello

import "testing"

func TestXxx(t *testing.T) {

	// testing the case where something is passed to our function
	t.Run("when something is passed", func(t *testing.T) {
		got := Hello("Marouane")

		want := "Hello, Marouane!"

		if got != want {
			t.Errorf("got %q expected %q", got, want)
		}
	})

	// testing the case when nothing is passed
	t.Run("when nothing is passed", func(t *testing.T) {

		got := Hello()

		want := "Hello, world!"

		if got != want {
			t.Errorf("got %q expected %q", got, want)
		}
	})
}
