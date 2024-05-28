package interpreter

import (
	"testing"
)

func TestInterpreterPattern(t *testing.T) {
	expression := "3 4 + 2 -"
	parser := NewParser(expression)
	exp, err := parser.Parse()
	if err != nil {
		t.Fatalf("Failed to parse expression: %s", err)
	}

	result := exp.Interpret()
	expected := 5
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
