package memento

import (
	"testing"
)

func TestMementoPattern(t *testing.T) {
	editor := &Editor{}
	caretaker := &Caretaker{}

	editor.SetState("State 1")
	caretaker.Add(editor.SaveStateToMemento())

	editor.SetState("State 2")
	caretaker.Add(editor.SaveStateToMemento())

	editor.SetState("State 3")
	if editor.GetState() != "State 3" {
		t.Errorf("Expected State 3, but got %s", editor.GetState())
	}

	editor.GetStateFromMemento(caretaker.Get(1))
	if editor.GetState() != "State 2" {
		t.Errorf("Expected State 2, but got %s", editor.GetState())
	}

	editor.GetStateFromMemento(caretaker.Get(0))
	if editor.GetState() != "State 1" {
		t.Errorf("Expected State 1, but got %s", editor.GetState())
	}
}
