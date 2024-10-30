package clone

import "testing"

func TestCloningWord(t *testing.T) {
	want := CloningWord("")

	got := ""

	if want != got {
		t.Errorf("got %q want %q", got, want)
	}
}
