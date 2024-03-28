package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func addContext(verboseFlag bool) {
	// Prompt for the .gitconfig path
	fmt.Print("Enter the absolute path to a .gitconfig file: ")
	var gitconfigPath string
	fmt.Scanln(&gitconfigPath)

	// Check if the .gitconfig file exists
	if _, err := os.Stat(gitconfigPath); err != nil {
		if verboseFlag {
			fmt.Print(err)
		}
		fmt.Printf("Error: .gitconfig file '%s' not found.\n", gitconfigPath)
		os.Exit(1)
	}

	// Prompt for a name for this .gitconfig context
	fmt.Print("Enter a name for your .gitconfig context: ")
	var name string
	fmt.Scanln(&name)

	// Get the absolute path of the .gitconfig file
	gitconfigAbsPath, err := filepath.Abs(gitconfigPath)
	if err != nil {
		fmt.Printf("Error: Unable to get absolute path for '%s'.\n", gitconfigPath)
		os.Exit(1)
	}

	// Prompt for an ssh config to use for this .gitconfig context
	fmt.Print("Enter the absolute path for your .ssh config to use with this context: ")
	var sshConfig string
	fmt.Scanln(&sshConfig)

	// Get the absolute path of the ssh config file
	sshConfigAbsPath, err := filepath.Abs(sshConfig)
	if err != nil {
		fmt.Printf("Error: Unable to get absolute path for '%s'.\n", sshConfigAbsPath)
		os.Exit(1)
	}

	// Store the name and path in the configuration file
	if err := saveConfig(name, gitconfigAbsPath, sshConfigAbsPath); err != nil {
		fmt.Printf("Error: Unable to save configuration. %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Configuration saved.")

	switchContext(name, verboseFlag)
}
