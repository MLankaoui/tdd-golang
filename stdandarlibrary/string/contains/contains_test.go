package contains

import "testing"

func TestWordContains(t *testing.T) {
	got := WordContains("hello", "llo")

	want := true

	if got != want {
		t.Errorf("got '%t' want '%t ", got, want)
	}
}
