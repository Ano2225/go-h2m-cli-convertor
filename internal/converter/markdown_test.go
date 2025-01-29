package converter

import "testing"

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