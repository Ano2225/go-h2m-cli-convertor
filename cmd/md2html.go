package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
	"github.com/tonuser/markdown-to-html/internal/converter"
	"github.com/tonuser/markdown-to-html/internal/utils"
)

// Commande md2html
var mdToHtmlCmd = &cobra.Command{
	Use:     "md2html",
	Aliases: []string{"markdown-to-html"},
	Short:   "Convertir Markdown en HTML",
	Long: "Convertir un fichier Markdown en HTML",
	RunE: func(cmd *cobra.Command, args []string) error {
		inputFile, _ := cmd.Flags().GetString("input")
		outputFile, _ := cmd.Flags().GetString("output")

		if inputFile == "" {
			fmt.Println("Entrez le chemin du fichier Markdown")
			fmt.Scanln(&inputFile)
		}
		if inputFile == "" {
			log.Fatal("Erreur: Aucun fichier Markdown fourni")
		}

		//Démarrage du loader
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Start()
		s.Suffix = " Conversion en cours..."

		// Lire le fichier Markdown
		markdown, err := utils.ReadFile(inputFile)
		if err != nil {
			s.Stop()
			log.Fatalf("Erreur de lecture : %v", err)
		}

		// Convertir en HTML
		html, err := converter.ConvertMarkdownToHTML(markdown)
		if err != nil {
			s.Stop()
			log.Fatalf("Erreur de conversion : %v", err)
		}

		//Arret du loader après conversion
		s.Stop()
		
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
		return nil
	},
}
