package prototype

import (
	"testing"
)

func TestPrototypePattern(t *testing.T) {
	circle := &Circle{Radius: 5}
	rectangle := &Rectangle{Width: 10, Height: 20}

	anotherCircle := circle.Clone()
	anotherRectangle := rectangle.Clone()

	if anotherCircle.GetType() != "Circle" {
		t.Errorf("Expected Circle, but got %s", anotherCircle.GetType())
	}

	if anotherRectangle.GetType() != "Rectangle" {
		t.Errorf("Expected Rectangle, but got %s", anotherRectangle.GetType())
	}

	if anotherCircle.(*Circle).Radius != circle.Radius {
		t.Errorf("Expected cloned Circle radius to be %d, but got %d", circle.Radius, anotherCircle.(*Circle).Radius)
	}

	if anotherRectangle.(*Rectangle).Width != rectangle.Width || anotherRectangle.(*Rectangle).Height != rectangle.Height {
		t.Errorf(
			"Expected cloned Rectangle dimensions to be %dx%d, but got %dx%d",
			rectangle.Width,
			rectangle.Height,
			anotherRectangle.(*Rectangle).Width,
			anotherRectangle.(*Rectangle).Height,
		)
	}
}
