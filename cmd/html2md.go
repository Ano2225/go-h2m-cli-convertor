package cmd

import (
	"fmt"
	"log"
	"path/filepath"
	"time"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
	"github.com/tonuser/markdown-to-html/internal/converter"
	"github.com/tonuser/markdown-to-html/internal/utils"
)

// Commande html2md
var htmlToMdCmd = &cobra.Command{
	Use:     "html2md",
	Aliases: []string{"html-to-markdown"},
	Short:   "Convertir HTML en Markdown",
	Long:  "Convertir un fichier HTML en Markdown",
	RunE: func(cmd *cobra.Command, args []string) error {
		inputFile, _ := cmd.Flags().GetString("input")
		outputFile, _ := cmd.Flags().GetString("output")

		if inputFile == "" {
			fmt.Println("Entrez le chemin du fichier HTML")
			fmt.Scanln(&inputFile)
		}
		 
		
		if inputFile == "" {
			log.Fatal("Erreur: Aucun fichier HTML fourni")
		}

		  // Vérifier l'extension du fichier
		  if filepath.Ext(inputFile) != ".html" {
            log.Fatal("Erreur: Le fichier d'entrée doit avoir l'extension .html")
        }


		//Démarrage du loader
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Start()
		s.Suffix = " Conversion en cours..."

		// Lire le fichier HTML
		htmlContent, err := utils.ReadFile(inputFile)
		if err != nil {
			s.Stop()
			log.Fatalf("Erreur de lecture : %v", err)
		}

		// Convertir en Markdown
		markdown, err := converter.ConvertHTMLtoMarkdown(string(htmlContent))
		if err != nil {
			s.Stop()
			log.Fatalf("Erreur de conversion : %v", err)
		}

		//Arret du loader après conversion
		s.Stop()

		if outputFile == "" {
			fmt.Println("Entrez le chemin du fichier Markdown")
			fmt.Scanln(&outputFile)
		}

		  // Vérifier l'extension du fichier
		  if filepath.Ext(outputFile) != ".md" {
            log.Fatal("Erreur: Le fichier de sortie doit avoir l'extension .md")
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
		return nil
	},
}


