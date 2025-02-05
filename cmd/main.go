package cmd

import (
    "bufio"
    "fmt"
    "os"
    "strings"
	"os/signal"
    "syscall"
    "github.com/spf13/cobra"
    "github.com/tonuser/markdown-to-html/internal/utils"
)

// crée une description détaillée stylisée
func formatLongDescription() string {
    desc := utils.TitleStyle.Sprint(utils.Logo)
    desc += "\n" + utils.HelpStyle.Sprint(`Un outil CLI pour la conversion de fichiers.

    Fonctionnalités principales:
    - Conversion Markdown → HTML
    - Conversion HTML → Markdown

    Exemples d'utilisation:
    `) + utils.CommandStyle.Sprint(`
    md2html -i document.md -o resultat.html
    html2md -i page.html -o resultat.md
    `)
    return desc
}

// RootCmd représente la commande racine
var RootCmd = &cobra.Command{
    Use:     "H2M",
    Short:   "H2M est un Convertisseur de fichiers Markdown et HTML",
    Long:    formatLongDescription(),
    Version: utils.AppVersion,
}

// REPL (Read-Eval-Print Loop) pour garder le programme actif
func startREPL() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("\nEntrez une commande (ou tapez 'exit' pour quitter) :")

    // Création d'un channel pour écouter les signaux OS
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

    // Goroutine pour écouter les signaux
    go func() {
        <-sigs
        fmt.Println("\nFermeture de H2M...")
        os.Exit(0) // Quitte proprement
    }()

    for {
        fmt.Print(utils.TitleStyle.Sprint("H2M> "))
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println(utils.ErrorStyle.Sprint("Erreur de lecture de l'entrée : "), err)
            continue
        }
        input = strings.TrimSpace(input)

        if input == "exit" {
            fmt.Println("Fermeture de H2M...")
            break
        }

        // Simuler l'exécution d'une commande
        args := strings.Split(input, " ")
        RootCmd.SetArgs(args)

        if err := RootCmd.Execute(); err != nil {
            fmt.Println(utils.ErrorStyle.Sprint("Erreur d'exécution de la commande : "), err)
        }
    }
}

// Lance l'application et active le mode interactif
func Execute() error {
    fmt.Println(formatLongDescription()) // Affiche la description au début
    if len(os.Args) > 1 {
        if err := RootCmd.Execute(); err != nil {
            fmt.Println(utils.ErrorStyle.Sprint("Erreur d'exécution de la commande : "), err)
            return err
        }
    } else {
        // Activer le mode interactif si aucune commande n'est donnée
        startREPL()
    }
    return nil
}

// initialise toutes les commandes et leur style
func setupCommands() {
    // Personnalisation du template d'aide
    cobra.AddTemplateFunc("styleLine", func(s string) string {
        return utils.HelpStyle.Sprint(s)
    })

    RootCmd.SetUsageTemplate(utils.UsageTemplate)

    // Configuration
    addGlobalFlags()
    addConversionCommands()
}

func init() {
    setupCommands()
}

// addGlobalFlags configure les flags globaux
func addGlobalFlags() {
    RootCmd.PersistentFlags().StringVarP(&utils.OutputFile, "output", "o", "",
        utils.HelpStyle.Sprint("fichier de sortie (obligatoire)"))
    RootCmd.PersistentFlags().StringVarP(&utils.InputFile, "input", "i", "",
        utils.HelpStyle.Sprint("fichier d'entrée (obligatoire)"))
    RootCmd.Flags().BoolP("version", "V", false,
        utils.HelpStyle.Sprint("affiche la version"))
}

// Ajoute les commandes de conversions
func addConversionCommands() {
    // Ajouter les sous-commandes
    RootCmd.AddCommand(mdToHtmlCmd)
    RootCmd.AddCommand(htmlToMdCmd)
}