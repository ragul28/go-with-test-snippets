package structs

import "testing"

func TestPerimeter(t *testing.T) {
	ractangle := Rectangle{10.0, 10.0}
	got := Perimeter(ractangle)
	went := 40.0

	if got != went {
		t.Errorf("got %.2f went %.2f", got, went)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, went float64) {
		t.Helper()
		got := shape.Area()
		if got != went {
			t.Errorf("got %g went %g", got, went)
		}
	}

	t.Run("ractangles", func(t *testing.T) {
		ractangle := Rectangle{12.0, 6.0}
		checkArea(t, ractangle, 72.0)
	})

	t.Run("Circles", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}

// Table driven tests
func TestAreaTT(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{12.0, 6.0}, want: 72.0},
		{name: "Circle", shape: Circle{10.0}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g went %g", tt.name, got, tt.want)
			}
		})
	}
}
