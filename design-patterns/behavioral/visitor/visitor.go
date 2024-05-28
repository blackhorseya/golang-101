package visitor

// Visitor 定義了訪問者接口
type Visitor interface {
	VisitCircle(circle *Circle)
	VisitRectangle(rectangle *Rectangle)
}
