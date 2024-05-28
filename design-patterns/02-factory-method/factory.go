package factorymethod

// ShapeFactory is the interface that all shape factories must implement.
type ShapeFactory interface {
	// CreateShape returns a new shape.
	CreateShape() Shape
}

// CircleFactory is a factory that creates circles.
type CircleFactory struct {
}

// CreateShape returns a new circle.
func (x *CircleFactory) CreateShape() Shape {
	return &Circle{}
}

// RectangleFactory is a factory that creates rectangles.
type RectangleFactory struct {
}

// CreateShape returns a new rectangle.
func (x *RectangleFactory) CreateShape() Shape {
	return &Rectangle{}
}
