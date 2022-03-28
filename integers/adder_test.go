package integers

import "testing"
import "fmt"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("got %d but wanted %d", sum, expected)
	}
}

func ExampleAdd() {
	sum := Add(5, 6)
	fmt.Println(sum)
	// Output: 11
}
