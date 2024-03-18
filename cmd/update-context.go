package main

import (
	"fmt"
	"os"
)

func updateContext(contextName string, currentContextPath string, verboseFlag bool) error {
	// Check if the file exists
	if _, err := os.Stat(currentContextPath); err == nil {
		// File exists, so delete it
		err := os.Remove(currentContextPath)
		if err != nil {
			if verboseFlag {
				fmt.Printf("Error deleting file: %v", err)
			}
			return fmt.Errorf("Error while updating context")
		}
	}

	// Create a new empty file
	file, err := os.Create(currentContextPath)
	if err != nil {
		if verboseFlag {
			fmt.Printf("Error creating file: %v", err)
		}
		return fmt.Errorf("Error while updating context")
	}
	defer file.Close()

	// Write content to the file
	_, err = file.WriteString(contextName)
	if err != nil {
		if verboseFlag {
			fmt.Printf("Error writing to file: %v", err)
		}
		return fmt.Errorf("Error while updating context")
	}
	return nil
}
