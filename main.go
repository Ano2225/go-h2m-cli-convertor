package main

import (
	"log"
	"github.com/tonuser/markdown-to-html/cmd"
)

func main() {
	// Initialise Cobra(bibliotheque utilisée pour ecrire les commandes en cli) et verifie qu'il n'y a pas d'erreur
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
