package main

import (
	"fmt"
	"os"
)

func updateContext(contextName string, currentContextPath string) error {
	// Check if the file exists
	if _, err := os.Stat(currentContextPath); err == nil {
		// File exists, so delete it
		err := os.Remove(currentContextPath)
		if err != nil {
			return fmt.Errorf("Error deleting file: %v", err)
		}
	}

	// Create a new empty file
	file, err := os.Create(currentContextPath)
	if err != nil {
		return fmt.Errorf("Error creating file: %v", err)
	}
	defer file.Close()

	// Write content to the file
	_, err = file.WriteString(contextName)
	if err != nil {
		return fmt.Errorf("Error writing to file: %v", err)
	}

	fmt.Printf("File %s deleted, recreated, and written with content successfully.\n", currentContextPath)
	return nil
}
