package state

import (
	"testing"
)

func TestLightState(t *testing.T) {
	light := NewLight()

	light.PressSwitch() // Turning on the light.
	if _, ok := light.state.(*OnState); !ok {
		t.Errorf("Expected light to be in OnState")
	}

	light.PressSwitch() // Turning off the light.
	if _, ok := light.state.(*OffState); !ok {
		t.Errorf("Expected light to be in OffState")
	}

	light.PressSwitch() // Turning on the light.
	if _, ok := light.state.(*OnState); !ok {
		t.Errorf("Expected light to be in OnState")
	}
}
