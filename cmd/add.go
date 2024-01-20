package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func initSymlink() {
	// Prompt for the .gitconfig path
	fmt.Print("Enter the absolute path to a .gitconfig file: ")
	var gitconfigPath string
	fmt.Scanln(&gitconfigPath)

	// Check if the .gitconfig file exists
	if _, err := os.Stat(gitconfigPath); err != nil {
		fmt.Print(err)
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

	// Get the absolute path of the target file
	targetAbsPath, err := filepath.Abs(targetFile)
	if err != nil {
		fmt.Printf("Error: Unable to get absolute path for '%s'.\n", targetFile)
		os.Exit(1)
	}

	// Store the name and path in the configuration file
	if err := saveConfig(name, gitconfigAbsPath, sshConfigAbsPath); err != nil {
		fmt.Printf("Error: Unable to save configuration. %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Configuration saved.")

	cmd := exec.Command("bash", "-c", "git config core.sshCommand \"ssh -F "+sshConfigAbsPath+"\"")

	// Run the command and capture its output
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running command: %v\n", err)
		fmt.Println("Standard Error Output:", string(output))
		return
	}

	// Check if a symlink with the target file name already exists
	if _, err := os.Stat(targetAbsPath); err == nil {
		// Check if the existing symlink already points to the specified .gitconfig file
		currentTarget, err := os.Readlink(targetAbsPath)
		if err == nil && currentTarget == gitconfigAbsPath {
			fmt.Println("Symlink already exists and points to the specified .gitconfig file. Doing nothing.")

			if err := updateContext(name, currentContextPath); err != nil {
				fmt.Printf("Error: Unable to update context")
				os.Exit(1)
			}

			fmt.Printf("Updated context to %s\n", name)
			os.Exit(0)
		} else {
			fmt.Printf("Removing existing symlink: %s\n", targetAbsPath)
			if err := os.Remove(targetAbsPath); err != nil {
				fmt.Printf("Error: Unable to remove existing symlink '%s'.\n", targetAbsPath)
				os.Exit(1)
			}
		}
	}

	// Create a new symlink
	if err := os.Symlink(gitconfigAbsPath, targetAbsPath); err != nil {
		fmt.Printf("Error: Unable to create symlink '%s' -> '%s'.\n", targetAbsPath, gitconfigAbsPath)
		os.Exit(1)
	}

	fmt.Printf("Symlink created: %s -> %s\n", targetAbsPath, gitconfigAbsPath)

	if err := updateContext(name, currentContextPath); err != nil {
		fmt.Printf("Error: Unable to update context")
		os.Exit(1)
	}

	fmt.Printf("Updated context to %s\n", name)
}
