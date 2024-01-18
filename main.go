package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

var configFileName = "~/.gitctx.config"

func main() {
	initCommand := flag.NewFlagSet("init", flag.ExitOnError)
	showCommand := flag.NewFlagSet("show", flag.ExitOnError)

	// Parse the command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: gitctx <subcommand> [options]")
		os.Exit(1)
	}

	// Parse the subcommand
	subcommand := os.Args[1]
	switch subcommand {
	case "init":
		initCommand.Parse(os.Args[2:])
		initSymlink()
	case "show":
		showCommand.Parse(os.Args[2:])
		showContext()
	default:
		nameSymlink(subcommand)
	}
}

func showContext() {
	cmd := exec.Command("bash", "-c", "git config --list")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running command: %v\n", err)
		fmt.Println("Standard Error Output:", string(output))
		return
	}

	fmt.Println(string(output))
}

func getUserHome() (string, error) {
	// Get the current user's information
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("Error: Unable to get current user information. %v\n", err)
		os.Exit(1)
	}

	// Get the absolute path to the home directory
	return currentUser.HomeDir, err
}

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

	userHomePath, err := getUserHome()
	if err != nil {
		fmt.Print("Error: Unable to get user's home path")
		os.Exit(1)
	}

	configFileName = userHomePath + "/.gitctx.config"

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

	// Set the target file to .gitconfig
	targetFile := userHomePath + "/.gitconfig"

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
}

func nameSymlink(name string) {

	userHomePath, err := getUserHome()
	if err != nil {
		fmt.Print("Error: Unable to get user's home path")
		os.Exit(1)
	}

	configFileName = userHomePath + "/.gitctx.config"

	// Read the configuration file
	config, err := readConfig()
	if err != nil {
		fmt.Printf("Error: Unable to read configuration. %v\n", err)
		os.Exit(1)
	}

	// Look for the name in the keys of the key-value pairs
	if path, ok := config[name]; ok {
		// Set the target file to .gitconfig
		targetFile := userHomePath + "/.gitconfig"

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
			return
		}

	} else {
		fmt.Printf("Error: Name '%s' not found in configuration.\n", name)
		os.Exit(1)
	}
}

func saveConfig(name, gitConfigPath string, sshConfigPath string) error {
	file, err := os.OpenFile(configFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "%s=%s=%s\n", name, gitConfigPath, sshConfigPath)
	if err != nil {
		return err
	}

	return nil
}

func readConfig() (map[string][]string, error) {
	config := make(map[string][]string)

	file, err := os.Open(configFileName)
	if err != nil {
		if os.IsNotExist(err) {
			return config, nil // File does not exist, return empty config
		}
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := parseConfigLine(line)
		if len(parts) == 3 {
			var pathsArray []string
			pathsArray = append(pathsArray, parts[1], parts[2])
			config[parts[0]] = pathsArray
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return config, nil
}

func parseConfigLine(line string) []string {
	parts := strings.SplitN(line, "=", 3)
	if len(parts) == 3 {
		return parts
	}
	return []string{}
}
