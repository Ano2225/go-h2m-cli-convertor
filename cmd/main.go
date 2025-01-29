package main

import (
    "flag"
    "fmt"
    "log"
   // "os"

    "github.com/tonuser/markdown-to-html/internal/converter"
    "github.com/tonuser/markdown-to-html/internal/utils"
)

func main() {
    // Définition des arguments en ligne de commande
    inputFile := flag.String("input", "", "Chemin du fichier Markdown à convertir")
    outputFile := flag.String("output", "", "Chemin du fichier HTML généré")
    flag.Parse()

    // Vérification si un fichier Markdown a été fourni
    if *inputFile == "" {
        log.Fatal("Erreur : Veuillez fournir un fichier Markdown avec --input")
    }

    // Lecture du fichier Markdown
    markdown, err := utils.ReadFile(*inputFile)
    if err != nil {
        log.Fatalf("Erreur de lecture du fichier : %v", err)
    }

    // Conversion en HTML
    html, err := converter.ConvertMarkdownToHTML(markdown)
    if err != nil {
        log.Fatalf("Erreur de conversion : %v", err)
    }

    // Écriture dans le fichier HTML ou affichage en console
    if *outputFile != "" {
        if err := utils.WriteFile(*outputFile, []byte(html)); err != nil {
            log.Fatalf("Erreur d'écriture du fichier : %v", err)
        }
        fmt.Printf("Conversion réussie ! HTML enregistré dans %s\n", *outputFile)
    } else {
        fmt.Println(html)
    }
}