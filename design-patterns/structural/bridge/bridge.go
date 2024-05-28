package bridge

// Color 定義顏色的實現接口
type Color interface {
	ApplyColor() string
}

// Red 是具體的顏色實現
type Red struct{}

func (r *Red) ApplyColor() string {
	return "Red"
}

// Blue 是具體的顏色實現
type Blue struct{}

func (b *Blue) ApplyColor() string {
	return "Blue"
}

// Shape 定義形狀的抽象
type Shape interface {
	Draw() string
}

// Circle 是擴充的形狀抽象，持有顏色的引用
type Circle struct {
	Color Color
}

func (c *Circle) Draw() string {
	return "Circle drawn with color " + c.Color.ApplyColor()
}

// Rectangle 是擴充的形狀抽象，持有顏色的引用
type Rectangle struct {
	Color Color
}

func (r *Rectangle) Draw() string {
	return "Rectangle drawn with color " + r.Color.ApplyColor()
}
