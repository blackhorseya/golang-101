package visitor

import "fmt"

// AreaVisitor 是具體訪問者，計算形狀的面積
type AreaVisitor struct{}

func (a *AreaVisitor) VisitCircle(circle *Circle) {
	area := 3.14 * circle.Radius * circle.Radius
	fmt.Printf("Circle Area: %.2f\n", area)
}

func (a *AreaVisitor) VisitRectangle(rectangle *Rectangle) {
	area := rectangle.Width * rectangle.Height
	fmt.Printf("Rectangle Area: %.2f\n", area)
}

// PerimeterVisitor 是具體訪問者，計算形狀的周長
type PerimeterVisitor struct{}

func (p *PerimeterVisitor) VisitCircle(circle *Circle) {
	perimeter := 2 * 3.14 * circle.Radius
	fmt.Printf("Circle Perimeter: %.2f\n", perimeter)
}

func (p *PerimeterVisitor) VisitRectangle(rectangle *Rectangle) {
	perimeter := 2 * (rectangle.Width + rectangle.Height)
	fmt.Printf("Rectangle Perimeter: %.2f\n", perimeter)
}
