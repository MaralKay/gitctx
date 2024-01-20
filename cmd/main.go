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

func main() {
	addCommand := flag.NewFlagSet("add", flag.ExitOnError)
	showCommand := flag.NewFlagSet("show", flag.ExitOnError)

	// Parse the command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: gitctx <subcommand> [options]")
		os.Exit(1)
	}

	// Parse the subcommand
	subcommand := os.Args[1]
	switch subcommand {
	case "add":
		addCommand.Parse(os.Args[2:])
		initSymlink()
	case "show":
		showCommand.Parse(os.Args[2:])
		showContext()
	default:
		nameSymlink(subcommand)
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