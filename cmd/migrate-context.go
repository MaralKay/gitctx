package main

import (
	"fmt"
	"os"
	"strings"
)

func migrateCurrentContextFile(currentContextPath string) {

	// Read the content of the file
	content, err := os.ReadFile(currentContextPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	if strings.Contains(string(content), "global: ") {
		fmt.Println("Nothing to migrate here.")
		os.Exit(0)
	}

	// Modify the content by adding "global: " in front of the existing content
	newContent := []byte("global: " + strings.TrimSpace(string(content)))

	// Write the modified content back to the file
	err = os.WriteFile(currentContextPath, newContent, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}

	fmt.Println("Context file migrated successfully.")
}
