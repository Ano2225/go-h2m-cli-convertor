package converter

import (
	//"log"
	"testing"
)

func TestConvertMarkdownToHTML(t *testing.T) {
	markdown := []byte("# Titre")
	expected := "<h1>Titre</h1>\n"

	html, err := ConvertMarkdownToHTML(markdown)
	if err != nil {
		t.Fatalf("Erreur de conversion : %v", err)
	}

	if html != expected {
		t.Errorf("Resultat inattendu: attendu %q, obtenu %q", expected, html)
	}

}

func TestConvertHtmlToMarkdown(t *testing.T) {
	htmlContent := string("<h1>Titre</h1>")
	expected := "# Titre"

	markdown, err := ConvertHTMLtoMarkdown(htmlContent)
	if err != nil {
		t.Fatalf("Erreur de conversion : %v", err)
	}

	if markdown != expected {
		t.Errorf("Resultat inattendu: attendu %q, obtenu %q", expected, markdown)
	}

}
