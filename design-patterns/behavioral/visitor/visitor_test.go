package visitor

import (
	"testing"
)

func TestVisitorPattern(t *testing.T) {
	circle := &Circle{Radius: 5}
	rectangle := &Rectangle{Width: 10, Height: 5}

	areaVisitor := &AreaVisitor{}
	perimeterVisitor := &PerimeterVisitor{}

	shapes := []Shape{circle, rectangle}

	for _, shape := range shapes {
		shape.Accept(areaVisitor)
		shape.Accept(perimeterVisitor)
	}
}
