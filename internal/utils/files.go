package utils

import (
	"os"
)

// ReadFile read a file and returns its content
func ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

//WriteFile write a file with the given content
func WriteFile(filePath string, content []byte) error {
	return os.WriteFile(filePath, content, 0644)
}
