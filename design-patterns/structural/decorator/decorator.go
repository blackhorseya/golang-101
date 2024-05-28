package decorator

// Text 定義一個組件接口
type Text interface {
	Render() string
}

// PlainText 是具體的組件，實現了 Text 接口
type PlainText struct {
	Content string
}

func (pt *PlainText) Render() string {
	return pt.Content
}

// TextDecorator 是裝飾器，實現了 Text 接口，並持有一個 Text 對象的引用
type TextDecorator struct {
	Text Text
}

func (td *TextDecorator) Render() string {
	return td.Text.Render()
}

// BoldDecorator 是具體裝飾器，為 Text 添加加粗行為
type BoldDecorator struct {
	TextDecorator
}

// NewBoldDecorator 用於創建 BoldDecorator 實例.
func NewBoldDecorator(text Text) *BoldDecorator {
	return &BoldDecorator{
		TextDecorator: TextDecorator{
			Text: text,
		},
	}
}

func (bd *BoldDecorator) Render() string {
	return "<b>" + bd.TextDecorator.Render() + "</b>"
}

// ItalicDecorator 是具體裝飾器，為 Text 添加斜體行為
type ItalicDecorator struct {
	TextDecorator
}

// NewItalicDecorator 用於創建 ItalicDecorator 實例.
func NewItalicDecorator(text Text) *ItalicDecorator {
	return &ItalicDecorator{
		TextDecorator: TextDecorator{
			Text: text,
		},
	}
}

func (id *ItalicDecorator) Render() string {
	return "<i>" + id.TextDecorator.Render() + "</i>"
}
