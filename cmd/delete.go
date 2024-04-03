package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func deleteContext(contextName string, configFileName string, verboseFlag bool) error {

	fmt.Print("Are you sure you want to delete context " + contextName + " ?     [y/n]")
	var answer string
	fmt.Scanln(&answer)

	// Check if the .gitconfig file exists
	if answer != "y" {
		fmt.Printf("Delete cancelled")
		os.Exit(0)
	}

	// Open the file
	file, err := os.OpenFile(configFileName, os.O_RDWR, 0644)
	if err != nil {
		if verboseFlag {
			fmt.Printf("Error opening file: %v", err)
		}
		return fmt.Errorf("Error while updating context")
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	var lines []string

	// Read lines from the file
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if parts[0] != contextName {
			lines = append(lines, line)
		}
	}

	// Check for any scanner errors
	if err := scanner.Err(); err != nil {
		if verboseFlag {
			fmt.Printf("Error scanning file: %v", err)
		}
		return fmt.Errorf("Error while updating context")
	}

	// Truncate the file before writing new content
	if err := file.Truncate(0); err != nil {
		if verboseFlag {
			fmt.Printf("Error truncating file: %v", err)
		}
		return fmt.Errorf("Error while updating context")
	}

	// Move the file offset to the beginning
	if _, err := file.Seek(0, 0); err != nil {
		if verboseFlag {
			fmt.Printf("Error seeking file: %v", err)
		}
		return fmt.Errorf("Error while updating context")
	}

	// Write the modified lines back to the file
	for _, line := range lines {
		_, err := fmt.Fprintln(file, line)
		if err != nil {
			if verboseFlag {
				fmt.Printf("Error writing to file: %v", err)
			}
			return fmt.Errorf("Error while updating context")
		}
	}

	fmt.Printf("Context %v deleted", contextName)
	return nil
}
