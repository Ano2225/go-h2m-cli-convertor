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

// md2html Command
var mdToHtmlCmd = &cobra.Command{
    Use:     "md2html",
    Aliases: []string{"markdown-to-html"},
    Short:   "Convert Markdown to HTML",
    Long:    "Convert a Markdown file to HTML",
    RunE: func(cmd *cobra.Command, args []string) error {
        inputFile, _ := cmd.Flags().GetString("input")
        outputFile, _ := cmd.Flags().GetString("output")

        if inputFile == "" {
            fmt.Println("Enter the path to the Markdown file:")
            fmt.Scanln(&inputFile)
        }
        if inputFile == "" {
            log.Fatal("Error: No Markdown file provided")
        }

        // Check file extension
        if filepath.Ext(inputFile) != ".md" {
            log.Fatal("Error: Input file must have a .md extension")
        }

        // Start the loading spinner
        s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
        s.Start()
        s.Suffix = " Converting..."

        // Read the Markdown file
        markdown, err := utils.ReadFile(inputFile)
        if err != nil {
            s.Stop()
            log.Fatalf("Read error: %v", err)
        }

        // Convert to HTML
        html, err := converter.ConvertMarkdownToHTML(markdown)
        if err != nil {
            s.Stop()
            log.Fatalf("Conversion error: %v", err)
        }

        // Stop the spinner after conversion
        s.Stop()

        if outputFile == "" {
            fmt.Println("Enter the path to the HTML file:")
            fmt.Scanln(&outputFile)
        }

        // Check file extension
        if filepath.Ext(outputFile) != ".html" {
            log.Fatal("Error: Output file must have a .html extension")
        }

        if outputFile != "" {
            if err := utils.WriteFile(outputFile, []byte(html)); err != nil {
                log.Fatalf("Write error: %v", err)
            }
            fmt.Printf("Conversion successful! HTML saved to %s %s\n", outputFile, utils.Signature)
        } else {
            fmt.Println("\nConversion result:")
            fmt.Println("-------------------------------")
            fmt.Println(html)
        }
        return nil
    },
}
