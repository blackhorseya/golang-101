package chain

import (
	"testing"
)

func TestChainOfResponsibility(t *testing.T) {
	handler1 := &ConcreteHandler1{}
	handler2 := &ConcreteHandler2{}
	handler3 := &ConcreteHandler3{}

	handler1.SetNext(handler2)
	handler2.SetNext(handler3)

	tests := []struct {
		level    int
		expected string
	}{
		{1, "ConcreteHandler1 handled the request"},
		{2, "ConcreteHandler2 handled the request"},
		{3, "ConcreteHandler3 handled the request"},
		{4, "No handler available"},
	}

	for _, test := range tests {
		result := handler1.HandleRequest(test.level)
		if result != test.expected {
			t.Errorf("For level %d, expected %s, but got %s", test.level, test.expected, result)
		}
	}
}
