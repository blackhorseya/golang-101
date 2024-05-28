package adapter

import (
	"testing"
)

func TestAdapterPattern(t *testing.T) {
	legacyCircle := &LegacyCircle{}
	circleAdapter := &CircleAdapter{
		LegacyCircle: legacyCircle,
	}

	expected := "Drawing a circle using the old method"
	result := circleAdapter.Draw()

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
