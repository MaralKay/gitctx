package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func nameSymlink(name string) {

	// Read the configuration file
	config, err := readConfig()
	if err != nil {
		fmt.Printf("Error: Unable to read configuration. %v\n", err)
		os.Exit(1)
	}

	// Look for the name in the keys of the key-value pairs
	if path, ok := config[name]; ok {
		// Get the absolute path of the target file
		targetAbsPath, err := filepath.Abs(targetFile)
		if err != nil {
			fmt.Printf("Error: Unable to get absolute path for '%s'.\n", targetFile)
			os.Exit(1)
		}

		// Check if a symlink with the target file name already exists
		if _, err := os.Stat(targetAbsPath); err == nil {
			// Check if the existing symlink already points to the specified .gitconfig file
			currentTarget, err := os.Readlink(targetAbsPath)
			if err == nil && currentTarget == path[0] {
				fmt.Println("Symlink already exists and points to the specified .gitconfig file. Doing nothing.")
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
		if err := os.Symlink(path[0], targetAbsPath); err != nil {
			fmt.Printf("Error: Unable to create symlink '%s' -> '%s'.\n", targetAbsPath, path)
			os.Exit(1)
		}

		fmt.Printf("Symlink created: %s -> %s\n", targetAbsPath, path)

		cmd := exec.Command("bash", "-c", "git config core.sshCommand \"ssh -F "+path[1]+"\"")

		// Run the command and capture its output
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error running command: %v\n", err)
			fmt.Println("Standard Error Output:", string(output))
		}

		if err := updateContext(name, currentContextPath); err != nil {
			fmt.Printf("Error: Unable to update context")
			os.Exit(1)
		}

		fmt.Printf("Updated context to %s\n", name)

	} else {
		fmt.Printf("Error: Name '%s' not found in configuration.\n", name)
		os.Exit(1)
	}
}
