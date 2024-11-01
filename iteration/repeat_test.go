package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", -1)

	want := ""

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// writing our first benchmark
func BenchmarkRepeat(b *testing.B) {
	// when the benchmark runs it will run b.N times and calculates how much is it gonna take.
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat(t *testing.TB) {
	repetedValue := Repeat("a", 5)
	fmt.Println(repetedValue)

	// Output: aaaaa
}
