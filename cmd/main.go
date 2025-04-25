package cmd

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/tonuser/markdown-to-html/internal/vars/constants"
	"github.com/tonuser/markdown-to-html/internal/vars/styles"
)

var (
	AppVersion = "1.0.0" // Version of the application
	OutputFile string    // Output file name
	InputFile  string    // Input file name

	//go:embed ascii-draw.txt
	Logo string // Logo ASCII art
)

// formatLongDescription formats the application's long description
// and includes the logo, main features and usage examples
func formatLongDescription() string {
	desc := styles.Title.Sprint(Logo)
	desc += "\n" + styles.Help.Sprint(`Un outil CLI pour la conversion de fichiers.

    Fonctionnalités principales:
    - Conversion Markdown → HTML
    - Conversion HTML → Markdown

    Exemples d'utilisation:
    `) + styles.Command.Sprint(`
    md2html -i document.md -o resultat.html
    html2md -i page.html -o resultat.md
    `)
	return desc
}

// RootCmd is the root command of the application
var RootCmd = &cobra.Command{
	Use:     "H2M",
	Short:   "H2M est un Convertisseur de fichiers Markdown et HTML",
	Long:    formatLongDescription(),
	Version: AppVersion,
}

// REPL (Read-Eval-Print Loop) for interactive mode
func startREPL() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nEntrez une commande (ou tapez 'exit' pour quitter) :")

	// signal handling for graceful shutdown
	// This will allow the program to exit cleanly on Ctrl+C or SIGTERM
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println("\nFermeture de H2M...")
		os.Exit(0)
	}()

	for {
		fmt.Print(styles.Title.Sprint("H2M> "))
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(styles.Error.Sprint("Erreur de lecture de l'entrée : "), err)
			continue
		}
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Fermeture de H2M...")
			break
		}

		// Parse the input and set it as arguments for the command
		args := strings.Split(input, " ")
		RootCmd.SetArgs(args)

		if err := RootCmd.Execute(); err != nil {
			fmt.Println(styles.Error.Sprint("Erreur d'exécution de la commande : "), err)
		}
	}
}

// Execute runs the root command and handles the command line arguments
func Execute() error {

	// check if the user requested help
	for _, arg := range os.Args {
		if arg == "--help" || arg == "-h" {
			return RootCmd.Execute()
		}
	}

	// print the logo and long description
	fmt.Println(formatLongDescription())
	if len(os.Args) > 1 {
		if err := RootCmd.Execute(); err != nil {
			fmt.Println(styles.Error.Sprint("Erreur d'exécution de la commande : "), err)
			return err
		}
	} else {
		// if no arguments are provided, start the REPL
		startREPL()
	}
	return nil
}

// setupCommands configures the commands and flags for the application
func setupCommands() {
	cobra.AddTemplateFunc("styleLine", func(s string) string {
		return styles.Help.Sprint(s)
	})

	RootCmd.SetUsageTemplate(constants.UsageTemplate)

	// Configuration
	addGlobalFlags()
	addConversionCommands()
}

func init() {
	setupCommands()
}

// addGlobalFlags adds global flags to the root command
func addGlobalFlags() {
	RootCmd.PersistentFlags().StringVarP(&OutputFile, "output", "o", "",
		styles.Help.Sprint("fichier de sortie (obligatoire)"))
	RootCmd.PersistentFlags().StringVarP(&InputFile, "input", "i", "",
		styles.Help.Sprint("fichier d'entrée (obligatoire)"))
	RootCmd.Flags().BoolP("version", "V", false,
		styles.Help.Sprint("affiche la version"))
}

// addConversionCommands adds the conversion commands to the root command
func addConversionCommands() {
	RootCmd.AddCommand(mdToHtmlCmd)
	RootCmd.AddCommand(htmlToMdCmd)
}
