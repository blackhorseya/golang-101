package decorator

import (
	"testing"
)

func TestDecoratorPattern(t *testing.T) {
	text := &PlainText{Content: "Hello, World!"}

	boldText := NewBoldDecorator(text)
	italicText := NewItalicDecorator(text)
	boldItalicText := NewBoldDecorator(italicText)

	if text.Render() != "Hello, World!" {
		t.Errorf("Expected 'Hello, World!', but got %s", text.Render())
	}

	if boldText.Render() != "<b>Hello, World!</b>" {
		t.Errorf("Expected '<b>Hello, World!</b>', but got %s", boldText.Render())
	}

	if italicText.Render() != "<i>Hello, World!</i>" {
		t.Errorf("Expected '<i>Hello, World!</i>', but got %s", italicText.Render())
	}

	if boldItalicText.Render() != "<b><i>Hello, World!</i></b>" {
		t.Errorf("Expected '<b><i>Hello, World!</i></b>', but got %s", boldItalicText.Render())
	}
}
