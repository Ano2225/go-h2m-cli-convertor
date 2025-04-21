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
	"github.com/tonuser/markdown-to-html/internal/vars/constants"
)

// html2md Command
var htmlToMdCmd = &cobra.Command{
	Use:     "html2md",
	Aliases: []string{"html-to-markdown"},
	Short:   "Convert HTML to Markdown",
	Long:    "Convert an HTML file to Markdown",
	RunE: func(cmd *cobra.Command, args []string) error {
		inputFile, _ := cmd.Flags().GetString("input")
		outputFile, _ := cmd.Flags().GetString("output")

		if inputFile == "" {
			fmt.Println("Enter the path to the HTML file:")
			fmt.Scanln(&inputFile)
		}

		if inputFile == "" {
			log.Fatal("Error: No HTML file provided")
		}

		// Check file extension
		if filepath.Ext(inputFile) != ".html" {
			log.Fatal("Error: Input file must have a .html extension")
		}

		// Start the loading spinner
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Start()
		s.Suffix = " Converting..."

		// Read the HTML file
		htmlContent, err := utils.ReadFile(inputFile)
		if err != nil {
			s.Stop()
			log.Fatalf("Read error: %v", err)
		}

		// Convert to Markdown
		markdown, err := converter.ConvertHTMLtoMarkdown(string(htmlContent))
		if err != nil {
			s.Stop()
			log.Fatalf("Conversion error: %v", err)
		}

		// Stop the spinner after conversion
		s.Stop()

		if outputFile == "" {
			fmt.Println("Enter the path to the Markdown file:")
			fmt.Scanln(&outputFile)
		}

		// Check file extension
		if filepath.Ext(outputFile) != ".md" {
			log.Fatal("Error: Output file must have a .md extension")
		}

		if outputFile != "" {
			if err := utils.WriteFile(outputFile, []byte(markdown)); err != nil {
				log.Fatalf("Write error: %v", err)
			}
			fmt.Printf("Conversion successful! Markdown saved to %s %s \n", outputFile, constants.Signature)
		} else {
			fmt.Println("\nConversion result:")
			fmt.Println("------------------------------")
			fmt.Println(markdown)
		}
		return nil
	},
}
