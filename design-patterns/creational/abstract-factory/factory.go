package abstractfactory

// GUIFactory 定義創建一族產品的方法
type GUIFactory interface {
	CreateButton() Button
	CreateTextBox() TextBox
}

// WindowsFactory 實現 GUIFactory 接口，創建 Windows 風格的產品
type WindowsFactory struct{}

func (wf WindowsFactory) CreateButton() Button {
	return &WindowsButton{}
}

func (wf WindowsFactory) CreateTextBox() TextBox {
	return &WindowsTextBox{}
}

// MacOSFactory 實現 GUIFactory 接口，創建 MacOS 風格的產品
type MacOSFactory struct{}

func (mf MacOSFactory) CreateButton() Button {
	return &MacOSButton{}
}

func (mf MacOSFactory) CreateTextBox() TextBox {
	return &MacOSTextBox{}
}
