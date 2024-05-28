package composite

import (
	"testing"
)

func TestCompositePattern(t *testing.T) {
	circle := &Circle{}
	rectangle := &Rectangle{}

	compositeGraphic := &CompositeGraphic{}
	compositeGraphic.Add(circle)
	compositeGraphic.Add(rectangle)

	expectedLeaf := "Circle"
	if circle.Draw() != expectedLeaf {
		t.Errorf("Expected %s, but got %s", expectedLeaf, circle.Draw())
	}

	expectedComposite := "Composite: [Circle Rectangle ]"
	if compositeGraphic.Draw() != expectedComposite {
		t.Errorf("Expected %s, but got %s", expectedComposite, compositeGraphic.Draw())
	}
}
