package containsany

import "testing"

func TestWordContainsAny(t *testing.T) {
	got := WordContainsAny("fail", "ui")

	want := true

	if got != want {
		t.Errorf("got '%t' want '%t'", got, want)
	}
}
