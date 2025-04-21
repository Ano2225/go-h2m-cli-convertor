package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertMarkdownToHTML(t *testing.T) {
	markdown := []byte("# Titre")
	expected := "<h1>Titre</h1>\n"

	html, err := ConvertMarkdownToHTML(markdown)

	assert.NoError(t, err)
	assert.Equal(t, expected, html)
}

func TestConvertHtmlToMarkdown(t *testing.T) {
	htmlContent := string("<h1>Titre</h1>")
	expected := "# Titre"

	markdown, err := ConvertHTMLtoMarkdown(htmlContent)

	assert.NoError(t, err)
	assert.Equal(t, expected, markdown)
}
