package shapes

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
	rectangle := Rectangle{10.0, 10.0}
	want := 40.0
	checkPerimeter(t, rectangle, want)
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 10.0}, 100.00},
		{Circle{10.0}, 314.1592653589793},
		{Triangle{12, 6, 10, 20, 30}, 36.0},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %.2f want %.2f", got, tt.want)
		}
	}
}
