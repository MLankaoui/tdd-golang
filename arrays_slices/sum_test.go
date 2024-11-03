package arraysslices

import "testing"

func TestSum(t *testing.T) {
	// creating and initializing an array

	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)

	want := 15

	if got != want {
		t.Errorf("got '%d' want '%d', %v", got, want, numbers)
	}
}
