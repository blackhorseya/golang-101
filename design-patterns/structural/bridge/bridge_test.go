package bridge

import (
	"testing"
)

func TestBridgePattern(t *testing.T) {
	red := &Red{}
	blue := &Blue{}

	circle := &Circle{Color: red}
	rectangle := &Rectangle{Color: blue}

	expectedCircle := "Circle drawn with color Red"
	expectedRectangle := "Rectangle drawn with color Blue"

	if circle.Draw() != expectedCircle {
		t.Errorf("Expected %s, but got %s", expectedCircle, circle.Draw())
	}

	if rectangle.Draw() != expectedRectangle {
		t.Errorf("Expected %s, but got %s", expectedRectangle, rectangle.Draw())
	}

	// 更改顏色
	circle.Color = blue
	expectedCircleBlue := "Circle drawn with color Blue"
	if circle.Draw() != expectedCircleBlue {
		t.Errorf("Expected %s, but got %s", expectedCircleBlue, circle.Draw())
	}
}
