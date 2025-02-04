package cmd

import (
	"fmt"
	"log"
	"github.com/spf13/cobra"
	"github.com/tonuser/markdown-to-html/internal/converter"
	"github.com/tonuser/markdown-to-html/internal/utils"

)

// Commande html2md
var htmlToMdCmd = &cobra.Command{
	Use:     "html2md",
	Aliases: []string{"html-to-markdown"},
	Short:   "Convertir HTML en Markdown",
	Run: func(cmd *cobra.Command, args []string) {
		inputFile, _ := cmd.Flags().GetString("input")
		outputFile, _ := cmd.Flags().GetString("output")

		if inputFile == "" {
			fmt.Println("Entrez le chemin du fichier HTML")
			fmt.Scanln(&inputFile)
		}
		if inputFile == "" {
			log.Fatal("Erreur: Aucun fichier HTML fourni")
		}

		// Lire le fichier HTML
		htmlContent, err := utils.ReadFile(inputFile)
		if err != nil {
			log.Fatalf("Erreur de lecture : %v", err)
		}

		// Convertir en Markdown
		markdown, err := converter.ConvertHTMLtoMarkdown(string(htmlContent))
		if err != nil {
			log.Fatalf("Erreur de conversion : %v", err)
		}

		if outputFile == "" {
			fmt.Println("Entrez le chemin du fichier Markdown (ou laisser vide pour afficher dans la console)")
			fmt.Scanln(&outputFile)
		}

		if outputFile != "" {
			if err := utils.WriteFile(outputFile, []byte(markdown)); err != nil {
				log.Fatalf("Erreur d'écriture : %v", err)
			}
			fmt.Printf("Conversion réussie ! Markdown enregistré dans %s\n", outputFile)
		} else {
			fmt.Println("\nRésultat de la conversion :")
			fmt.Println("------------------------------")
			fmt.Println(markdown)
		}
	},
}

func init() {
	htmlToMdCmd.Flags().StringP("input", "i", "", "Fichier HTML d'entrée")
	htmlToMdCmd.Flags().StringP("output", "o", "", "Fichier Markdown de sortie")
}
