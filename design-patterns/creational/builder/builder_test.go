package builder

import (
	"testing"
)

func TestBuilderPattern(t *testing.T) {
	builder := NewConcreteHouseBuilder()
	director := NewDirector(builder)
	director.Construct()
	house := builder.GetHouse()

	if house.Foundation != "Concrete Foundation" {
		t.Errorf("Expected foundation to be 'Concrete Foundation', but got %s", house.Foundation)
	}

	if house.Walls != "Brick Walls" {
		t.Errorf("Expected walls to be 'Brick Walls', but got %s", house.Walls)
	}

	if house.Roof != "Slate Roof" {
		t.Errorf("Expected roof to be 'Slate Roof', but got %s", house.Roof)
	}
}
