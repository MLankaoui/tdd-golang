package compare

import "testing"

func TestWordsCompare(t *testing.T) {
	got := WordsCompare("a", "b")

	want := -1

	if got != want {
		t.Errorf("got '%d' want '%d", got, want)
	}
}
