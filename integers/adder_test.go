package integers

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	sum := Add(2, 3)

	expected := 5

	if sum != expected {
		t.Errorf("got '%d' expected '%d'", sum, expected)
	}
}

func ExampleAdd() {
	sum := Add(3, 3)

	fmt.Println(sum)

	// Output: 6
}
