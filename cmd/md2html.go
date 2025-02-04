package cmd

import (
	"fmt"
	"log"
	"github.com/spf13/cobra"
	"github.com/tonuser/markdown-to-html/internal/converter"
	"github.com/tonuser/markdown-to-html/internal/utils"
)

// Commande md2html
var mdToHtmlCmd = &cobra.Command{
	Use:     "md2html",
	Aliases: []string{"markdown-to-html"},
	Short:   "Convertir Markdown en HTML",
	Run: func(cmd *cobra.Command, args []string) {
		inputFile, _ := cmd.Flags().GetString("input")
		outputFile, _ := cmd.Flags().GetString("output")

		if inputFile == "" {
			fmt.Println("Entrez le chemin du fichier Markdown")
			fmt.Scanln(&inputFile)
		}
		if inputFile == "" {
			log.Fatal("Erreur: Aucun fichier Markdown fourni")
		}

		// Lire le fichier Markdown
		markdown, err := utils.ReadFile(inputFile)
		if err != nil {
			log.Fatalf("Erreur de lecture : %v", err)
		}

		// Convertir en HTML
		html, err := converter.ConvertMarkdownToHTML(markdown)
		if err != nil {
			log.Fatalf("Erreur de conversion : %v", err)
		}

		if outputFile == "" {
			fmt.Println("Entrez le chemin du fichier HTML (ou laisser vide pour afficher dans la console)")
			fmt.Scanln(&outputFile)
		}

		if outputFile != "" {
			if err := utils.WriteFile(outputFile, []byte(html)); err != nil {
				log.Fatalf("Erreur d'écriture : %v", err)
			}
			fmt.Printf("Conversion réussie ! HTML enregistré dans %s\n", outputFile)
		} else {
			fmt.Println("\nRésultat de la conversion :")
			fmt.Println("-------------------------------")
			fmt.Println(html)
		}
	},
}

func init() {
	mdToHtmlCmd.Flags().StringP("input", "i", "", "Fichier Markdown d'entrée")
	mdToHtmlCmd.Flags().StringP("output", "o", "", "Fichier HTML de sortie")
}
