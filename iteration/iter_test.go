package iteration

import "testing"
import "fmt"

func TestRepeat(t *testing.T) {
	got := Repeat("b", 5)
	want := "bbbbb"

	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	repeatCount := 100000

	for i := 0; i < b.N; i++ {
		Repeat("a", repeatCount)
	}
}

func ExampleRepeat() {
	s := Repeat("a", 10)
	fmt.Printf("%s", s)
	// Output: aaaaaaaaaa
}
