package main

import (
	"log"
	"github.com/tonuser/markdown-to-html/cmd"
)

func main() {
	// Run the root command
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
