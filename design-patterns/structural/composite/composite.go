package composite

// Graphic 定義圖形的接口
type Graphic interface {
	Draw() string
}

// Circle 是圖形的葉子節點
type Circle struct{}

func (c *Circle) Draw() string {
	return "Circle"
}

// Rectangle 是圖形的葉子節點
type Rectangle struct{}

func (r *Rectangle) Draw() string {
	return "Rectangle"
}

// CompositeGraphic 是圖形的組合節點
type CompositeGraphic struct {
	children []Graphic
}

func (cg *CompositeGraphic) Add(graphic Graphic) {
	cg.children = append(cg.children, graphic)
}

func (cg *CompositeGraphic) Draw() string {
	result := "Composite: ["
	for _, child := range cg.children {
		result += child.Draw() + " "
	}
	result += "]"
	return result
}
