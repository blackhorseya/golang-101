package templatemethod

const msgOK = "Document processed successfully"

// Document 定義文檔處理的抽象類
type Document interface {
	Open()
	Parse()
	Save()
	Close()
}

type DocumentTemplate struct{}

func (dt *DocumentTemplate) Process(doc Document) string {
	doc.Open()
	doc.Parse()
	doc.Save()
	doc.Close()

	return msgOK
}
