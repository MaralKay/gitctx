package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func updateContext(contextName string, currentContextPath string, repoName string, verboseFlag bool) error {

	// Open the file
	file, err := os.OpenFile(currentContextPath, os.O_RDWR|os.O_CREATE, 0644)
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
	found := false

	// Read lines from the file
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) == 2 && strings.TrimSpace(parts[0]) == repoName {
			// Found the line with the repository name, update the context
			line = parts[0] + ": " + contextName
			found = true
		}
		if repoName != "global" {
			if len(parts) == 2 && strings.TrimSpace(parts[0]) == "global" {
				// Found the line with the repository name, update the context
				line = parts[0] + ": " + contextName
			}
		}
		lines = append(lines, line)
	}
	if !found {
		newContext := repoName + ": " + contextName
		newContext = strings.ReplaceAll(newContext, "\n", "")
		lines = append(lines, newContext)
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
	return nil
}
