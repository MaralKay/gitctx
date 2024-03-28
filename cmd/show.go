package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func showContext() {

	// Open the file
	file, err := os.OpenFile(currentContextPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		if verboseFlag {
			fmt.Printf("Error opening file: %v", err)
		}
		fmt.Printf("Error while updating context")
	}
	defer file.Close()

	repoName := ""
	if isGitRepo() {
		repoName = strings.TrimSpace(getRepoName())
	}
	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	var lines []string

	// Read lines from the file
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ": ", 2)
		if repoName != "" {
			if len(parts) == 2 && strings.TrimSpace(parts[0]) == repoName {
				lines = append(lines, line)
			}
			if len(parts) == 2 && strings.TrimSpace(parts[0]) == "global" {
				lines = append(lines, line)
			}
		} else {
			if len(parts) == 2 && strings.TrimSpace(parts[0]) == "global" {
				lines = append(lines, line)
			}
		}
	}
	// Check for any scanner errors
	if err := scanner.Err(); err != nil {
		if verboseFlag {
			fmt.Printf("Error scanning file: %v", err)
		}
		fmt.Printf("Error while updating context")
	}

	for _, l := range lines {
		fmt.Println(l)
	}
}
