package utils

import (
	"os"
)

// ReadFile lit le contenu d'un fichier et le retourne sous forme de bytes
func ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

//WriteFile ecrit du contenu dans un fichier
func WriteFile(filePath string, content []byte) error {
	return os.WriteFile(filePath, content, 0644)
}