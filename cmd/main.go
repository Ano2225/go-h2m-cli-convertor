package main

import (
	"bufio"
	//"flag"
	"fmt"
	"log"
	"os"
	"strings"

	// "os"

	"github.com/tonuser/markdown-to-html/internal/converter"
	"github.com/tonuser/markdown-to-html/internal/utils"
)

const (
    mdToHtml = "1"
    htmlToMd = "2"
)

func main() {
    //Affichage du menu
    fmt.Println("=== Convertisseur de fichier ===")
    fmt.Println("1. Markdown vers HTML")
    fmt.Println("2. HTML vers Markdown")
    fmt.Println("Choisissez une option (1-2):")

    reader := bufio.NewReader(os.Stdin)
    choice, _ := reader.ReadString('\n')
    choice = strings.TrimSpace(choice)

    switch (choice) {
    case mdToHtml:
        MarkdownToHTML()
    case htmlToMd:
        HtmlToMd()
    default:
        log.Fatal("Option invalide")
    }

}


func MarkdownToHTML() {
    var inputFile, outputFile string

    //Demande du fichier d'entrée
    fmt.Println("Entrez le chemin du fichier Markdown")
    fmt.Scanln(&inputFile)

    if inputFile == "" {
        log.Fatal("Erreur: Aucun fichier markdown fourni")
    }

    // Lecture du fichier Markdown
    markdown, err := utils.ReadFile(inputFile)
    if err != nil {
        log.Fatalf("Erreur de lecture du fichier MarkDown : %v", err)
    }

    // Conversion en HTML
    html, err := converter.ConvertMarkdownToHTML(markdown)
    if err != nil {
        log.Fatalf("Erreur de conversion du fichier HTML : %v", err)
    }

    //Demande du fichier de sortie
    fmt.Println("Entrez le chemin du fichier HTML (ou laisser vide pour afficher dans la console)")
    fmt.Scanln(&outputFile)

    if outputFile != "" {
        if err := utils.WriteFile(outputFile, []byte(html)); err != nil {
            log.Fatalf("Erreur d'écriture du fichier HTML : %v", err)
        }
        fmt.Printf("Conversion réussie ! HTML enregistré dans %s\n", outputFile)
    } else {
        fmt.Println("\n Resultat de la conversion :")
        fmt.Println("-------------------------------")
        fmt.Println(html)
    }
}

func HtmlToMd() {
    var inputFile, outputFile string

    //Demande du fichier d'entrée HTML
    fmt.Println("Entrez le chemin du fichier HTML")
    fmt.Scanln(&inputFile)

    if inputFile == "" {
        log.Fatal("Erreur: Aucun fichier HTML fourni")
    }

    //Lecture du fichier
    htmlContent, err := utils.ReadFile(inputFile)
    if err != nil {
        log.Fatalf("Erreur de lecture du fichier HTML: %v", err)
    }

    //Conversion en Markdown
    markdown, err := converter.ConvertHTMLtoMarkdown(string(htmlContent))
    if err != nil {
        log.Fatalf("Erreur de conversion du fichier HTML : %v", err)
    }

    //Demande du fichier de sortie
    fmt.Println("Entrez le chemin du fichier Markdown (ou laisser vide pour afficher dans la console)")
    fmt.Scanln(&outputFile)

    if outputFile != "" {
        if err := utils.WriteFile(outputFile, []byte(markdown)); err != nil {
            log.Fatalf("Erreur d'ecriture du fichier Markdown: %v", err)
        }
        fmt.Printf("Conversion réussie ! Markdown enregistré dans %s\n", outputFile)
    } else {
        fmt.Println("\nRésultat de la conversion :")
        fmt.Println("------------------------------")
        fmt.Println(markdown)
    }

}