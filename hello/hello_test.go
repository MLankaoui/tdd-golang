package hello

import "testing"

func TestXxx(t *testing.T) {
	got := Hello("Marouane")

	want := "Hello, Marouane!"

	if got != want {
		t.Errorf("got %q expected %q", got, want)
	}
}
