package factorymethod

// Shape is the interface that all shapes must implement.
type Shape interface {
	// Draw returns a string representation of the shape.
	Draw() string
}

// Circle is a shape that represents a circle.
type Circle struct {
}

func (x *Circle) Draw() string {
	return "Circle"
}

// Rectangle is a shape that represents a rectangle.
type Rectangle struct {
}

func (x *Rectangle) Draw() string {
	return "Rectangle"
}
