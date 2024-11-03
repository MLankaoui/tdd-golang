package containsfunc

import "testing"

func TestContainsFuncRune(t *testing.T) {
	t.Run("if it is true", func(t *testing.T) {
		got := ContainsFuncRune("hello")

		want := true

		assertCorrectMessage(t, got, want)
	})

	t.Run("testing if it false", func(t *testing.T) {
		got := ContainsFuncRune("test")

		want := false

		assertCorrectMessage(t, got, want)

	})
}

func assertCorrectMessage(t testing.TB, got, want bool) {
	t.Helper()

	if got != want {
		t.Errorf("got '%t' want '%t'", got, want)
	}
}
