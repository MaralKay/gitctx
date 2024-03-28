package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
)

var userHomePath = getUserHome()
var configFileName = userHomePath + "/.gitctx.config"
var currentContextPath = userHomePath + "/.gitctx.current"
var targetFile = userHomePath + "/.gitconfig"
var verboseFlag = false

func main() {
	helpCommand := flag.NewFlagSet("-h", flag.ExitOnError)
	addCommand := flag.NewFlagSet("add", flag.ExitOnError)
	showCommand := flag.NewFlagSet("show", flag.ExitOnError)
	migrateContextCommand := flag.NewFlagSet("migrate", flag.ExitOnError)

	// Set Usage message for commands
	helpCommand.Usage = showHelp
	addCommand.Usage = showHelp
	showCommand.Usage = showHelp
	migrateContextCommand.Usage = showHelp

	// Parse the command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: gitctx <subcommand> [options]")
		os.Exit(1)
	}

	if len(os.Args) == 3 {
		if os.Args[1] == "-v" || os.Args[2] == "-v" {
			verboseFlag = true
		}
	} else if len(os.Args) == 4 {
		if os.Args[2] == "-v" || os.Args[3] == "-v" {
			verboseFlag = true
		}
	}

	// Parse the subcommand
	subcommand := os.Args[1]
	switch subcommand {
	case "-h":
		helpCommand.Parse(os.Args[2:])
		showHelp()
	case "list":
		listContexts()
	case "add":
		addCommand.Parse(os.Args[2:])
		addContext(verboseFlag)
	case "show":
		showCommand.Parse(os.Args[2:])
		showContext()
	case "migrate":
		migrateCurrentContextFile(currentContextPath)
	default:
		switchContext(subcommand, verboseFlag)
	}
}

func getUserHome() string {
	// Get the current user's information
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("Error: Unable to get current user information. %v\n", err)
		os.Exit(1)
	}

	return currentUser.HomeDir
}
