package visitor

// Circle 是具體元素，表示圓形
type Circle struct {
	Radius float64
}

func (c *Circle) Accept(visitor Visitor) {
	visitor.VisitCircle(c)
}

// Rectangle 是具體元素，表示矩形
type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Accept(visitor Visitor) {
	visitor.VisitRectangle(r)
}
