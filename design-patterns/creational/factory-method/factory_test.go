package factorymethod

import (
	"testing"
)

func TestCircleFactoryCreatesCircle(t *testing.T) {
	factory := &CircleFactory{}
	shape := factory.CreateShape()

	_, ok := shape.(*Circle)
	if !ok {
		t.Errorf("Expected CircleFactory to create a Circle")
	}
}

func TestRectangleFactoryCreatesRectangle(t *testing.T) {
	factory := &RectangleFactory{}
	shape := factory.CreateShape()

	_, ok := shape.(*Rectangle)
	if !ok {
		t.Errorf("Expected RectangleFactory to create a Rectangle")
	}
}
