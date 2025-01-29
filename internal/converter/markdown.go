package converter

import (
	"bytes"

	"github.com/yuin/goldmark"
)

//ConvertMarkdownToHTML
func ConvertMarkdownToHTML(markdown []byte) (string, error) {
	var buf bytes.Buffer
	md := goldmark.New()
	if err := md.Convert(markdown, &buf);
	 err != nil {
		return "" , err
	}
	return buf.String(), nil
}
