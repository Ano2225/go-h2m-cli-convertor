package main

import (
	"log"
	"github.com/tonuser/markdown-to-html/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
