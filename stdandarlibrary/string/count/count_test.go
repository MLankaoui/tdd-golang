package count

import "testing"

func TestCountMany(t *testing.T) {
	t.Run("when there is no substring passed", func(t *testing.T) {
		got := CountMany("free", "")

		want := 5

		assertCorrectMessage(t, got, want)
	})

	t.Run("whe there is one passed", func(t *testing.T) {
		got := CountMany("hello", "l")

		want := 2

		assertCorrectMessage(t, got, want)

	})
}

func assertCorrectMessage(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got '%d' want '%d", got, want)
	}
}
