package abstractfactory

// Button is the interface that all buttons must implement.
type Button interface {
	Paint() string
}

// TextBox is the interface that all text boxes must implement.
type TextBox interface {
	Paint() string
}

// WindowsButton 實現 Button 接口
type WindowsButton struct{}

func (wb WindowsButton) Paint() string {
	return "Rendering a button in Windows style"
}

// MacOSButton 實現 Button 接口
type MacOSButton struct{}

func (mb MacOSButton) Paint() string {
	return "Rendering a button in MacOS style"
}

// WindowsTextBox 實現 TextBox 接口
type WindowsTextBox struct{}

func (wtb WindowsTextBox) Paint() string {
	return "Rendering a text box in Windows style"
}

// MacOSTextBox 實現 TextBox 接口
type MacOSTextBox struct{}

func (mtb MacOSTextBox) Paint() string {
	return "Rendering a text box in MacOS style"
}
