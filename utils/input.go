package utils

import (
	"log"
	"os"
	"strings"
)

// reads the input for a given day
func ReadInput(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read input file: %s", err)
	}
	return strings.TrimSpace(string(data))
}
