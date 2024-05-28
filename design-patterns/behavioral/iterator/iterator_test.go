package iterator

import (
	"testing"
)

func TestNameIterator(t *testing.T) {
	names := []string{"Alice", "Bob", "Charlie", "Diana"}
	collection := NewNameCollection(names)
	iterator := collection.CreateIterator()

	expectedNames := []string{"Alice", "Bob", "Charlie", "Diana"}
	index := 0

	for iterator.HasNext() {
		name := iterator.Next()
		if name != expectedNames[index] {
			t.Errorf("Expected %s, but got %s", expectedNames[index], name)
		}
		index++
	}

	if index != len(expectedNames) {
		t.Errorf("Expected to iterate over %d names, but iterated over %d", len(expectedNames), index)
	}
}
