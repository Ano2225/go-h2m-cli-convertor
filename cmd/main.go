package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd représente la commande racine
var RootCmd = &cobra.Command{
	Use:   "H2M",
	Short: "H2M est un Convertisseur de fichiers Markdown et HTML",
	Long:  `Un outil CLI complet pour convertir entre Markdown et HTML`,
}

// Execute initialise Cobra
func Execute() error {
	return RootCmd.Execute()
}

func init() {
	// Ajouter les sous-commandes
	RootCmd.AddCommand(mdToHtmlCmd)
	RootCmd.AddCommand(htmlToMdCmd)
}
