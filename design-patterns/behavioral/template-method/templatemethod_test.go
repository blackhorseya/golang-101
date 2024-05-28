package templatemethod

import (
	"testing"
)

func TestPDFDocument(t *testing.T) {
	pdfDoc := &PDFDocument{}
	template := &DocumentTemplate{}

	expectedOutput := "Document processed successfully"
	if template.Process(pdfDoc) != expectedOutput {
		t.Errorf("Expected %s, but got %s", expectedOutput, template.Process(pdfDoc))
	}
}

func TestWordDocument(t *testing.T) {
	wordDoc := &WordDocument{}
	template := &DocumentTemplate{}

	expectedOutput := "Document processed successfully"
	if template.Process(wordDoc) != expectedOutput {
		t.Errorf("Expected %s, but got %s", expectedOutput, template.Process(wordDoc))
	}
}
