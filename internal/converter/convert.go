package converter

import (
	"bytes"
	"strings"
	"github.com/yuin/goldmark"
    "github.com/JohannesKaufmann/html-to-markdown"
)

//ConvertMarkdownToHTML converts Markdown to HTML
func ConvertMarkdownToHTML(markdown []byte) (string, error) {
	var buf bytes.Buffer
	md := goldmark.New()
	if err := md.Convert(markdown, &buf);
	 err != nil {
		return "" , err
	}
	return buf.String(), nil
}


//ConvertHtmltoMarkdown converts HTML to Markdown
func ConvertHTMLtoMarkdown(htmlContent string) (string, error) {
    // Créer un nouveau convertisseur
    converter := md.NewConverter("", true, nil)

    // Conversion
    markdownContent, err := converter.ConvertString(htmlContent)
    if err != nil {
        return "", err
    }

    return strings.TrimSpace(markdownContent), nil
}
