package strategy

import (
	"reflect"
	"testing"
)

func TestBubbleSortStrategy(t *testing.T) {
	data := []int{34, 7, 23, 32, 5, 62}
	expected := []int{5, 7, 23, 32, 34, 62}

	bubbleSort := &BubbleSortStrategy{}
	context := NewSortContext(bubbleSort)
	sortedData := context.Sort(data)

	if !reflect.DeepEqual(sortedData, expected) {
		t.Errorf("Expected %v, but got %v", expected, sortedData)
	}
}

func TestQuickSortStrategy(t *testing.T) {
	data := []int{34, 7, 23, 32, 5, 62}
	expected := []int{5, 7, 23, 32, 34, 62}

	quickSort := &QuickSortStrategy{}
	context := NewSortContext(quickSort)
	sortedData := context.Sort(data)

	if !reflect.DeepEqual(sortedData, expected) {
		t.Errorf("Expected %v, but got %v", expected, sortedData)
	}
}
