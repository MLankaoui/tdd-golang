package arraysslices

import "testing"

func TestSum(t *testing.T) {
	// creating and initializing an array

	t.Run("when we have a fixed size we use arrays", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 6}

		got := Sum(numbers)

		want := 16

		if got != want {
			t.Errorf("got '%d' want '%d', %v", got, want, numbers)
		}
	})

	t.Run("when we don t have a fixed size we use slices", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)

		want := 6

		if got != want {
			t.Errorf("got '%d' want '%d', %v", got, want, numbers)
		}
	})
}
