package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	r := Rectangle{Width: 10.0, Height: 20.0}
	got := r.Perimeter()
	want := 60.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, &rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, &circle, 314.1592653589793)
	})
}

func TestArea2(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: &Rectangle{Width: 12, Height: 8}, want: 96.0},
		{shape: &Circle{Radius: 10}, want: 314.1592653589793},
		{shape: &Triangle{Width: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %.2f want %.2f", got, tt.want)
		}
	}
}
