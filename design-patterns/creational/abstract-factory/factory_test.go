package abstractfactory

import (
	"testing"
)

func TestWindowsFactory(t *testing.T) {
	factory := WindowsFactory{}

	button := factory.CreateButton()
	textBox := factory.CreateTextBox()

	if button.Paint() != "Rendering a button in Windows style" {
		t.Errorf("Expected 'Rendering a button in Windows style', but got %s", button.Paint())
	}

	if textBox.Paint() != "Rendering a text box in Windows style" {
		t.Errorf("Expected 'Rendering a text box in Windows style', but got %s", textBox.Paint())
	}
}

func TestMacOSFactory(t *testing.T) {
	factory := MacOSFactory{}

	button := factory.CreateButton()
	textBox := factory.CreateTextBox()

	if button.Paint() != "Rendering a button in MacOS style" {
		t.Errorf("Expected 'Rendering a button in MacOS style', but got %s", button.Paint())
	}

	if textBox.Paint() != "Rendering a text box in MacOS style" {
		t.Errorf("Expected 'Rendering a text box in MacOS style', but got %s", textBox.Paint())
	}
}
