package prototype

// Shape 定義形狀的原型接口
type Shape interface {
	Clone() Shape
	GetType() string
}

// Circle 實現 Shape 接口
type Circle struct {
	Radius int
}

func (c *Circle) Clone() Shape {
	return &Circle{Radius: c.Radius}
}

func (c *Circle) GetType() string {
	return "Circle"
}

// Rectangle 實現 Shape 接口
type Rectangle struct {
	Width, Height int
}

func (r *Rectangle) Clone() Shape {
	return &Rectangle{Width: r.Width, Height: r.Height}
}

func (r *Rectangle) GetType() string {
	return "Rectangle"
}
