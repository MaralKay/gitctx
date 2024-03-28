package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func switchContext(name string, verboseFlag bool) {

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
				if verboseFlag {
					fmt.Println("Symlink already exists and points to the specified .gitconfig file. Doing nothing.")
					goto jumpToSSHConfig
				}
			} else {
				if verboseFlag {
					fmt.Printf("Removing existing symlink: %s\n", targetAbsPath)
				}

				if err := os.Remove(targetAbsPath); err != nil {
					if verboseFlag {
						fmt.Printf("Error: Unable to remove existing symlink '%s'.\n", targetAbsPath)
					}
					fmt.Println("Error: Unable to switch context.")
					os.Exit(1)
				}
			}
		}

		// Create a new symlink
		createSymlink(path, targetAbsPath)
		goto jumpToSSHConfig

	jumpToSSHConfig:
		configureSSH(path[1])

		if err := updateContext(name, currentContextPath, verboseFlag); err != nil {
			fmt.Printf("Error: Unable to update context")
			os.Exit(1)
		}

		fmt.Printf("Updated context to %s\n", name)

	} else {
		fmt.Printf("Error: Name '%s' not found in configuration.\n", name)
		os.Exit(1)
	}
}

func createSymlink(path []string, targetPath string) {
	// Create a new symlink
	if err := os.Symlink(path[0], targetPath); err != nil {
		if verboseFlag {
			fmt.Printf("Error: Unable to create symlink '%s' -> '%s'.\n", targetPath, path)
		}
		fmt.Println("Error: Unable to switch context.")
		os.Exit(1)
	}
	if verboseFlag {
		fmt.Printf("Symlink created: %s -> %s\n", targetPath, path)
	}
}

func configureSSH(path string) {

	if isGitRepo() {
		cmd := exec.Command("bash", "-c", "git config core.sshCommand \"ssh -F "+path+"\"")

		// Run the command and capture its output
		output, err := cmd.CombinedOutput()
		if err != nil {
			if verboseFlag {
				fmt.Println("Standard Error Output:", string(output))
			}
			fmt.Printf("Error setting up ssh config: %v\n", err)
			os.Exit(1)
		}
	} else {
		cmd := exec.Command("bash", "-c", "git config --global core.sshCommand \"ssh -F "+path+"\"")

		// Run the command and capture its output
		output, err := cmd.CombinedOutput()
		if err != nil {
			if verboseFlag {
				fmt.Println("Standard Error Output:", string(output))
			}
			fmt.Printf("Error setting up ssh config: %v\n", err)
			os.Exit(1)
		}
	}
}

func isGitRepo() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}
