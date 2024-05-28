package templatemethod

import "fmt"

// PDFDocument 是具體的 PDF 文檔處理類
type PDFDocument struct {
	DocumentTemplate
}

func (pdf *PDFDocument) Open() {
	fmt.Println("Opening PDF document")
}

func (pdf *PDFDocument) Parse() {
	fmt.Println("Parsing PDF document")
}

func (pdf *PDFDocument) Save() {
	fmt.Println("Saving PDF document")
}

func (pdf *PDFDocument) Close() {
	fmt.Println("Closing PDF document")
}

// WordDocument 是具體的 Word 文檔處理類
type WordDocument struct {
	DocumentTemplate
}

func (word *WordDocument) Open() {
	fmt.Println("Opening Word document")
}

func (word *WordDocument) Parse() {
	fmt.Println("Parsing Word document")
}

func (word *WordDocument) Save() {
	fmt.Println("Saving Word document")
}

func (word *WordDocument) Close() {
	fmt.Println("Closing Word document")
}
